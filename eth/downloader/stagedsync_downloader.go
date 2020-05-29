package downloader

import (
	"fmt"

	"github.com/ledgerwatch/turbo-geth/core"
	"github.com/ledgerwatch/turbo-geth/log"
)

func (d *Downloader) doStagedSyncWithFetchers(p *peerConnection, headersFetchers []func() error) error {
	var err error
	defer log.Info("Staged sync finished")

	/*
	* Stage 1. Download Headers
	 */
	log.Info("Sync stage 1/7. Downloading headers...")
	err = d.shallQuit(func() error {
		return d.DownloadHeaders(headersFetchers)
	})
	if err != nil {
		return err
	}

	/*
	* Stage 2. Download Block bodies
	 */
	log.Info("Sync stage 2/7. Downloading block bodies...")
	cont := true

	for cont && err == nil {
		err = d.shallQuit(func() error {
			return d.spawnBodyDownloadStage(p.id, &cont)
		})
		if err != nil {
			return err
		}
	}

	log.Info("Sync stage 2/7. Downloading block bodies... Complete!")

	/*
	* Stage 3. Recover senders from tx signatures
	 */
	log.Info("Sync stage 3/7. Recovering senders from tx signatures...")

	err = d.shallQuit(func() error {
		return d.spawnRecoverSendersStage()
	})
	if err != nil {
		return err
	}
	log.Info("Sync stage 3/7. Recovering senders from tx signatures... Complete!")

	/*
	* Stage 4. Execute block bodies w/o calculating trie roots
	 */
	log.Info("Sync stage 4/7. Executing blocks w/o hash checks...")
	var syncHeadNumber uint64
	err = d.shallQuit(func() error {
		return spawnExecuteBlocksStage(d.stateDB, d.blockchain, &syncHeadNumber, d.quitCh)
	})
	if err != nil {
		return err
	}

	log.Info("Sync stage 4/7. Executing blocks w/o hash checks... Complete!")

	// Further stages go there
	log.Info("Sync stage 5/7. Validating final hash")
	err = d.shallQuit(func() error {
		return spawnCheckFinalHashStage(d.stateDB, syncHeadNumber, d.datadir)
	})
	if err != nil {
		return err
	}

	log.Info("Sync stage 5/7. Validating final hash... Complete!")

	if d.history {
		log.Info("Sync stage 6/7. Generating account history index")
		err = d.shallQuit(func() error {
			return spawnAccountHistoryIndex(d.stateDB, d.datadir, core.UsePlainStateExecution, d.quitCh)
		})
		if err != nil {
			return err
		}
		log.Info("Sync stage 6/7. Generating account history index... Complete!")
	} else {
		log.Info("Sync stage 6/7, generating account history index is disabled. Enable by adding `h` to --storage-mode")
	}

	if d.history {
		log.Info("Sync stage 7/7. Generating storage history index")
		err = d.shallQuit(func() error {
			return spawnStorageHistoryIndex(d.stateDB, d.datadir, core.UsePlainStateExecution, d.quitCh)
		})
		if err != nil {
			return err
		}
		log.Info("Sync stage 7/7. Generating storage history index... Complete!")
	} else {
		log.Info("Sync stage 7/7, generating storage history index is disabled. Enable by adding `h` to --storage-mode")
	}

	return err
}

func (d *Downloader) DownloadHeaders(headersFetchers []func() error) error {
	err := d.shallQuit(func() error {
		return d.spawnSync(headersFetchers)
	})
	if err != nil {
		return err
	}

	log.Info("Sync stage 1/7. Downloading headers... Complete!")
	log.Info("Checking for unwinding...")
	// Check unwinds backwards and if they are outstanding, invoke corresponding functions
	for stage := Finish - 1; stage > Headers; stage-- {
		var unwindPoint uint64
		err = d.shallQuit(func() error {
			return GetStageUnwind(d.stateDB, stage, &unwindPoint)
		})
		if err != nil {
			return err
		}

		if unwindPoint == 0 {
			continue
		}

		var stageFn func() error
		switch stage {
		case Bodies:
			stageFn = func() error { return d.unwindBodyDownloadStage(unwindPoint) }
		case Senders:
			stageFn = func() error { return d.unwindSendersStage(unwindPoint) }
		case Execution:
			stageFn = func() error { return unwindExecutionStage(unwindPoint, d.stateDB) }
		case HashCheck:
			stageFn = func() error { return unwindHashCheckStage(unwindPoint, d.stateDB) }
		case AccountHistoryIndex:
			stageFn = func() error { return unwindAccountHistoryIndex(unwindPoint, d.stateDB, core.UsePlainStateExecution) }
		case StorageHistoryIndex:
			stageFn = func() error { return unwindStorageHistoryIndex(unwindPoint, d.stateDB, core.UsePlainStateExecution) }
		default:
			return fmt.Errorf("unrecognized stage for unwinding: %d", stage)
		}

		if err = d.shallQuit(stageFn); err != nil {
			return fmt.Errorf("error unwinding stage: %d: %v", stage, err)
		}
	}
	log.Info("Checking for unwinding... Complete!")
	return nil
}

func (d *Downloader) shallQuit(fn func() error) error {
	select {
	case <-d.quitCh:
		return errCanceled
	default:
	}
	return fn()
}
