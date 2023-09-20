// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: userPoint.proto

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
	UserPointService_CreateUserPoint_FullMethodName = "/pb.UserPointService/CreateUserPoint"
	UserPointService_GetUserPoint_FullMethodName    = "/pb.UserPointService/GetUserPoint"
	UserPointService_GetUserPoints_FullMethodName   = "/pb.UserPointService/GetUserPoints"
	UserPointService_UpdateUserPoint_FullMethodName = "/pb.UserPointService/UpdateUserPoint"
	UserPointService_DeleteUserPoint_FullMethodName = "/pb.UserPointService/DeleteUserPoint"
	UserPointService_EvaluatePoints_FullMethodName  = "/pb.UserPointService/EvaluatePoints"
)

// UserPointServiceClient is the client API for UserPointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserPointServiceClient interface {
	CreateUserPoint(ctx context.Context, in *CreateUserPointRequest, opts ...grpc.CallOption) (*UserPointResponse, error)
	GetUserPoint(ctx context.Context, in *UserPointRequest, opts ...grpc.CallOption) (*UserPointResponse, error)
	GetUserPoints(ctx context.Context, in *GetUserPointsRequest, opts ...grpc.CallOption) (*UserPointResponseList, error)
	UpdateUserPoint(ctx context.Context, in *UpdateUserPointRequest, opts ...grpc.CallOption) (*UserPointResponse, error)
	DeleteUserPoint(ctx context.Context, in *UserPointRequest, opts ...grpc.CallOption) (*DeleteUserPointResponse, error)
	EvaluatePoints(ctx context.Context, in *GetUserPointsRequest, opts ...grpc.CallOption) (*UserPointEvaluateResponse, error)
}

type userPointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserPointServiceClient(cc grpc.ClientConnInterface) UserPointServiceClient {
	return &userPointServiceClient{cc}
}

