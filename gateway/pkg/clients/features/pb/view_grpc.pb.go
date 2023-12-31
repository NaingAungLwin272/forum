// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: view.proto

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
	ViewService_CreateView_FullMethodName                 = "/pb.ViewService/CreateView"
	ViewService_GetViews_FullMethodName                   = "/pb.ViewService/GetViews"
	ViewService_GetView_FullMethodName                    = "/pb.ViewService/GetView"
	ViewService_UpdateView_FullMethodName                 = "/pb.ViewService/UpdateView"
	ViewService_DeleteView_FullMethodName                 = "/pb.ViewService/DeleteView"
	ViewService_GetViewsByUserId_FullMethodName           = "/pb.ViewService/GetViewsByUserId"
	ViewService_GetViewsByUserIdQuestionId_FullMethodName = "/pb.ViewService/GetViewsByUserIdQuestionId"
)

// ViewServiceClient is the client API for ViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewServiceClient interface {
	CreateView(ctx context.Context, in *CreateViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
	GetViews(ctx context.Context, in *GetViewsRequest, opts ...grpc.CallOption) (ViewService_GetViewsClient, error)
	GetView(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
	UpdateView(ctx context.Context, in *UpdateViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
	DeleteView(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*DeleteViewResponse, error)
	GetViewsByUserId(ctx context.Context, in *ViewRequestByUserId, opts ...grpc.CallOption) (ViewService_GetViewsByUserIdClient, error)
	GetViewsByUserIdQuestionId(ctx context.Context, in *ViewRequestByUserIdQuestionId, opts ...grpc.CallOption) (ViewService_GetViewsByUserIdQuestionIdClient, error)
}

type viewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViewServiceClient(cc grpc.ClientConnInterface) ViewServiceClient {
	return &viewServiceClient{cc}
}

func (c *viewServiceClient) CreateView(ctx context.Context, in *CreateViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, ViewService_CreateView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetViews(ctx context.Context, in *GetViewsRequest, opts ...grpc.CallOption) (ViewService_GetViewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ViewService_ServiceDesc.Streams[0], ViewService_GetViews_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &viewServiceGetViewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewService_GetViewsClient interface {
	Recv() (*View, error)
	grpc.ClientStream
}

type viewServiceGetViewsClient struct {
	grpc.ClientStream
}

func (x *viewServiceGetViewsClient) Recv() (*View, error) {
	m := new(View)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewServiceClient) GetView(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, ViewService_GetView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) UpdateView(ctx context.Context, in *UpdateViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, ViewService_UpdateView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) DeleteView(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*DeleteViewResponse, error) {
	out := new(DeleteViewResponse)
	err := c.cc.Invoke(ctx, ViewService_DeleteView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetViewsByUserId(ctx context.Context, in *ViewRequestByUserId, opts ...grpc.CallOption) (ViewService_GetViewsByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &ViewService_ServiceDesc.Streams[1], ViewService_GetViewsByUserId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &viewServiceGetViewsByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewService_GetViewsByUserIdClient interface {
	Recv() (*View, error)
	grpc.ClientStream
}

type viewServiceGetViewsByUserIdClient struct {
	grpc.ClientStream
}

func (x *viewServiceGetViewsByUserIdClient) Recv() (*View, error) {
	m := new(View)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewServiceClient) GetViewsByUserIdQuestionId(ctx context.Context, in *ViewRequestByUserIdQuestionId, opts ...grpc.CallOption) (ViewService_GetViewsByUserIdQuestionIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &ViewService_ServiceDesc.Streams[2], ViewService_GetViewsByUserIdQuestionId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &viewServiceGetViewsByUserIdQuestionIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewService_GetViewsByUserIdQuestionIdClient interface {
	Recv() (*View, error)
	grpc.ClientStream
}

type viewServiceGetViewsByUserIdQuestionIdClient struct {
	grpc.ClientStream
}

func (x *viewServiceGetViewsByUserIdQuestionIdClient) Recv() (*View, error) {
	m := new(View)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ViewServiceServer is the server API for ViewService service.
// All implementations must embed UnimplementedViewServiceServer
// for forward compatibility
type ViewServiceServer interface {
	CreateView(context.Context, *CreateViewRequest) (*ViewResponse, error)
	GetViews(*GetViewsRequest, ViewService_GetViewsServer) error
	GetView(context.Context, *ViewRequest) (*ViewResponse, error)
	UpdateView(context.Context, *UpdateViewRequest) (*ViewResponse, error)
	DeleteView(context.Context, *ViewRequest) (*DeleteViewResponse, error)
	GetViewsByUserId(*ViewRequestByUserId, ViewService_GetViewsByUserIdServer) error
	GetViewsByUserIdQuestionId(*ViewRequestByUserIdQuestionId, ViewService_GetViewsByUserIdQuestionIdServer) error
	mustEmbedUnimplementedViewServiceServer()
}

// UnimplementedViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedViewServiceServer struct {
}

func (UnimplementedViewServiceServer) CreateView(context.Context, *CreateViewRequest) (*ViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateView not implemented")
}
func (UnimplementedViewServiceServer) GetViews(*GetViewsRequest, ViewService_GetViewsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetViews not implemented")
}
func (UnimplementedViewServiceServer) GetView(context.Context, *ViewRequest) (*ViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetView not implemented")
}
func (UnimplementedViewServiceServer) UpdateView(context.Context, *UpdateViewRequest) (*ViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateView not implemented")
}
func (UnimplementedViewServiceServer) DeleteView(context.Context, *ViewRequest) (*DeleteViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteView not implemented")
}
func (UnimplementedViewServiceServer) GetViewsByUserId(*ViewRequestByUserId, ViewService_GetViewsByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetViewsByUserId not implemented")
}
func (UnimplementedViewServiceServer) GetViewsByUserIdQuestionId(*ViewRequestByUserIdQuestionId, ViewService_GetViewsByUserIdQuestionIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetViewsByUserIdQuestionId not implemented")
}
func (UnimplementedViewServiceServer) mustEmbedUnimplementedViewServiceServer() {}

// UnsafeViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewServiceServer will
// result in compilation errors.
type UnsafeViewServiceServer interface {
	mustEmbedUnimplementedViewServiceServer()
}

func RegisterViewServiceServer(s grpc.ServiceRegistrar, srv ViewServiceServer) {
	s.RegisterService(&ViewService_ServiceDesc, srv)
}

func _ViewService_CreateView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).CreateView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewService_CreateView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).CreateView(ctx, req.(*CreateViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetViews_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetViewsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewServiceServer).GetViews(m, &viewServiceGetViewsServer{stream})
}

type ViewService_GetViewsServer interface {
	Send(*View) error
	grpc.ServerStream
}

type viewServiceGetViewsServer struct {
	grpc.ServerStream
}

func (x *viewServiceGetViewsServer) Send(m *View) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewService_GetView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewService_GetView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetView(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_UpdateView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).UpdateView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewService_UpdateView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).UpdateView(ctx, req.(*UpdateViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_DeleteView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).DeleteView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewService_DeleteView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).DeleteView(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetViewsByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ViewRequestByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewServiceServer).GetViewsByUserId(m, &viewServiceGetViewsByUserIdServer{stream})
}

type ViewService_GetViewsByUserIdServer interface {
	Send(*View) error
	grpc.ServerStream
}

type viewServiceGetViewsByUserIdServer struct {
	grpc.ServerStream
}

func (x *viewServiceGetViewsByUserIdServer) Send(m *View) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewService_GetViewsByUserIdQuestionId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ViewRequestByUserIdQuestionId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewServiceServer).GetViewsByUserIdQuestionId(m, &viewServiceGetViewsByUserIdQuestionIdServer{stream})
}

type ViewService_GetViewsByUserIdQuestionIdServer interface {
	Send(*View) error
	grpc.ServerStream
}

type viewServiceGetViewsByUserIdQuestionIdServer struct {
	grpc.ServerStream
}

func (x *viewServiceGetViewsByUserIdQuestionIdServer) Send(m *View) error {
	return x.ServerStream.SendMsg(m)
}

// ViewService_ServiceDesc is the grpc.ServiceDesc for ViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ViewService",
	HandlerType: (*ViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateView",
			Handler:    _ViewService_CreateView_Handler,
		},
		{
			MethodName: "GetView",
			Handler:    _ViewService_GetView_Handler,
		},
		{
			MethodName: "UpdateView",
			Handler:    _ViewService_UpdateView_Handler,
		},
		{
			MethodName: "DeleteView",
			Handler:    _ViewService_DeleteView_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetViews",
			Handler:       _ViewService_GetViews_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetViewsByUserId",
			Handler:       _ViewService_GetViewsByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetViewsByUserIdQuestionId",
			Handler:       _ViewService_GetViewsByUserIdQuestionId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "view.proto",
}
