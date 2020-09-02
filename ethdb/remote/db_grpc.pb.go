// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package remote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DBClient is the client API for DB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBClient interface {
	Size(ctx context.Context, in *SizeRequest, opts ...grpc.CallOption) (*SizeReply, error)
	BucketSize(ctx context.Context, in *BucketSizeRequest, opts ...grpc.CallOption) (*BucketSizeReply, error)
}

type dBClient struct {
	cc grpc.ClientConnInterface
}

func NewDBClient(cc grpc.ClientConnInterface) DBClient {
	return &dBClient{cc}
}

var dBSizeStreamDesc = &grpc.StreamDesc{
	StreamName: "Size",
}

func (c *dBClient) Size(ctx context.Context, in *SizeRequest, opts ...grpc.CallOption) (*SizeReply, error) {
	out := new(SizeReply)
	err := c.cc.Invoke(ctx, "/remote.DB/Size", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var dBBucketSizeStreamDesc = &grpc.StreamDesc{
	StreamName: "BucketSize",
}

func (c *dBClient) BucketSize(ctx context.Context, in *BucketSizeRequest, opts ...grpc.CallOption) (*BucketSizeReply, error) {
	out := new(BucketSizeReply)
	err := c.cc.Invoke(ctx, "/remote.DB/BucketSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBService is the service API for DB service.
// Fields should be assigned to their respective handler implementations only before
// RegisterDBService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type DBService struct {
	Size       func(context.Context, *SizeRequest) (*SizeReply, error)
	BucketSize func(context.Context, *BucketSizeRequest) (*BucketSizeReply, error)
}

func (s *DBService) size(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Size(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/remote.DB/Size",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Size(ctx, req.(*SizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *DBService) bucketSize(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.BucketSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/remote.DB/BucketSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.BucketSize(ctx, req.(*BucketSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterDBService registers a service implementation with a gRPC server.
func RegisterDBService(s grpc.ServiceRegistrar, srv *DBService) {
	srvCopy := *srv
	if srvCopy.Size == nil {
		srvCopy.Size = func(context.Context, *SizeRequest) (*SizeReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Size not implemented")
		}
	}
	if srvCopy.BucketSize == nil {
		srvCopy.BucketSize = func(context.Context, *BucketSizeRequest) (*BucketSizeReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method BucketSize not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "remote.DB",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Size",
				Handler:    srvCopy.size,
			},
			{
				MethodName: "BucketSize",
				Handler:    srvCopy.bucketSize,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "remote/db.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewDBService creates a new DBService containing the
// implemented methods of the DB service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewDBService(s interface{}) *DBService {
	ns := &DBService{}
	if h, ok := s.(interface {
		Size(context.Context, *SizeRequest) (*SizeReply, error)
	}); ok {
		ns.Size = h.Size
	}
	if h, ok := s.(interface {
		BucketSize(context.Context, *BucketSizeRequest) (*BucketSizeReply, error)
	}); ok {
		ns.BucketSize = h.BucketSize
	}
	return ns
}

// UnstableDBService is the service API for DB service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableDBService interface {
	Size(context.Context, *SizeRequest) (*SizeReply, error)
	BucketSize(context.Context, *BucketSizeRequest) (*BucketSizeReply, error)
}
