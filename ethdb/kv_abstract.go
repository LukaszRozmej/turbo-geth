package ethdb

import (
	"context"
	"errors"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
)

var (
	ErrAttemptToDeleteNonDeprecatedBucket = errors.New("only buckets from dbutils.DeprecatedBuckets can be deleted")
	ErrUnknownBucket                      = errors.New("unknown bucket. add it to dbutils.Buckets")
)

type KV interface {
	View(ctx context.Context, f func(tx Tx) error) error
	Update(ctx context.Context, f func(tx Tx) error) error
	Close()

	Begin(ctx context.Context, parent Tx, writable bool) (Tx, error)
	AllBuckets() dbutils.BucketsCfg
}

type Tx interface {
	Cursor(bucket string) Cursor
	CursorDupSort(bucket string) CursorDupSort
	CursorDupFixed(bucket string) CursorDupFixed
	Get(bucket string, key []byte) (val []byte, err error)

	Commit(ctx context.Context) error
	Rollback()
	BucketSize(name string) (uint64, error)
}

// Interface used for buckets migration, don't use it in usual app code
type BucketMigrator interface {
	DropBucket(string) error
	CreateBucket(string) error
	ExistsBucket(string) bool
	ClearBucket(string) error
	ExistingBuckets() ([]string, error)
}

type Cursor interface {
	Prefix(v []byte) Cursor
	Prefetch(v uint) Cursor

	First() ([]byte, []byte, error)
	Seek(seek []byte) ([]byte, []byte, error)
	SeekExact(key []byte) ([]byte, error)
	Next() ([]byte, []byte, error) // Next - returns next key/value (can iterate over DupSort key/values automatically)
	Prev() ([]byte, []byte, error)
	// Last - returns last key and last possible value
	Last() ([]byte, []byte, error)

	Put(key, value []byte) error
	// PutNoOverwrite(key, value []byte) error
	// Reserve()

	// PutCurrent - replace the item at the current cursor position.
	//	The key parameter must still be provided, and must match it.
	//	If using sorted duplicates (#MDB_DUPSORT) the data item must still
	//	sort into the same place. This is intended to be used when the
	//	new data is the same size as the old. Otherwise it will simply
	//	perform a delete of the old record followed by an insert.
	PutCurrent(key, value []byte) error
	// Current - return key/data at current cursor position
	Current() ([]byte, []byte, error)

	// DeleteCurrent This function deletes the key/data pair to which the cursor refers.
	// This does not invalidate the cursor, so operations such as MDB_NEXT
	// can still be used on it.
	// Both MDB_NEXT and MDB_GET_CURRENT will return the same record after
	// this operation.
	DeleteCurrent() error
	Delete(key []byte) error
	Append(key []byte, value []byte) error // Returns error if provided data not sorted or has duplicates
}

type CursorDupSort interface {
	Cursor

	SeekBothExact(key, value []byte) ([]byte, []byte, error)
	SeekBothRange(key, value []byte) ([]byte, []byte, error)
	FirstDup() ([]byte, error)
	NextDup() ([]byte, []byte, error)   // NextDup - iterate only over duplicates of current key
	NextNoDup() ([]byte, []byte, error) // NextNoDup - iterate with skipping all duplicates
	LastDup() ([]byte, error)

	CountDuplicates() (uint64, error)  // Count returns the number of duplicates for the current key. See mdb_cursor_count
	DeleteCurrentDuplicates() error    // Delete all of the data items for the current key
	AppendDup(key, value []byte) error // Returns error if provided data not sorted or has duplicates

	//PutIfNoDup()      // Store the key-value pair only if key is not present
}

type CursorDupFixed interface {
	CursorDupSort

	// GetMulti - return up to a page of duplicate data items from current cursor position
	// After return - move cursor to prepare for #MDB_NEXT_MULTIPLE
	// See also lmdb.WrapMulti
	GetMulti() ([]byte, error)
	// NextMulti - return up to a page of duplicate data items from next cursor position
	// After return - move cursor to prepare for #MDB_NEXT_MULTIPLE
	// See also lmdb.WrapMulti
	NextMulti() ([]byte, []byte, error)
	// PutMulti store multiple contiguous data elements in a single request.
	// Panics if len(page) is not a multiple of stride.
	// The cursor's bucket must be DupFixed and DupSort.
	PutMulti(key []byte, page []byte, stride int) error
	// ReserveMulti()
}

type HasStats interface {
	DiskSize(context.Context) (uint64, error) // db size
}

type Backend interface {
	AddLocal([]byte) ([]byte, error)
	Etherbase() (common.Address, error)
	NetVersion() (uint64, error)
	BloomStatus() (uint64, uint64, common.Hash)
}

type DbProvider uint8

const (
	Bolt DbProvider = iota
	Remote
	Lmdb
)
