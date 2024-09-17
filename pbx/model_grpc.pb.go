// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: model.proto

package pbx

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
	Node_MessageLoop_FullMethodName = "/pbx.Node/MessageLoop"
)

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	// Client sends a stream of ClientMsg, server responds with a stream of ServerMsg
	MessageLoop(ctx context.Context, opts ...grpc.CallOption) (Node_MessageLoopClient, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) MessageLoop(ctx context.Context, opts ...grpc.CallOption) (Node_MessageLoopClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], Node_MessageLoop_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeMessageLoopClient{stream}
	return x, nil
}

type Node_MessageLoopClient interface {
	Send(*ClientMsg) error
	Recv() (*ServerMsg, error)
	grpc.ClientStream
}

type nodeMessageLoopClient struct {
	grpc.ClientStream
}

func (x *nodeMessageLoopClient) Send(m *ClientMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeMessageLoopClient) Recv() (*ServerMsg, error) {
	m := new(ServerMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	// Client sends a stream of ClientMsg, server responds with a stream of ServerMsg
	MessageLoop(Node_MessageLoopServer) error
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) MessageLoop(Node_MessageLoopServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageLoop not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_MessageLoop_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).MessageLoop(&nodeMessageLoopServer{stream})
}

type Node_MessageLoopServer interface {
	Send(*ServerMsg) error
	Recv() (*ClientMsg, error)
	grpc.ServerStream
}

type nodeMessageLoopServer struct {
	grpc.ServerStream
}

func (x *nodeMessageLoopServer) Send(m *ServerMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeMessageLoopServer) Recv() (*ClientMsg, error) {
	m := new(ClientMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbx.Node",
	HandlerType: (*NodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessageLoop",
			Handler:       _Node_MessageLoop_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "model.proto",
}

const (
	Plugin_FireHose_FullMethodName     = "/pbx.Plugin/FireHose"
	Plugin_Find_FullMethodName         = "/pbx.Plugin/Find"
	Plugin_Account_FullMethodName      = "/pbx.Plugin/Account"
	Plugin_Topic_FullMethodName        = "/pbx.Plugin/Topic"
	Plugin_Subscription_FullMethodName = "/pbx.Plugin/Subscription"
	Plugin_Message_FullMethodName      = "/pbx.Plugin/Message"
)

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	// This plugin method is called by Tinode server for every message received from the clients. The
	// method returns a ServerCtrl message. Non-zero ServerCtrl.code indicates that no further
	// processing is needed. The Tinode server will generate a {ctrl} message from the returned ServerCtrl
	// and forward it to the client session.
	// ServerCtrl.code equals to 0 instructs the server to continue with default processing of the client message.
	FireHose(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*ServerResp, error)
	// An alteranative user and topic discovery mechanism.
	// A search request issued on a 'fnd' topic. This method is called to generate an alternative result set.
	Find(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchFound, error)
	// Account created, updated or deleted
	Account(ctx context.Context, in *AccountEvent, opts ...grpc.CallOption) (*Unused, error)
	// Topic created, updated [or deleted -- not supported yet]
	Topic(ctx context.Context, in *TopicEvent, opts ...grpc.CallOption) (*Unused, error)
	// Subscription created, updated or deleted
	Subscription(ctx context.Context, in *SubscriptionEvent, opts ...grpc.CallOption) (*Unused, error)
	// Message published or deleted
	Message(ctx context.Context, in *MessageEvent, opts ...grpc.CallOption) (*Unused, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) FireHose(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*ServerResp, error) {
	out := new(ServerResp)
	err := c.cc.Invoke(ctx, Plugin_FireHose_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Find(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchFound, error) {
	out := new(SearchFound)
	err := c.cc.Invoke(ctx, Plugin_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Account(ctx context.Context, in *AccountEvent, opts ...grpc.CallOption) (*Unused, error) {
	out := new(Unused)
	err := c.cc.Invoke(ctx, Plugin_Account_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Topic(ctx context.Context, in *TopicEvent, opts ...grpc.CallOption) (*Unused, error) {
	out := new(Unused)
	err := c.cc.Invoke(ctx, Plugin_Topic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Subscription(ctx context.Context, in *SubscriptionEvent, opts ...grpc.CallOption) (*Unused, error) {
	out := new(Unused)
	err := c.cc.Invoke(ctx, Plugin_Subscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Message(ctx context.Context, in *MessageEvent, opts ...grpc.CallOption) (*Unused, error) {
	out := new(Unused)
	err := c.cc.Invoke(ctx, Plugin_Message_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	// This plugin method is called by Tinode server for every message received from the clients. The
	// method returns a ServerCtrl message. Non-zero ServerCtrl.code indicates that no further
	// processing is needed. The Tinode server will generate a {ctrl} message from the returned ServerCtrl
	// and forward it to the client session.
	// ServerCtrl.code equals to 0 instructs the server to continue with default processing of the client message.
	FireHose(context.Context, *ClientReq) (*ServerResp, error)
	// An alteranative user and topic discovery mechanism.
	// A search request issued on a 'fnd' topic. This method is called to generate an alternative result set.
	Find(context.Context, *SearchQuery) (*SearchFound, error)
	// Account created, updated or deleted
	Account(context.Context, *AccountEvent) (*Unused, error)
	// Topic created, updated [or deleted -- not supported yet]
	Topic(context.Context, *TopicEvent) (*Unused, error)
	// Subscription created, updated or deleted
	Subscription(context.Context, *SubscriptionEvent) (*Unused, error)
	// Message published or deleted
	Message(context.Context, *MessageEvent) (*Unused, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) FireHose(context.Context, *ClientReq) (*ServerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FireHose not implemented")
}
func (UnimplementedPluginServer) Find(context.Context, *SearchQuery) (*SearchFound, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedPluginServer) Account(context.Context, *AccountEvent) (*Unused, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (UnimplementedPluginServer) Topic(context.Context, *TopicEvent) (*Unused, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topic not implemented")
}
func (UnimplementedPluginServer) Subscription(context.Context, *SubscriptionEvent) (*Unused, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscription not implemented")
}
func (UnimplementedPluginServer) Message(context.Context, *MessageEvent) (*Unused, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_FireHose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).FireHose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_FireHose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).FireHose(ctx, req.(*ClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Find(ctx, req.(*SearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Account_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Account(ctx, req.(*AccountEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Topic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Topic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Topic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Topic(ctx, req.(*TopicEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Subscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Subscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Subscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Subscription(ctx, req.(*SubscriptionEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Message_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Message(ctx, req.(*MessageEvent))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbx.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FireHose",
			Handler:    _Plugin_FireHose_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Plugin_Find_Handler,
		},
		{
			MethodName: "Account",
			Handler:    _Plugin_Account_Handler,
		},
		{
			MethodName: "Topic",
			Handler:    _Plugin_Topic_Handler,
		},
		{
			MethodName: "Subscription",
			Handler:    _Plugin_Subscription_Handler,
		},
		{
			MethodName: "Message",
			Handler:    _Plugin_Message_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}
