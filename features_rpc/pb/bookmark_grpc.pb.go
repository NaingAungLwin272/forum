// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: bookmark.proto

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
	BookmarkService_CreateBookmark_FullMethodName                 = "/pb.BookmarkService/CreateBookmark"
	BookmarkService_GetBookmarks_FullMethodName                   = "/pb.BookmarkService/GetBookmarks"
	BookmarkService_GetBookmark_FullMethodName                    = "/pb.BookmarkService/GetBookmark"
	BookmarkService_UpdateBookmark_FullMethodName                 = "/pb.BookmarkService/UpdateBookmark"
	BookmarkService_DeleteBookmark_FullMethodName                 = "/pb.BookmarkService/DeleteBookmark"
	BookmarkService_GetBookmarksByUserId_FullMethodName           = "/pb.BookmarkService/GetBookmarksByUserId"
	BookmarkService_GetBookmarkCount_FullMethodName               = "/pb.BookmarkService/GetBookmarkCount"
	BookmarkService_GetBookmarksByUserIdQuestionId_FullMethodName = "/pb.BookmarkService/GetBookmarksByUserIdQuestionId"
)

// BookmarkServiceClient is the client API for BookmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookmarkServiceClient interface {
	CreateBookmark(ctx context.Context, in *CreateBookmarkRequest, opts ...grpc.CallOption) (*BookmarkResponse, error)
	GetBookmarks(ctx context.Context, in *GetBookmarksRequest, opts ...grpc.CallOption) (BookmarkService_GetBookmarksClient, error)
	GetBookmark(ctx context.Context, in *BookmarkRequest, opts ...grpc.CallOption) (*BookmarkResponse, error)
	UpdateBookmark(ctx context.Context, in *UpdateBookmarkRequest, opts ...grpc.CallOption) (*BookmarkResponse, error)
	DeleteBookmark(ctx context.Context, in *BookmarkRequest, opts ...grpc.CallOption) (*DeleteBookmarkResponse, error)
	GetBookmarksByUserId(ctx context.Context, in *BookmarkRequestByUserId, opts ...grpc.CallOption) (BookmarkService_GetBookmarksByUserIdClient, error)
	GetBookmarkCount(ctx context.Context, in *BookmarkRequestByUserId, opts ...grpc.CallOption) (*BookmarkCountResponse, error)
	GetBookmarksByUserIdQuestionId(ctx context.Context, in *BookmarkRequestByUserIdQuestionId, opts ...grpc.CallOption) (BookmarkService_GetBookmarksByUserIdQuestionIdClient, error)
}

type bookmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarkServiceClient(cc grpc.ClientConnInterface) BookmarkServiceClient {
	return &bookmarkServiceClient{cc}
}