func (c *userPointServiceClient) CreateUserPoint(ctx context.Context, in *CreateUserPointRequest, opts ...grpc.CallOption) (*UserPointResponse, error) {
	out := new(UserPointResponse)
	err := c.cc.Invoke(ctx, UserPointService_CreateUserPoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPointServiceClient) GetUserPoint(ctx context.Context, in *UserPointRequest, opts ...grpc.CallOption) (*UserPointResponse, error) {
	out := new(UserPointResponse)
	err := c.cc.Invoke(ctx, UserPointService_GetUserPoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPointServiceClient) GetUserPoints(ctx context.Context, in *GetUserPointsRequest, opts ...grpc.CallOption) (*UserPointResponseList, error) {
	out := new(UserPointResponseList)
	err := c.cc.Invoke(ctx, UserPointService_GetUserPoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPointServiceClient) UpdateUserPoint(ctx context.Context, in *UpdateUserPointRequest, opts ...grpc.CallOption) (*UserPointResponse, error) {
	out := new(UserPointResponse)
	err := c.cc.Invoke(ctx, UserPointService_UpdateUserPoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPointServiceClient) DeleteUserPoint(ctx context.Context, in *UserPointRequest, opts ...grpc.CallOption) (*DeleteUserPointResponse, error) {
	out := new(DeleteUserPointResponse)
	err := c.cc.Invoke(ctx, UserPointService_DeleteUserPoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPointServiceClient) EvaluatePoints(ctx context.Context, in *GetUserPointsRequest, opts ...grpc.CallOption) (*UserPointEvaluateResponse, error) {
	out := new(UserPointEvaluateResponse)
	err := c.cc.Invoke(ctx, UserPointService_EvaluatePoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserPointServiceServer is the server API for UserPointService service.
// All implementations must embed UnimplementedUserPointServiceServer
// for forward compatibility
type UserPointServiceServer interface {
	CreateUserPoint(context.Context, *CreateUserPointRequest) (*UserPointResponse, error)
	GetUserPoint(context.Context, *UserPointRequest) (*UserPointResponse, error)
	GetUserPoints(context.Context, *GetUserPointsRequest) (*UserPointResponseList, error)
	UpdateUserPoint(context.Context, *UpdateUserPointRequest) (*UserPointResponse, error)
	DeleteUserPoint(context.Context, *UserPointRequest) (*DeleteUserPointResponse, error)
	EvaluatePoints(context.Context, *GetUserPointsRequest) (*UserPointEvaluateResponse, error)
	mustEmbedUnimplementedUserPointServiceServer()
}

// UnimplementedUserPointServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserPointServiceServer struct {
}

func (UnimplementedUserPointServiceServer) CreateUserPoint(context.Context, *CreateUserPointRequest) (*UserPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserPoint not implemented")
}
func (UnimplementedUserPointServiceServer) GetUserPoint(context.Context, *UserPointRequest) (*UserPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPoint not implemented")
}
func (UnimplementedUserPointServiceServer) GetUserPoints(context.Context, *GetUserPointsRequest) (*UserPointResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPoints not implemented")
}
func (UnimplementedUserPointServiceServer) UpdateUserPoint(context.Context, *UpdateUserPointRequest) (*UserPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPoint not implemented")
}
func (UnimplementedUserPointServiceServer) DeleteUserPoint(context.Context, *UserPointRequest) (*DeleteUserPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserPoint not implemented")
}
func (UnimplementedUserPointServiceServer) EvaluatePoints(context.Context, *GetUserPointsRequest) (*UserPointEvaluateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluatePoints not implemented")
}
func (UnimplementedUserPointServiceServer) mustEmbedUnimplementedUserPointServiceServer() {}

// UnsafeUserPointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserPointServiceServer will
// result in compilation errors.
type UnsafeUserPointServiceServer interface {
	mustEmbedUnimplementedUserPointServiceServer()
}

func RegisterUserPointServiceServer(s grpc.ServiceRegistrar, srv UserPointServiceServer) {
	s.RegisterService(&UserPointService_ServiceDesc, srv)
}

func _UserPointService_CreateUserPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPointServiceServer).CreateUserPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserPointService_CreateUserPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPointServiceServer).CreateUserPoint(ctx, req.(*CreateUserPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPointService_GetUserPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPointServiceServer).GetUserPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserPointService_GetUserPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPointServiceServer).GetUserPoint(ctx, req.(*UserPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPointService_GetUserPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPointServiceServer).GetUserPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserPointService_GetUserPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPointServiceServer).GetUserPoints(ctx, req.(*GetUserPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPointService_UpdateUserPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPointServiceServer).UpdateUserPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserPointService_UpdateUserPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPointServiceServer).UpdateUserPoint(ctx, req.(*UpdateUserPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPointService_DeleteUserPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPointServiceServer).DeleteUserPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserPointService_DeleteUserPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPointServiceServer).DeleteUserPoint(ctx, req.(*UserPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPointService_EvaluatePoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPointServiceServer).EvaluatePoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserPointService_EvaluatePoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPointServiceServer).EvaluatePoints(ctx, req.(*GetUserPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserPointService_ServiceDesc is the grpc.ServiceDesc for UserPointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserPointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserPointService",
	HandlerType: (*UserPointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserPoint",
			Handler:    _UserPointService_CreateUserPoint_Handler,
		},
		{
			MethodName: "GetUserPoint",
			Handler:    _UserPointService_GetUserPoint_Handler,
		},
		{
			MethodName: "GetUserPoints",
			Handler:    _UserPointService_GetUserPoints_Handler,
		},
		{
			MethodName: "UpdateUserPoint",
			Handler:    _UserPointService_UpdateUserPoint_Handler,
		},
		{
			MethodName: "DeleteUserPoint",
			Handler:    _UserPointService_DeleteUserPoint_Handler,
		},
		{
			MethodName: "EvaluatePoints",
			Handler:    _UserPointService_EvaluatePoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userPoint.proto",
}