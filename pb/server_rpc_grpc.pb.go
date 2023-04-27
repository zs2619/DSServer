// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: server_rpc.proto

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

// GameDscRealmClient is the client API for GameDscRealm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameDscRealmClient interface {
	CreateRealm(ctx context.Context, in *RpcCreateRealmReq, opts ...grpc.CallOption) (*RpcCreateRealmResp, error)
	JoinRealm(ctx context.Context, in *RpcJoinRealmReq, opts ...grpc.CallOption) (*RpcJoinRealmResp, error)
	QuickJoinRealm(ctx context.Context, in *RpcQuickJoinRealmReq, opts ...grpc.CallOption) (*RpcQuickJoinRealmResp, error)
	QueryRealmList(ctx context.Context, in *RpcQueryRealmListRealmReq, opts ...grpc.CallOption) (*RpcQueryRealmListRealmResp, error)
	NotifyJoinRealm(ctx context.Context, opts ...grpc.CallOption) (GameDscRealm_NotifyJoinRealmClient, error)
}

type gameDscRealmClient struct {
	cc grpc.ClientConnInterface
}

func NewGameDscRealmClient(cc grpc.ClientConnInterface) GameDscRealmClient {
	return &gameDscRealmClient{cc}
}

func (c *gameDscRealmClient) CreateRealm(ctx context.Context, in *RpcCreateRealmReq, opts ...grpc.CallOption) (*RpcCreateRealmResp, error) {
	out := new(RpcCreateRealmResp)
	err := c.cc.Invoke(ctx, "/pb.GameDscRealm/CreateRealm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameDscRealmClient) JoinRealm(ctx context.Context, in *RpcJoinRealmReq, opts ...grpc.CallOption) (*RpcJoinRealmResp, error) {
	out := new(RpcJoinRealmResp)
	err := c.cc.Invoke(ctx, "/pb.GameDscRealm/JoinRealm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameDscRealmClient) QuickJoinRealm(ctx context.Context, in *RpcQuickJoinRealmReq, opts ...grpc.CallOption) (*RpcQuickJoinRealmResp, error) {
	out := new(RpcQuickJoinRealmResp)
	err := c.cc.Invoke(ctx, "/pb.GameDscRealm/QuickJoinRealm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameDscRealmClient) QueryRealmList(ctx context.Context, in *RpcQueryRealmListRealmReq, opts ...grpc.CallOption) (*RpcQueryRealmListRealmResp, error) {
	out := new(RpcQueryRealmListRealmResp)
	err := c.cc.Invoke(ctx, "/pb.GameDscRealm/QueryRealmList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameDscRealmClient) NotifyJoinRealm(ctx context.Context, opts ...grpc.CallOption) (GameDscRealm_NotifyJoinRealmClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameDscRealm_ServiceDesc.Streams[0], "/pb.GameDscRealm/NotifyJoinRealm", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameDscRealmNotifyJoinRealmClient{stream}
	return x, nil
}

type GameDscRealm_NotifyJoinRealmClient interface {
	Send(*NotifyJoinRealmReq) error
	Recv() (*NotifyJoinRealmResp, error)
	grpc.ClientStream
}

type gameDscRealmNotifyJoinRealmClient struct {
	grpc.ClientStream
}

func (x *gameDscRealmNotifyJoinRealmClient) Send(m *NotifyJoinRealmReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameDscRealmNotifyJoinRealmClient) Recv() (*NotifyJoinRealmResp, error) {
	m := new(NotifyJoinRealmResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameDscRealmServer is the server API for GameDscRealm service.
// All implementations must embed UnimplementedGameDscRealmServer
// for forward compatibility
type GameDscRealmServer interface {
	CreateRealm(context.Context, *RpcCreateRealmReq) (*RpcCreateRealmResp, error)
	JoinRealm(context.Context, *RpcJoinRealmReq) (*RpcJoinRealmResp, error)
	QuickJoinRealm(context.Context, *RpcQuickJoinRealmReq) (*RpcQuickJoinRealmResp, error)
	QueryRealmList(context.Context, *RpcQueryRealmListRealmReq) (*RpcQueryRealmListRealmResp, error)
	NotifyJoinRealm(GameDscRealm_NotifyJoinRealmServer) error
	mustEmbedUnimplementedGameDscRealmServer()
}

// UnimplementedGameDscRealmServer must be embedded to have forward compatible implementations.
type UnimplementedGameDscRealmServer struct {
}

func (UnimplementedGameDscRealmServer) CreateRealm(context.Context, *RpcCreateRealmReq) (*RpcCreateRealmResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRealm not implemented")
}
func (UnimplementedGameDscRealmServer) JoinRealm(context.Context, *RpcJoinRealmReq) (*RpcJoinRealmResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRealm not implemented")
}
func (UnimplementedGameDscRealmServer) QuickJoinRealm(context.Context, *RpcQuickJoinRealmReq) (*RpcQuickJoinRealmResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuickJoinRealm not implemented")
}
func (UnimplementedGameDscRealmServer) QueryRealmList(context.Context, *RpcQueryRealmListRealmReq) (*RpcQueryRealmListRealmResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRealmList not implemented")
}
func (UnimplementedGameDscRealmServer) NotifyJoinRealm(GameDscRealm_NotifyJoinRealmServer) error {
	return status.Errorf(codes.Unimplemented, "method NotifyJoinRealm not implemented")
}
func (UnimplementedGameDscRealmServer) mustEmbedUnimplementedGameDscRealmServer() {}

// UnsafeGameDscRealmServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameDscRealmServer will
// result in compilation errors.
type UnsafeGameDscRealmServer interface {
	mustEmbedUnimplementedGameDscRealmServer()
}

func RegisterGameDscRealmServer(s grpc.ServiceRegistrar, srv GameDscRealmServer) {
	s.RegisterService(&GameDscRealm_ServiceDesc, srv)
}

func _GameDscRealm_CreateRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcCreateRealmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameDscRealmServer).CreateRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameDscRealm/CreateRealm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameDscRealmServer).CreateRealm(ctx, req.(*RpcCreateRealmReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameDscRealm_JoinRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcJoinRealmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameDscRealmServer).JoinRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameDscRealm/JoinRealm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameDscRealmServer).JoinRealm(ctx, req.(*RpcJoinRealmReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameDscRealm_QuickJoinRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcQuickJoinRealmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameDscRealmServer).QuickJoinRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameDscRealm/QuickJoinRealm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameDscRealmServer).QuickJoinRealm(ctx, req.(*RpcQuickJoinRealmReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameDscRealm_QueryRealmList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcQueryRealmListRealmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameDscRealmServer).QueryRealmList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameDscRealm/QueryRealmList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameDscRealmServer).QueryRealmList(ctx, req.(*RpcQueryRealmListRealmReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameDscRealm_NotifyJoinRealm_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameDscRealmServer).NotifyJoinRealm(&gameDscRealmNotifyJoinRealmServer{stream})
}

type GameDscRealm_NotifyJoinRealmServer interface {
	Send(*NotifyJoinRealmResp) error
	Recv() (*NotifyJoinRealmReq, error)
	grpc.ServerStream
}

type gameDscRealmNotifyJoinRealmServer struct {
	grpc.ServerStream
}

func (x *gameDscRealmNotifyJoinRealmServer) Send(m *NotifyJoinRealmResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameDscRealmNotifyJoinRealmServer) Recv() (*NotifyJoinRealmReq, error) {
	m := new(NotifyJoinRealmReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameDscRealm_ServiceDesc is the grpc.ServiceDesc for GameDscRealm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameDscRealm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GameDscRealm",
	HandlerType: (*GameDscRealmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRealm",
			Handler:    _GameDscRealm_CreateRealm_Handler,
		},
		{
			MethodName: "JoinRealm",
			Handler:    _GameDscRealm_JoinRealm_Handler,
		},
		{
			MethodName: "QuickJoinRealm",
			Handler:    _GameDscRealm_QuickJoinRealm_Handler,
		},
		{
			MethodName: "QueryRealmList",
			Handler:    _GameDscRealm_QueryRealmList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NotifyJoinRealm",
			Handler:       _GameDscRealm_NotifyJoinRealm_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server_rpc.proto",
}

// DsaDscARealmClient is the client API for DsaDscARealm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DsaDscARealmClient interface {
	WaitCreateRealm(ctx context.Context, opts ...grpc.CallOption) (DsaDscARealm_WaitCreateRealmClient, error)
	WaitJoinRealm(ctx context.Context, opts ...grpc.CallOption) (DsaDscARealm_WaitJoinRealmClient, error)
}

type dsaDscARealmClient struct {
	cc grpc.ClientConnInterface
}

func NewDsaDscARealmClient(cc grpc.ClientConnInterface) DsaDscARealmClient {
	return &dsaDscARealmClient{cc}
}

func (c *dsaDscARealmClient) WaitCreateRealm(ctx context.Context, opts ...grpc.CallOption) (DsaDscARealm_WaitCreateRealmClient, error) {
	stream, err := c.cc.NewStream(ctx, &DsaDscARealm_ServiceDesc.Streams[0], "/pb.DsaDscARealm/WaitCreateRealm", opts...)
	if err != nil {
		return nil, err
	}
	x := &dsaDscARealmWaitCreateRealmClient{stream}
	return x, nil
}

type DsaDscARealm_WaitCreateRealmClient interface {
	Send(*RpcCreateRealmResult) error
	Recv() (*RpcCreateRealmInfo, error)
	grpc.ClientStream
}

type dsaDscARealmWaitCreateRealmClient struct {
	grpc.ClientStream
}

func (x *dsaDscARealmWaitCreateRealmClient) Send(m *RpcCreateRealmResult) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dsaDscARealmWaitCreateRealmClient) Recv() (*RpcCreateRealmInfo, error) {
	m := new(RpcCreateRealmInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dsaDscARealmClient) WaitJoinRealm(ctx context.Context, opts ...grpc.CallOption) (DsaDscARealm_WaitJoinRealmClient, error) {
	stream, err := c.cc.NewStream(ctx, &DsaDscARealm_ServiceDesc.Streams[1], "/pb.DsaDscARealm/WaitJoinRealm", opts...)
	if err != nil {
		return nil, err
	}
	x := &dsaDscARealmWaitJoinRealmClient{stream}
	return x, nil
}

type DsaDscARealm_WaitJoinRealmClient interface {
	Send(*RpcJoinRealmResult) error
	Recv() (*RpcJoinRealmInfo, error)
	grpc.ClientStream
}

type dsaDscARealmWaitJoinRealmClient struct {
	grpc.ClientStream
}

func (x *dsaDscARealmWaitJoinRealmClient) Send(m *RpcJoinRealmResult) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dsaDscARealmWaitJoinRealmClient) Recv() (*RpcJoinRealmInfo, error) {
	m := new(RpcJoinRealmInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DsaDscARealmServer is the server API for DsaDscARealm service.
// All implementations must embed UnimplementedDsaDscARealmServer
// for forward compatibility
type DsaDscARealmServer interface {
	WaitCreateRealm(DsaDscARealm_WaitCreateRealmServer) error
	WaitJoinRealm(DsaDscARealm_WaitJoinRealmServer) error
	mustEmbedUnimplementedDsaDscARealmServer()
}

// UnimplementedDsaDscARealmServer must be embedded to have forward compatible implementations.
type UnimplementedDsaDscARealmServer struct {
}

func (UnimplementedDsaDscARealmServer) WaitCreateRealm(DsaDscARealm_WaitCreateRealmServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitCreateRealm not implemented")
}
func (UnimplementedDsaDscARealmServer) WaitJoinRealm(DsaDscARealm_WaitJoinRealmServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitJoinRealm not implemented")
}
func (UnimplementedDsaDscARealmServer) mustEmbedUnimplementedDsaDscARealmServer() {}

// UnsafeDsaDscARealmServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DsaDscARealmServer will
// result in compilation errors.
type UnsafeDsaDscARealmServer interface {
	mustEmbedUnimplementedDsaDscARealmServer()
}

func RegisterDsaDscARealmServer(s grpc.ServiceRegistrar, srv DsaDscARealmServer) {
	s.RegisterService(&DsaDscARealm_ServiceDesc, srv)
}

func _DsaDscARealm_WaitCreateRealm_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DsaDscARealmServer).WaitCreateRealm(&dsaDscARealmWaitCreateRealmServer{stream})
}

type DsaDscARealm_WaitCreateRealmServer interface {
	Send(*RpcCreateRealmInfo) error
	Recv() (*RpcCreateRealmResult, error)
	grpc.ServerStream
}

type dsaDscARealmWaitCreateRealmServer struct {
	grpc.ServerStream
}

func (x *dsaDscARealmWaitCreateRealmServer) Send(m *RpcCreateRealmInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dsaDscARealmWaitCreateRealmServer) Recv() (*RpcCreateRealmResult, error) {
	m := new(RpcCreateRealmResult)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DsaDscARealm_WaitJoinRealm_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DsaDscARealmServer).WaitJoinRealm(&dsaDscARealmWaitJoinRealmServer{stream})
}

type DsaDscARealm_WaitJoinRealmServer interface {
	Send(*RpcJoinRealmInfo) error
	Recv() (*RpcJoinRealmResult, error)
	grpc.ServerStream
}

type dsaDscARealmWaitJoinRealmServer struct {
	grpc.ServerStream
}

func (x *dsaDscARealmWaitJoinRealmServer) Send(m *RpcJoinRealmInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dsaDscARealmWaitJoinRealmServer) Recv() (*RpcJoinRealmResult, error) {
	m := new(RpcJoinRealmResult)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DsaDscARealm_ServiceDesc is the grpc.ServiceDesc for DsaDscARealm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DsaDscARealm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DsaDscARealm",
	HandlerType: (*DsaDscARealmServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitCreateRealm",
			Handler:       _DsaDscARealm_WaitCreateRealm_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "WaitJoinRealm",
			Handler:       _DsaDscARealm_WaitJoinRealm_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server_rpc.proto",
}