func (c *bookmarkServiceClient) CreateBookmark(ctx context.Context, in *CreateBookmarkRequest, opts ...grpc.CallOption) (*BookmarkResponse, error) {
	out := new(BookmarkResponse)
	err := c.cc.Invoke(ctx, BookmarkService_CreateBookmark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) GetBookmarks(ctx context.Context, in *GetBookmarksRequest, opts ...grpc.CallOption) (BookmarkService_GetBookmarksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookmarkService_ServiceDesc.Streams[0], BookmarkService_GetBookmarks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bookmarkServiceGetBookmarksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookmarkService_GetBookmarksClient interface {
	Recv() (*Bookmark, error)
	grpc.ClientStream
}

type bookmarkServiceGetBookmarksClient struct {
	grpc.ClientStream
}

func (x *bookmarkServiceGetBookmarksClient) Recv() (*Bookmark, error) {
	m := new(Bookmark)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookmarkServiceClient) GetBookmark(ctx context.Context, in *BookmarkRequest, opts ...grpc.CallOption) (*BookmarkResponse, error) {
	out := new(BookmarkResponse)
	err := c.cc.Invoke(ctx, BookmarkService_GetBookmark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) UpdateBookmark(ctx context.Context, in *UpdateBookmarkRequest, opts ...grpc.CallOption) (*BookmarkResponse, error) {
	out := new(BookmarkResponse)
	err := c.cc.Invoke(ctx, BookmarkService_UpdateBookmark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) DeleteBookmark(ctx context.Context, in *BookmarkRequest, opts ...grpc.CallOption) (*DeleteBookmarkResponse, error) {
	out := new(DeleteBookmarkResponse)
	err := c.cc.Invoke(ctx, BookmarkService_DeleteBookmark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) GetBookmarksByUserId(ctx context.Context, in *BookmarkRequestByUserId, opts ...grpc.CallOption) (BookmarkService_GetBookmarksByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookmarkService_ServiceDesc.Streams[1], BookmarkService_GetBookmarksByUserId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bookmarkServiceGetBookmarksByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookmarkService_GetBookmarksByUserIdClient interface {
	Recv() (*Bookmark, error)
	grpc.ClientStream
}

type bookmarkServiceGetBookmarksByUserIdClient struct {
	grpc.ClientStream
}

func (x *bookmarkServiceGetBookmarksByUserIdClient) Recv() (*Bookmark, error) {
	m := new(Bookmark)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookmarkServiceClient) GetBookmarkCount(ctx context.Context, in *BookmarkRequestByUserId, opts ...grpc.CallOption) (*BookmarkCountResponse, error) {
	out := new(BookmarkCountResponse)
	err := c.cc.Invoke(ctx, BookmarkService_GetBookmarkCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) GetBookmarksByUserIdQuestionId(ctx context.Context, in *BookmarkRequestByUserIdQuestionId, opts ...grpc.CallOption) (BookmarkService_GetBookmarksByUserIdQuestionIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookmarkService_ServiceDesc.Streams[2], BookmarkService_GetBookmarksByUserIdQuestionId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bookmarkServiceGetBookmarksByUserIdQuestionIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookmarkService_GetBookmarksByUserIdQuestionIdClient interface {
	Recv() (*Bookmark, error)
	grpc.ClientStream
}

type bookmarkServiceGetBookmarksByUserIdQuestionIdClient struct {
	grpc.ClientStream
}

func (x *bookmarkServiceGetBookmarksByUserIdQuestionIdClient) Recv() (*Bookmark, error) {
	m := new(Bookmark)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookmarkServiceServer is the server API for BookmarkService service.
// All implementations must embed UnimplementedBookmarkServiceServer
// for forward compatibility
type BookmarkServiceServer interface {
	CreateBookmark(context.Context, *CreateBookmarkRequest) (*BookmarkResponse, error)
	GetBookmarks(*GetBookmarksRequest, BookmarkService_GetBookmarksServer) error
	GetBookmark(context.Context, *BookmarkRequest) (*BookmarkResponse, error)
	UpdateBookmark(context.Context, *UpdateBookmarkRequest) (*BookmarkResponse, error)
	DeleteBookmark(context.Context, *BookmarkRequest) (*DeleteBookmarkResponse, error)
	GetBookmarksByUserId(*BookmarkRequestByUserId, BookmarkService_GetBookmarksByUserIdServer) error
	GetBookmarkCount(context.Context, *BookmarkRequestByUserId) (*BookmarkCountResponse, error)
	GetBookmarksByUserIdQuestionId(*BookmarkRequestByUserIdQuestionId, BookmarkService_GetBookmarksByUserIdQuestionIdServer) error
	mustEmbedUnimplementedBookmarkServiceServer()
}

// UnimplementedBookmarkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookmarkServiceServer struct {
}

func (UnimplementedBookmarkServiceServer) CreateBookmark(context.Context, *CreateBookmarkRequest) (*BookmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookmark not implemented")
}
func (UnimplementedBookmarkServiceServer) GetBookmarks(*GetBookmarksRequest, BookmarkService_GetBookmarksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBookmarks not implemented")
}
func (UnimplementedBookmarkServiceServer) GetBookmark(context.Context, *BookmarkRequest) (*BookmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookmark not implemented")
}
func (UnimplementedBookmarkServiceServer) UpdateBookmark(context.Context, *UpdateBookmarkRequest) (*BookmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookmark not implemented")
}
func (UnimplementedBookmarkServiceServer) DeleteBookmark(context.Context, *BookmarkRequest) (*DeleteBookmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookmark not implemented")
}
func (UnimplementedBookmarkServiceServer) GetBookmarksByUserId(*BookmarkRequestByUserId, BookmarkService_GetBookmarksByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBookmarksByUserId not implemented")
}
func (UnimplementedBookmarkServiceServer) GetBookmarkCount(context.Context, *BookmarkRequestByUserId) (*BookmarkCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookmarkCount not implemented")
}
func (UnimplementedBookmarkServiceServer) GetBookmarksByUserIdQuestionId(*BookmarkRequestByUserIdQuestionId, BookmarkService_GetBookmarksByUserIdQuestionIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBookmarksByUserIdQuestionId not implemented")
}
func (UnimplementedBookmarkServiceServer) mustEmbedUnimplementedBookmarkServiceServer() {}

// UnsafeBookmarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookmarkServiceServer will
// result in compilation errors.
type UnsafeBookmarkServiceServer interface {
	mustEmbedUnimplementedBookmarkServiceServer()
}

func RegisterBookmarkServiceServer(s grpc.ServiceRegistrar, srv BookmarkServiceServer) {
	s.RegisterService(&BookmarkService_ServiceDesc, srv)
}

func _BookmarkService_CreateBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).CreateBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_CreateBookmark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).CreateBookmark(ctx, req.(*CreateBookmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_GetBookmarks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBookmarksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookmarkServiceServer).GetBookmarks(m, &bookmarkServiceGetBookmarksServer{stream})
}

type BookmarkService_GetBookmarksServer interface {
	Send(*Bookmark) error
	grpc.ServerStream
}

type bookmarkServiceGetBookmarksServer struct {
	grpc.ServerStream
}

func (x *bookmarkServiceGetBookmarksServer) Send(m *Bookmark) error {
	return x.ServerStream.SendMsg(m)
}

func _BookmarkService_GetBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).GetBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_GetBookmark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).GetBookmark(ctx, req.(*BookmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_UpdateBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).UpdateBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_UpdateBookmark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).UpdateBookmark(ctx, req.(*UpdateBookmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_DeleteBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).DeleteBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_DeleteBookmark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).DeleteBookmark(ctx, req.(*BookmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_GetBookmarksByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookmarkRequestByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookmarkServiceServer).GetBookmarksByUserId(m, &bookmarkServiceGetBookmarksByUserIdServer{stream})
}

type BookmarkService_GetBookmarksByUserIdServer interface {
	Send(*Bookmark) error
	grpc.ServerStream
}

type bookmarkServiceGetBookmarksByUserIdServer struct {
	grpc.ServerStream
}

func (x *bookmarkServiceGetBookmarksByUserIdServer) Send(m *Bookmark) error {
	return x.ServerStream.SendMsg(m)
}

func _BookmarkService_GetBookmarkCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookmarkRequestByUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).GetBookmarkCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_GetBookmarkCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).GetBookmarkCount(ctx, req.(*BookmarkRequestByUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_GetBookmarksByUserIdQuestionId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookmarkRequestByUserIdQuestionId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookmarkServiceServer).GetBookmarksByUserIdQuestionId(m, &bookmarkServiceGetBookmarksByUserIdQuestionIdServer{stream})
}

type BookmarkService_GetBookmarksByUserIdQuestionIdServer interface {
	Send(*Bookmark) error
	grpc.ServerStream
}

type bookmarkServiceGetBookmarksByUserIdQuestionIdServer struct {
	grpc.ServerStream
}

func (x *bookmarkServiceGetBookmarksByUserIdQuestionIdServer) Send(m *Bookmark) error {
	return x.ServerStream.SendMsg(m)
}

// BookmarkService_ServiceDesc is the grpc.ServiceDesc for BookmarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BookmarkService",
	HandlerType: (*BookmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBookmark",
			Handler:    _BookmarkService_CreateBookmark_Handler,
		},
		{
			MethodName: "GetBookmark",
			Handler:    _BookmarkService_GetBookmark_Handler,
		},
		{
			MethodName: "UpdateBookmark",
			Handler:    _BookmarkService_UpdateBookmark_Handler,
		},
		{
			MethodName: "DeleteBookmark",
			Handler:    _BookmarkService_DeleteBookmark_Handler,
		},
		{
			MethodName: "GetBookmarkCount",
			Handler:    _BookmarkService_GetBookmarkCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBookmarks",
			Handler:       _BookmarkService_GetBookmarks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBookmarksByUserId",
			Handler:       _BookmarkService_GetBookmarksByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBookmarksByUserIdQuestionId",
			Handler:       _BookmarkService_GetBookmarksByUserIdQuestionId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bookmark.proto",
}
