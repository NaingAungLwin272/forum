// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: badge.proto

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
	BadgeService_CreateBadge_FullMethodName = "/pb.BadgeService/CreateBadge"
	BadgeService_GetBadge_FullMethodName    = "/pb.BadgeService/GetBadge"
	BadgeService_GetBadges_FullMethodName   = "/pb.BadgeService/GetBadges"
	BadgeService_UpdateBadge_FullMethodName = "/pb.BadgeService/UpdateBadge"
	BadgeService_DeleteBadge_FullMethodName = "/pb.BadgeService/DeleteBadge"
)

// BadgeServiceClient is the client API for BadgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BadgeServiceClient interface {
	CreateBadge(ctx context.Context, in *CreateBadgeRequest, opts ...grpc.CallOption) (*BadgeResponse, error)
	GetBadge(ctx context.Context, in *BadgeRequest, opts ...grpc.CallOption) (*BadgeResponse, error)
	GetBadges(ctx context.Context, in *GetBadgesRequest, opts ...grpc.CallOption) (*BadgeResponseList, error)
	UpdateBadge(ctx context.Context, in *UpdateBadgeRequest, opts ...grpc.CallOption) (*BadgeResponse, error)
	DeleteBadge(ctx context.Context, in *BadgeRequest, opts ...grpc.CallOption) (*DeleteBadgeResponse, error)
}

type badgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBadgeServiceClient(cc grpc.ClientConnInterface) BadgeServiceClient {
	return &badgeServiceClient{cc}
}

func (c *badgeServiceClient) CreateBadge(ctx context.Context, in *CreateBadgeRequest, opts ...grpc.CallOption) (*BadgeResponse, error) {
	out := new(BadgeResponse)
	err := c.cc.Invoke(ctx, BadgeService_CreateBadge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) GetBadge(ctx context.Context, in *BadgeRequest, opts ...grpc.CallOption) (*BadgeResponse, error) {
	out := new(BadgeResponse)
	err := c.cc.Invoke(ctx, BadgeService_GetBadge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) GetBadges(ctx context.Context, in *GetBadgesRequest, opts ...grpc.CallOption) (*BadgeResponseList, error) {
	out := new(BadgeResponseList)
	err := c.cc.Invoke(ctx, BadgeService_GetBadges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) UpdateBadge(ctx context.Context, in *UpdateBadgeRequest, opts ...grpc.CallOption) (*BadgeResponse, error) {
	out := new(BadgeResponse)
	err := c.cc.Invoke(ctx, BadgeService_UpdateBadge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) DeleteBadge(ctx context.Context, in *BadgeRequest, opts ...grpc.CallOption) (*DeleteBadgeResponse, error) {
	out := new(DeleteBadgeResponse)
	err := c.cc.Invoke(ctx, BadgeService_DeleteBadge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadgeServiceServer is the server API for BadgeService service.
// All implementations must embed UnimplementedBadgeServiceServer
// for forward compatibility
type BadgeServiceServer interface {
	CreateBadge(context.Context, *CreateBadgeRequest) (*BadgeResponse, error)
	GetBadge(context.Context, *BadgeRequest) (*BadgeResponse, error)
	GetBadges(context.Context, *GetBadgesRequest) (*BadgeResponseList, error)
	UpdateBadge(context.Context, *UpdateBadgeRequest) (*BadgeResponse, error)
	DeleteBadge(context.Context, *BadgeRequest) (*DeleteBadgeResponse, error)
	mustEmbedUnimplementedBadgeServiceServer()
}

// UnimplementedBadgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBadgeServiceServer struct {
}

func (UnimplementedBadgeServiceServer) CreateBadge(context.Context, *CreateBadgeRequest) (*BadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBadge not implemented")
}
func (UnimplementedBadgeServiceServer) GetBadge(context.Context, *BadgeRequest) (*BadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBadge not implemented")
}
func (UnimplementedBadgeServiceServer) GetBadges(context.Context, *GetBadgesRequest) (*BadgeResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBadges not implemented")
}
func (UnimplementedBadgeServiceServer) UpdateBadge(context.Context, *UpdateBadgeRequest) (*BadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBadge not implemented")
}
func (UnimplementedBadgeServiceServer) DeleteBadge(context.Context, *BadgeRequest) (*DeleteBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBadge not implemented")
}
func (UnimplementedBadgeServiceServer) mustEmbedUnimplementedBadgeServiceServer() {}

// UnsafeBadgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BadgeServiceServer will
// result in compilation errors.
type UnsafeBadgeServiceServer interface {
	mustEmbedUnimplementedBadgeServiceServer()
}

func RegisterBadgeServiceServer(s grpc.ServiceRegistrar, srv BadgeServiceServer) {
	s.RegisterService(&BadgeService_ServiceDesc, srv)
}

func _BadgeService_CreateBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).CreateBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BadgeService_CreateBadge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).CreateBadge(ctx, req.(*CreateBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_GetBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).GetBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BadgeService_GetBadge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).GetBadge(ctx, req.(*BadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_GetBadges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBadgesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).GetBadges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BadgeService_GetBadges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).GetBadges(ctx, req.(*GetBadgesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_UpdateBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).UpdateBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BadgeService_UpdateBadge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).UpdateBadge(ctx, req.(*UpdateBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_DeleteBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).DeleteBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BadgeService_DeleteBadge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).DeleteBadge(ctx, req.(*BadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BadgeService_ServiceDesc is the grpc.ServiceDesc for BadgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BadgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BadgeService",
	HandlerType: (*BadgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBadge",
			Handler:    _BadgeService_CreateBadge_Handler,
		},
		{
			MethodName: "GetBadge",
			Handler:    _BadgeService_GetBadge_Handler,
		},
		{
			MethodName: "GetBadges",
			Handler:    _BadgeService_GetBadges_Handler,
		},
		{
			MethodName: "UpdateBadge",
			Handler:    _BadgeService_UpdateBadge_Handler,
		},
		{
			MethodName: "DeleteBadge",
			Handler:    _BadgeService_DeleteBadge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "badge.proto",
}
