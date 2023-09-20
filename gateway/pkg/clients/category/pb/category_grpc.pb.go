// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: category.proto

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

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*CategoryResponseList, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	DeleteCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
	GetCategoryByType(ctx context.Context, in *GetCategoryByTypeRequest, opts ...grpc.CallOption) (*CategoryResponseList, error)
	GetCategoryCount(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*CategoryCountResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*CategoryResponseList, error) {
	out := new(CategoryResponseList)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryByType(ctx context.Context, in *GetCategoryByTypeRequest, opts ...grpc.CallOption) (*CategoryResponseList, error) {
	out := new(CategoryResponseList)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategoryByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryCount(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*CategoryCountResponse, error) {
	out := new(CategoryCountResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategoryCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error)
	GetCategory(context.Context, *CategoryRequest) (*CategoryResponse, error)
	GetCategories(context.Context, *GetCategoriesRequest) (*CategoryResponseList, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*CategoryResponse, error)
	DeleteCategory(context.Context, *CategoryRequest) (*DeleteCategoryResponse, error)
	GetCategoryByType(context.Context, *GetCategoryByTypeRequest) (*CategoryResponseList, error)
	GetCategoryCount(context.Context, *GetCategoriesRequest) (*CategoryCountResponse, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategory(context.Context, *CategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategories(context.Context, *GetCategoriesRequest) (*CategoryResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedCategoryServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteCategory(context.Context, *CategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategoryByType(context.Context, *GetCategoryByTypeRequest) (*CategoryResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryByType not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategoryCount(context.Context, *GetCategoriesRequest) (*CategoryCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryCount not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategoryByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryByType(ctx, req.(*GetCategoryByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategoryCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryCount(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryService_CreateCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CategoryService_GetCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _CategoryService_GetCategories_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CategoryService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CategoryService_DeleteCategory_Handler,
		},
		{
			MethodName: "GetCategoryByType",
			Handler:    _CategoryService_GetCategoryByType_Handler,
		},
		{
			MethodName: "GetCategoryCount",
			Handler:    _CategoryService_GetCategoryCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}