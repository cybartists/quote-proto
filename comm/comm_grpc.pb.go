// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/comm.proto

package comm

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	entity "quote-proto/entity"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClient interface {
	//推送逐笔成交行情数据
	NewTickRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewTickRecordStreamClient, error)
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) NewTickRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewTickRecordStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[0], "/sa.rpc.comm.Proxy/NewTickRecordStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyNewTickRecordStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_NewTickRecordStreamClient interface {
	Recv() (*entity.Transaction, error)
	grpc.ClientStream
}

type proxyNewTickRecordStreamClient struct {
	grpc.ClientStream
}

func (x *proxyNewTickRecordStreamClient) Recv() (*entity.Transaction, error) {
	m := new(entity.Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServer is the server API for Proxy service.
// All implementations must embed UnimplementedProxyServer
// for forward compatibility
type ProxyServer interface {
	//推送逐笔成交行情数据
	NewTickRecordStream(*entity.Void, Proxy_NewTickRecordStreamServer) error
	mustEmbedUnimplementedProxyServer()
}

// UnimplementedProxyServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (UnimplementedProxyServer) NewTickRecordStream(*entity.Void, Proxy_NewTickRecordStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NewTickRecordStream not implemented")
}
func (UnimplementedProxyServer) mustEmbedUnimplementedProxyServer() {}

// UnsafeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServer will
// result in compilation errors.
type UnsafeProxyServer interface {
	mustEmbedUnimplementedProxyServer()
}

func RegisterProxyServer(s grpc.ServiceRegistrar, srv ProxyServer) {
	s.RegisterService(&Proxy_ServiceDesc, srv)
}

func _Proxy_NewTickRecordStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(entity.Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).NewTickRecordStream(m, &proxyNewTickRecordStreamServer{stream})
}

type Proxy_NewTickRecordStreamServer interface {
	Send(*entity.Transaction) error
	grpc.ServerStream
}

type proxyNewTickRecordStreamServer struct {
	grpc.ServerStream
}

func (x *proxyNewTickRecordStreamServer) Send(m *entity.Transaction) error {
	return x.ServerStream.SendMsg(m)
}

// Proxy_ServiceDesc is the grpc.ServiceDesc for Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sa.rpc.comm.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewTickRecordStream",
			Handler:       _Proxy_NewTickRecordStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/comm.proto",
}
