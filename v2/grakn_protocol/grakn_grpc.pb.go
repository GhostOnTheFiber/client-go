// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grakn_protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GraknClient is the client API for Grakn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraknClient interface {
	DatabaseContains(ctx context.Context, in *Database_Contains_Req, opts ...grpc.CallOption) (*Database_Contains_Res, error)
	DatabaseCreate(ctx context.Context, in *Database_Create_Req, opts ...grpc.CallOption) (*Database_Create_Res, error)
	DatabaseAll(ctx context.Context, in *Database_All_Req, opts ...grpc.CallOption) (*Database_All_Res, error)
	DatabaseDelete(ctx context.Context, in *Database_Delete_Req, opts ...grpc.CallOption) (*Database_Delete_Res, error)
	SessionOpen(ctx context.Context, in *Session_Open_Req, opts ...grpc.CallOption) (*Session_Open_Res, error)
	SessionClose(ctx context.Context, in *Session_Close_Req, opts ...grpc.CallOption) (*Session_Close_Res, error)
	// Opens a bi-directional stream representing a stateful transaction, streaming
	// requests and responses back-and-forth. The first request message must
	// be {Transaction.Open.Req}. Closing the stream closes the transaction.
	Transaction(ctx context.Context, opts ...grpc.CallOption) (Grakn_TransactionClient, error)
}

type graknClient struct {
	cc grpc.ClientConnInterface
}

func NewGraknClient(cc grpc.ClientConnInterface) GraknClient {
	return &graknClient{cc}
}

