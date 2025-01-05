// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: wishlist-service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WishlistService_CreateWishlist_FullMethodName = "/proto.WishlistService/CreateWishlist"
)

// WishlistServiceClient is the client API for WishlistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WishlistServiceClient interface {
	CreateWishlist(ctx context.Context, in *CreateWishlistRequest, opts ...grpc.CallOption) (*CreateWishlistResponse, error)
}

type wishlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWishlistServiceClient(cc grpc.ClientConnInterface) WishlistServiceClient {
	return &wishlistServiceClient{cc}
}

func (c *wishlistServiceClient) CreateWishlist(ctx context.Context, in *CreateWishlistRequest, opts ...grpc.CallOption) (*CreateWishlistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWishlistResponse)
	err := c.cc.Invoke(ctx, WishlistService_CreateWishlist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WishlistServiceServer is the server API for WishlistService service.
// All implementations must embed UnimplementedWishlistServiceServer
// for forward compatibility.
type WishlistServiceServer interface {
	CreateWishlist(context.Context, *CreateWishlistRequest) (*CreateWishlistResponse, error)
	mustEmbedUnimplementedWishlistServiceServer()
}

// UnimplementedWishlistServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWishlistServiceServer struct{}

func (UnimplementedWishlistServiceServer) CreateWishlist(context.Context, *CreateWishlistRequest) (*CreateWishlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWishlist not implemented")
}
func (UnimplementedWishlistServiceServer) mustEmbedUnimplementedWishlistServiceServer() {}
func (UnimplementedWishlistServiceServer) testEmbeddedByValue()                         {}

// UnsafeWishlistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WishlistServiceServer will
// result in compilation errors.
type UnsafeWishlistServiceServer interface {
	mustEmbedUnimplementedWishlistServiceServer()
}

func RegisterWishlistServiceServer(s grpc.ServiceRegistrar, srv WishlistServiceServer) {
	// If the following call pancis, it indicates UnimplementedWishlistServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WishlistService_ServiceDesc, srv)
}

func _WishlistService_CreateWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWishlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).CreateWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WishlistService_CreateWishlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).CreateWishlist(ctx, req.(*CreateWishlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WishlistService_ServiceDesc is the grpc.ServiceDesc for WishlistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WishlistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WishlistService",
	HandlerType: (*WishlistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWishlist",
			Handler:    _WishlistService_CreateWishlist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wishlist-service.proto",
}