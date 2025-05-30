// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: homestay.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Homestay_HomestayDetail_FullMethodName = "/pb.homestay/homestayDetail"
	Homestay_HomestayList_FullMethodName   = "/pb.homestay/homestayList"
	Homestay_CreateHomestay_FullMethodName = "/pb.homestay/createHomestay"
)

// HomestayClient is the client API for Homestay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomestayClient interface {
	// homestayDetail
	HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error)
	HomestayList(ctx context.Context, in *HomestayListReq, opts ...grpc.CallOption) (*HomestayListResp, error)
	CreateHomestay(ctx context.Context, in *CreateHomestayReq, opts ...grpc.CallOption) (*Empty, error)
}

type homestayClient struct {
	cc grpc.ClientConnInterface
}

func NewHomestayClient(cc grpc.ClientConnInterface) HomestayClient {
	return &homestayClient{cc}
}

func (c *homestayClient) HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error) {
	out := new(HomestayDetailResp)
	err := c.cc.Invoke(ctx, Homestay_HomestayDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homestayClient) HomestayList(ctx context.Context, in *HomestayListReq, opts ...grpc.CallOption) (*HomestayListResp, error) {
	out := new(HomestayListResp)
	err := c.cc.Invoke(ctx, Homestay_HomestayList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homestayClient) CreateHomestay(ctx context.Context, in *CreateHomestayReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Homestay_CreateHomestay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomestayServer is the server API for Homestay service.
// All implementations must embed UnimplementedHomestayServer
// for forward compatibility
type HomestayServer interface {
	// homestayDetail
	HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error)
	HomestayList(context.Context, *HomestayListReq) (*HomestayListResp, error)
	CreateHomestay(context.Context, *CreateHomestayReq) (*Empty, error)
	mustEmbedUnimplementedHomestayServer()
}

// UnimplementedHomestayServer must be embedded to have forward compatible implementations.
type UnimplementedHomestayServer struct {
}

func (UnimplementedHomestayServer) HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomestayDetail not implemented")
}
func (UnimplementedHomestayServer) HomestayList(context.Context, *HomestayListReq) (*HomestayListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomestayList not implemented")
}
func (UnimplementedHomestayServer) CreateHomestay(context.Context, *CreateHomestayReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHomestay not implemented")
}
func (UnimplementedHomestayServer) mustEmbedUnimplementedHomestayServer() {}

// UnsafeHomestayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomestayServer will
// result in compilation errors.
type UnsafeHomestayServer interface {
	mustEmbedUnimplementedHomestayServer()
}

func RegisterHomestayServer(s grpc.ServiceRegistrar, srv HomestayServer) {
	s.RegisterService(&Homestay_ServiceDesc, srv)
}

func _Homestay_HomestayDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomestayDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomestayServer).HomestayDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Homestay_HomestayDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomestayServer).HomestayDetail(ctx, req.(*HomestayDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Homestay_HomestayList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomestayListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomestayServer).HomestayList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Homestay_HomestayList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomestayServer).HomestayList(ctx, req.(*HomestayListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Homestay_CreateHomestay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHomestayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomestayServer).CreateHomestay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Homestay_CreateHomestay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomestayServer).CreateHomestay(ctx, req.(*CreateHomestayReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Homestay_ServiceDesc is the grpc.ServiceDesc for Homestay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Homestay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.homestay",
	HandlerType: (*HomestayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "homestayDetail",
			Handler:    _Homestay_HomestayDetail_Handler,
		},
		{
			MethodName: "homestayList",
			Handler:    _Homestay_HomestayList_Handler,
		},
		{
			MethodName: "createHomestay",
			Handler:    _Homestay_CreateHomestay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "homestay.proto",
}
