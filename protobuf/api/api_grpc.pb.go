// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api.proto

package api

import (
	context "context"
	eth "github.com/chainbound/fiber-go/protobuf/eth"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	SubscribeNewTxs(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (API_SubscribeNewTxsClient, error)
	SendTransaction(ctx context.Context, opts ...grpc.CallOption) (API_SendTransactionClient, error)
	SendRawTransaction(ctx context.Context, opts ...grpc.CallOption) (API_SendRawTransactionClient, error)
	SendTransactionSequence(ctx context.Context, opts ...grpc.CallOption) (API_SendTransactionSequenceClient, error)
	SendRawTransactionSequence(ctx context.Context, opts ...grpc.CallOption) (API_SendRawTransactionSequenceClient, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) SubscribeNewTxs(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (API_SubscribeNewTxsClient, error) {
	stream, err := c.cc.NewStream(ctx, &API_ServiceDesc.Streams[0], "/api.API/SubscribeNewTxs", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISubscribeNewTxsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_SubscribeNewTxsClient interface {
	Recv() (*eth.Transaction, error)
	grpc.ClientStream
}

type aPISubscribeNewTxsClient struct {
	grpc.ClientStream
}

func (x *aPISubscribeNewTxsClient) Recv() (*eth.Transaction, error) {
	m := new(eth.Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) SendTransaction(ctx context.Context, opts ...grpc.CallOption) (API_SendTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &API_ServiceDesc.Streams[1], "/api.API/SendTransaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISendTransactionClient{stream}
	return x, nil
}

type API_SendTransactionClient interface {
	Send(*eth.Transaction) error
	Recv() (*TransactionResponse, error)
	grpc.ClientStream
}

type aPISendTransactionClient struct {
	grpc.ClientStream
}

func (x *aPISendTransactionClient) Send(m *eth.Transaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPISendTransactionClient) Recv() (*TransactionResponse, error) {
	m := new(TransactionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) SendRawTransaction(ctx context.Context, opts ...grpc.CallOption) (API_SendRawTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &API_ServiceDesc.Streams[2], "/api.API/SendRawTransaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISendRawTransactionClient{stream}
	return x, nil
}

type API_SendRawTransactionClient interface {
	Send(*RawTxMsg) error
	Recv() (*TransactionResponse, error)
	grpc.ClientStream
}

type aPISendRawTransactionClient struct {
	grpc.ClientStream
}

func (x *aPISendRawTransactionClient) Send(m *RawTxMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPISendRawTransactionClient) Recv() (*TransactionResponse, error) {
	m := new(TransactionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) SendTransactionSequence(ctx context.Context, opts ...grpc.CallOption) (API_SendTransactionSequenceClient, error) {
	stream, err := c.cc.NewStream(ctx, &API_ServiceDesc.Streams[3], "/api.API/SendTransactionSequence", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISendTransactionSequenceClient{stream}
	return x, nil
}

type API_SendTransactionSequenceClient interface {
	Send(*TxSequenceMsg) error
	Recv() (*TxSequenceResponse, error)
	grpc.ClientStream
}

type aPISendTransactionSequenceClient struct {
	grpc.ClientStream
}

func (x *aPISendTransactionSequenceClient) Send(m *TxSequenceMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPISendTransactionSequenceClient) Recv() (*TxSequenceResponse, error) {
	m := new(TxSequenceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) SendRawTransactionSequence(ctx context.Context, opts ...grpc.CallOption) (API_SendRawTransactionSequenceClient, error) {
	stream, err := c.cc.NewStream(ctx, &API_ServiceDesc.Streams[4], "/api.API/SendRawTransactionSequence", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISendRawTransactionSequenceClient{stream}
	return x, nil
}

type API_SendRawTransactionSequenceClient interface {
	Send(*RawTxSequenceMsg) error
	Recv() (*TxSequenceResponse, error)
	grpc.ClientStream
}

type aPISendRawTransactionSequenceClient struct {
	grpc.ClientStream
}

func (x *aPISendRawTransactionSequenceClient) Send(m *RawTxSequenceMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPISendRawTransactionSequenceClient) Recv() (*TxSequenceResponse, error) {
	m := new(TxSequenceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	SubscribeNewTxs(*TxFilter, API_SubscribeNewTxsServer) error
	SendTransaction(API_SendTransactionServer) error
	SendRawTransaction(API_SendRawTransactionServer) error
	SendTransactionSequence(API_SendTransactionSequenceServer) error
	SendRawTransactionSequence(API_SendRawTransactionSequenceServer) error
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) SubscribeNewTxs(*TxFilter, API_SubscribeNewTxsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeNewTxs not implemented")
}
func (UnimplementedAPIServer) SendTransaction(API_SendTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (UnimplementedAPIServer) SendRawTransaction(API_SendRawTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method SendRawTransaction not implemented")
}
func (UnimplementedAPIServer) SendTransactionSequence(API_SendTransactionSequenceServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTransactionSequence not implemented")
}
func (UnimplementedAPIServer) SendRawTransactionSequence(API_SendRawTransactionSequenceServer) error {
	return status.Errorf(codes.Unimplemented, "method SendRawTransactionSequence not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_SubscribeNewTxs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TxFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).SubscribeNewTxs(m, &aPISubscribeNewTxsServer{stream})
}

type API_SubscribeNewTxsServer interface {
	Send(*eth.Transaction) error
	grpc.ServerStream
}

type aPISubscribeNewTxsServer struct {
	grpc.ServerStream
}

func (x *aPISubscribeNewTxsServer) Send(m *eth.Transaction) error {
	return x.ServerStream.SendMsg(m)
}

func _API_SendTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).SendTransaction(&aPISendTransactionServer{stream})
}

type API_SendTransactionServer interface {
	Send(*TransactionResponse) error
	Recv() (*eth.Transaction, error)
	grpc.ServerStream
}

type aPISendTransactionServer struct {
	grpc.ServerStream
}

func (x *aPISendTransactionServer) Send(m *TransactionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPISendTransactionServer) Recv() (*eth.Transaction, error) {
	m := new(eth.Transaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _API_SendRawTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).SendRawTransaction(&aPISendRawTransactionServer{stream})
}

type API_SendRawTransactionServer interface {
	Send(*TransactionResponse) error
	Recv() (*RawTxMsg, error)
	grpc.ServerStream
}

type aPISendRawTransactionServer struct {
	grpc.ServerStream
}

func (x *aPISendRawTransactionServer) Send(m *TransactionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPISendRawTransactionServer) Recv() (*RawTxMsg, error) {
	m := new(RawTxMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _API_SendTransactionSequence_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).SendTransactionSequence(&aPISendTransactionSequenceServer{stream})
}

type API_SendTransactionSequenceServer interface {
	Send(*TxSequenceResponse) error
	Recv() (*TxSequenceMsg, error)
	grpc.ServerStream
}

type aPISendTransactionSequenceServer struct {
	grpc.ServerStream
}

func (x *aPISendTransactionSequenceServer) Send(m *TxSequenceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPISendTransactionSequenceServer) Recv() (*TxSequenceMsg, error) {
	m := new(TxSequenceMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _API_SendRawTransactionSequence_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).SendRawTransactionSequence(&aPISendRawTransactionSequenceServer{stream})
}

type API_SendRawTransactionSequenceServer interface {
	Send(*TxSequenceResponse) error
	Recv() (*RawTxSequenceMsg, error)
	grpc.ServerStream
}

type aPISendRawTransactionSequenceServer struct {
	grpc.ServerStream
}

func (x *aPISendRawTransactionSequenceServer) Send(m *TxSequenceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPISendRawTransactionSequenceServer) Recv() (*RawTxSequenceMsg, error) {
	m := new(RawTxSequenceMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.API",
	HandlerType: (*APIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeNewTxs",
			Handler:       _API_SubscribeNewTxs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendTransaction",
			Handler:       _API_SendTransaction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendRawTransaction",
			Handler:       _API_SendRawTransaction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendTransactionSequence",
			Handler:       _API_SendTransactionSequence_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendRawTransactionSequence",
			Handler:       _API_SendRawTransactionSequence_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
