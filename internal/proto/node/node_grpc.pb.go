// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: node.proto

package node

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
	NodeCommunication_Ping_FullMethodName          = "/NodeCommunication/Ping"
	NodeCommunication_GetKnownNodes_FullMethodName = "/NodeCommunication/GetKnownNodes"
)

// NodeCommunicationClient is the client API for NodeCommunication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeCommunicationClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResult, error)
	GetKnownNodes(ctx context.Context, in *KnownNodes, opts ...grpc.CallOption) (*KnownNodes, error)
}

type nodeCommunicationClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeCommunicationClient(cc grpc.ClientConnInterface) NodeCommunicationClient {
	return &nodeCommunicationClient{cc}
}

func (c *nodeCommunicationClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResult, error) {
	out := new(PingResult)
	err := c.cc.Invoke(ctx, NodeCommunication_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationClient) GetKnownNodes(ctx context.Context, in *KnownNodes, opts ...grpc.CallOption) (*KnownNodes, error) {
	out := new(KnownNodes)
	err := c.cc.Invoke(ctx, NodeCommunication_GetKnownNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeCommunicationServer is the server API for NodeCommunication service.
// All implementations must embed UnimplementedNodeCommunicationServer
// for forward compatibility
type NodeCommunicationServer interface {
	Ping(context.Context, *PingRequest) (*PingResult, error)
	GetKnownNodes(context.Context, *KnownNodes) (*KnownNodes, error)
	mustEmbedUnimplementedNodeCommunicationServer()
}

// UnimplementedNodeCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedNodeCommunicationServer struct {
}

func (UnimplementedNodeCommunicationServer) Ping(context.Context, *PingRequest) (*PingResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedNodeCommunicationServer) GetKnownNodes(context.Context, *KnownNodes) (*KnownNodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnownNodes not implemented")
}
func (UnimplementedNodeCommunicationServer) mustEmbedUnimplementedNodeCommunicationServer() {}

// UnsafeNodeCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeCommunicationServer will
// result in compilation errors.
type UnsafeNodeCommunicationServer interface {
	mustEmbedUnimplementedNodeCommunicationServer()
}

func RegisterNodeCommunicationServer(s grpc.ServiceRegistrar, srv NodeCommunicationServer) {
	s.RegisterService(&NodeCommunication_ServiceDesc, srv)
}

func _NodeCommunication_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeCommunication_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunication_GetKnownNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KnownNodes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationServer).GetKnownNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeCommunication_GetKnownNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationServer).GetKnownNodes(ctx, req.(*KnownNodes))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeCommunication_ServiceDesc is the grpc.ServiceDesc for NodeCommunication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeCommunication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NodeCommunication",
	HandlerType: (*NodeCommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _NodeCommunication_Ping_Handler,
		},
		{
			MethodName: "GetKnownNodes",
			Handler:    _NodeCommunication_GetKnownNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}