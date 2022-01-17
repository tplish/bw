// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: protos/api.proto

package protos

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

// BwApiClient is the client API for BwApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BwApiClient interface {
	ListItems(ctx context.Context, in *ListItemsReq, opts ...grpc.CallOption) (*ListItemsResp, error)
}

type bwApiClient struct {
	cc grpc.ClientConnInterface
}

func NewBwApiClient(cc grpc.ClientConnInterface) BwApiClient {
	return &bwApiClient{cc}
}

func (c *bwApiClient) ListItems(ctx context.Context, in *ListItemsReq, opts ...grpc.CallOption) (*ListItemsResp, error) {
	out := new(ListItemsResp)
	err := c.cc.Invoke(ctx, "/bwapi.BwApi/ListItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BwApiServer is the server API for BwApi service.
// All implementations must embed UnimplementedBwApiServer
// for forward compatibility
type BwApiServer interface {
	ListItems(context.Context, *ListItemsReq) (*ListItemsResp, error)
	mustEmbedUnimplementedBwApiServer()
}

// UnimplementedBwApiServer must be embedded to have forward compatible implementations.
type UnimplementedBwApiServer struct {
}

func (UnimplementedBwApiServer) ListItems(context.Context, *ListItemsReq) (*ListItemsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedBwApiServer) mustEmbedUnimplementedBwApiServer() {}

// UnsafeBwApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BwApiServer will
// result in compilation errors.
type UnsafeBwApiServer interface {
	mustEmbedUnimplementedBwApiServer()
}

func RegisterBwApiServer(s grpc.ServiceRegistrar, srv BwApiServer) {
	s.RegisterService(&BwApi_ServiceDesc, srv)
}

func _BwApi_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BwApiServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bwapi.BwApi/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BwApiServer).ListItems(ctx, req.(*ListItemsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BwApi_ServiceDesc is the grpc.ServiceDesc for BwApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BwApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bwapi.BwApi",
	HandlerType: (*BwApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListItems",
			Handler:    _BwApi_ListItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/api.proto",
}