func (c *graknClient) DatabaseContains(ctx context.Context, in *Database_Contains_Req, opts ...grpc.CallOption) (*Database_Contains_Res, error) {
	out := new(Database_Contains_Res)
	err := c.cc.Invoke(ctx, "/grakn.protocol.Grakn/database_contains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graknClient) DatabaseCreate(ctx context.Context, in *Database_Create_Req, opts ...grpc.CallOption) (*Database_Create_Res, error) {
	out := new(Database_Create_Res)
	err := c.cc.Invoke(ctx, "/grakn.protocol.Grakn/database_create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graknClient) DatabaseAll(ctx context.Context, in *Database_All_Req, opts ...grpc.CallOption) (*Database_All_Res, error) {
	out := new(Database_All_Res)
	err := c.cc.Invoke(ctx, "/grakn.protocol.Grakn/database_all", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graknClient) DatabaseDelete(ctx context.Context, in *Database_Delete_Req, opts ...grpc.CallOption) (*Database_Delete_Res, error) {
	out := new(Database_Delete_Res)
	err := c.cc.Invoke(ctx, "/grakn.protocol.Grakn/database_delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graknClient) SessionOpen(ctx context.Context, in *Session_Open_Req, opts ...grpc.CallOption) (*Session_Open_Res, error) {
	out := new(Session_Open_Res)
	err := c.cc.Invoke(ctx, "/grakn.protocol.Grakn/session_open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graknClient) SessionClose(ctx context.Context, in *Session_Close_Req, opts ...grpc.CallOption) (*Session_Close_Res, error) {
	out := new(Session_Close_Res)
	err := c.cc.Invoke(ctx, "/grakn.protocol.Grakn/session_close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graknClient) Transaction(ctx context.Context, opts ...grpc.CallOption) (Grakn_TransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Grakn_serviceDesc.Streams[0], "/grakn.protocol.Grakn/transaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &graknTransactionClient{stream}
	return x, nil
}

type Grakn_TransactionClient interface {
	Send(*Transaction_Req) error
	Recv() (*Transaction_Res, error)
	grpc.ClientStream
}

type graknTransactionClient struct {
	grpc.ClientStream
}

func (x *graknTransactionClient) Send(m *Transaction_Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *graknTransactionClient) Recv() (*Transaction_Res, error) {
	m := new(Transaction_Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GraknServer is the server API for Grakn service.
// All implementations must embed UnimplementedGraknServer
// for forward compatibility
type GraknServer interface {
	DatabaseContains(context.Context, *Database_Contains_Req) (*Database_Contains_Res, error)
	DatabaseCreate(context.Context, *Database_Create_Req) (*Database_Create_Res, error)
	DatabaseAll(context.Context, *Database_All_Req) (*Database_All_Res, error)
	DatabaseDelete(context.Context, *Database_Delete_Req) (*Database_Delete_Res, error)
	SessionOpen(context.Context, *Session_Open_Req) (*Session_Open_Res, error)
	SessionClose(context.Context, *Session_Close_Req) (*Session_Close_Res, error)
	// Opens a bi-directional stream representing a stateful transaction, streaming
	// requests and responses back-and-forth. The first request message must
	// be {Transaction.Open.Req}. Closing the stream closes the transaction.
	Transaction(Grakn_TransactionServer) error
	mustEmbedUnimplementedGraknServer()
}

// UnimplementedGraknServer must be embedded to have forward compatible implementations.
type UnimplementedGraknServer struct {
}

func (UnimplementedGraknServer) DatabaseContains(context.Context, *Database_Contains_Req) (*Database_Contains_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseContains not implemented")
}
func (UnimplementedGraknServer) DatabaseCreate(context.Context, *Database_Create_Req) (*Database_Create_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseCreate not implemented")
}
func (UnimplementedGraknServer) DatabaseAll(context.Context, *Database_All_Req) (*Database_All_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseAll not implemented")
}
func (UnimplementedGraknServer) DatabaseDelete(context.Context, *Database_Delete_Req) (*Database_Delete_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseDelete not implemented")
}
func (UnimplementedGraknServer) SessionOpen(context.Context, *Session_Open_Req) (*Session_Open_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionOpen not implemented")
}
func (UnimplementedGraknServer) SessionClose(context.Context, *Session_Close_Req) (*Session_Close_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionClose not implemented")
}
func (UnimplementedGraknServer) Transaction(Grakn_TransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method Transaction not implemented")
}
func (UnimplementedGraknServer) mustEmbedUnimplementedGraknServer() {}

// UnsafeGraknServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraknServer will
// result in compilation errors.
type UnsafeGraknServer interface {
	mustEmbedUnimplementedGraknServer()
}

func RegisterGraknServer(s grpc.ServiceRegistrar, srv GraknServer) {
	s.RegisterService(&_Grakn_serviceDesc, srv)
}

func _Grakn_DatabaseContains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Database_Contains_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraknServer).DatabaseContains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grakn.protocol.Grakn/database_contains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraknServer).DatabaseContains(ctx, req.(*Database_Contains_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grakn_DatabaseCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Database_Create_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraknServer).DatabaseCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grakn.protocol.Grakn/database_create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraknServer).DatabaseCreate(ctx, req.(*Database_Create_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grakn_DatabaseAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Database_All_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraknServer).DatabaseAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grakn.protocol.Grakn/database_all",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraknServer).DatabaseAll(ctx, req.(*Database_All_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grakn_DatabaseDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Database_Delete_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraknServer).DatabaseDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grakn.protocol.Grakn/database_delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraknServer).DatabaseDelete(ctx, req.(*Database_Delete_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grakn_SessionOpen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session_Open_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraknServer).SessionOpen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grakn.protocol.Grakn/session_open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraknServer).SessionOpen(ctx, req.(*Session_Open_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grakn_SessionClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session_Close_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraknServer).SessionClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grakn.protocol.Grakn/session_close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraknServer).SessionClose(ctx, req.(*Session_Close_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grakn_Transaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GraknServer).Transaction(&graknTransactionServer{stream})
}

type Grakn_TransactionServer interface {
	Send(*Transaction_Res) error
	Recv() (*Transaction_Req, error)
	grpc.ServerStream
}

type graknTransactionServer struct {
	grpc.ServerStream
}

func (x *graknTransactionServer) Send(m *Transaction_Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *graknTransactionServer) Recv() (*Transaction_Req, error) {
	m := new(Transaction_Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Grakn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grakn.protocol.Grakn",
	HandlerType: (*GraknServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "database_contains",
			Handler:    _Grakn_DatabaseContains_Handler,
		},
		{
			MethodName: "database_create",
			Handler:    _Grakn_DatabaseCreate_Handler,
		},
		{
			MethodName: "database_all",
			Handler:    _Grakn_DatabaseAll_Handler,
		},
		{
			MethodName: "database_delete",
			Handler:    _Grakn_DatabaseDelete_Handler,
		},
		{
			MethodName: "session_open",
			Handler:    _Grakn_SessionOpen_Handler,
		},
		{
			MethodName: "session_close",
			Handler:    _Grakn_SessionClose_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "transaction",
			Handler:       _Grakn_Transaction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/grakn.proto",
}