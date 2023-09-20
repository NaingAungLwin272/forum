// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: question.proto

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
	QuestionService_CreateQuestion_FullMethodName           = "/pb.QuestionService/CreateQuestion"
	QuestionService_GetQuestions_FullMethodName             = "/pb.QuestionService/GetQuestions"
	QuestionService_GetQuestion_FullMethodName              = "/pb.QuestionService/GetQuestion"
	QuestionService_UpdateQuestion_FullMethodName           = "/pb.QuestionService/UpdateQuestion"
	QuestionService_DeleteQuestion_FullMethodName           = "/pb.QuestionService/DeleteQuestion"
	QuestionService_GetQuestionByUserId_FullMethodName      = "/pb.QuestionService/GetQuestionByUserId"
	QuestionService_GetAnswersByUserId_FullMethodName       = "/pb.QuestionService/GetAnswersByUserId"
	QuestionService_GetQuestionCount_FullMethodName         = "/pb.QuestionService/GetQuestionCount"
	QuestionService_FilterQuestion_FullMethodName           = "/pb.QuestionService/FilterQuestion"
	QuestionService_GetQuestionCountAll_FullMethodName      = "/pb.QuestionService/GetQuestionCountAll"
	QuestionService_GetFilteredQuestionCount_FullMethodName = "/pb.QuestionService/GetFilteredQuestionCount"
)

// QuestionServiceClient is the client API for QuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionServiceClient interface {
	CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error)
	GetQuestions(ctx context.Context, in *GetQuestionsRequest, opts ...grpc.CallOption) (QuestionService_GetQuestionsClient, error)
	GetQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error)
	UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error)
	DeleteQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*DeleteQuestionResponse, error)
	GetQuestionByUserId(ctx context.Context, in *QuestionResquestByUserId, opts ...grpc.CallOption) (QuestionService_GetQuestionByUserIdClient, error)
	GetAnswersByUserId(ctx context.Context, in *QuestionResquestByUserId, opts ...grpc.CallOption) (QuestionService_GetAnswersByUserIdClient, error)
	GetQuestionCount(ctx context.Context, in *QuestionResquestByUserId, opts ...grpc.CallOption) (*QuestionCountResponse, error)
	FilterQuestion(ctx context.Context, in *FilterQuestionRequest, opts ...grpc.CallOption) (QuestionService_FilterQuestionClient, error)
	GetQuestionCountAll(ctx context.Context, in *GetQuestionCountRequest, opts ...grpc.CallOption) (*QuestionCountResponse, error)
	GetFilteredQuestionCount(ctx context.Context, in *FilterQuestionRequest, opts ...grpc.CallOption) (*QuestionCountResponse, error)
}

type questionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionServiceClient(cc grpc.ClientConnInterface) QuestionServiceClient {
	return &questionServiceClient{cc}
}

func (c *questionServiceClient) CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error) {
	out := new(QuestionResponse)
	err := c.cc.Invoke(ctx, QuestionService_CreateQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetQuestions(ctx context.Context, in *GetQuestionsRequest, opts ...grpc.CallOption) (QuestionService_GetQuestionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuestionService_ServiceDesc.Streams[0], QuestionService_GetQuestions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &questionServiceGetQuestionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuestionService_GetQuestionsClient interface {
	Recv() (*Question, error)
	grpc.ClientStream
}

type questionServiceGetQuestionsClient struct {
	grpc.ClientStream
}

func (x *questionServiceGetQuestionsClient) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *questionServiceClient) GetQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error) {
	out := new(QuestionResponse)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error) {
	out := new(QuestionResponse)
	err := c.cc.Invoke(ctx, QuestionService_UpdateQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) DeleteQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*DeleteQuestionResponse, error) {
	out := new(DeleteQuestionResponse)
	err := c.cc.Invoke(ctx, QuestionService_DeleteQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetQuestionByUserId(ctx context.Context, in *QuestionResquestByUserId, opts ...grpc.CallOption) (QuestionService_GetQuestionByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuestionService_ServiceDesc.Streams[1], QuestionService_GetQuestionByUserId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &questionServiceGetQuestionByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuestionService_GetQuestionByUserIdClient interface {
	Recv() (*Question, error)
	grpc.ClientStream
}

type questionServiceGetQuestionByUserIdClient struct {
	grpc.ClientStream
}

func (x *questionServiceGetQuestionByUserIdClient) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *questionServiceClient) GetAnswersByUserId(ctx context.Context, in *QuestionResquestByUserId, opts ...grpc.CallOption) (QuestionService_GetAnswersByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuestionService_ServiceDesc.Streams[2], QuestionService_GetAnswersByUserId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &questionServiceGetAnswersByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuestionService_GetAnswersByUserIdClient interface {
	Recv() (*Question, error)
	grpc.ClientStream
}

type questionServiceGetAnswersByUserIdClient struct {
	grpc.ClientStream
}

func (x *questionServiceGetAnswersByUserIdClient) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *questionServiceClient) GetQuestionCount(ctx context.Context, in *QuestionResquestByUserId, opts ...grpc.CallOption) (*QuestionCountResponse, error) {
	out := new(QuestionCountResponse)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) FilterQuestion(ctx context.Context, in *FilterQuestionRequest, opts ...grpc.CallOption) (QuestionService_FilterQuestionClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuestionService_ServiceDesc.Streams[3], QuestionService_FilterQuestion_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &questionServiceFilterQuestionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuestionService_FilterQuestionClient interface {
	Recv() (*Question, error)
	grpc.ClientStream
}

type questionServiceFilterQuestionClient struct {
	grpc.ClientStream
}

func (x *questionServiceFilterQuestionClient) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *questionServiceClient) GetQuestionCountAll(ctx context.Context, in *GetQuestionCountRequest, opts ...grpc.CallOption) (*QuestionCountResponse, error) {
	out := new(QuestionCountResponse)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionCountAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetFilteredQuestionCount(ctx context.Context, in *FilterQuestionRequest, opts ...grpc.CallOption) (*QuestionCountResponse, error) {
	out := new(QuestionCountResponse)
	err := c.cc.Invoke(ctx, QuestionService_GetFilteredQuestionCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServiceServer is the server API for QuestionService service.
// All implementations must embed UnimplementedQuestionServiceServer
// for forward compatibility
type QuestionServiceServer interface {
	CreateQuestion(context.Context, *CreateQuestionRequest) (*QuestionResponse, error)
	GetQuestions(*GetQuestionsRequest, QuestionService_GetQuestionsServer) error
	GetQuestion(context.Context, *QuestionRequest) (*QuestionResponse, error)
	UpdateQuestion(context.Context, *UpdateQuestionRequest) (*QuestionResponse, error)
	DeleteQuestion(context.Context, *QuestionRequest) (*DeleteQuestionResponse, error)
	GetQuestionByUserId(*QuestionResquestByUserId, QuestionService_GetQuestionByUserIdServer) error
	GetAnswersByUserId(*QuestionResquestByUserId, QuestionService_GetAnswersByUserIdServer) error
	GetQuestionCount(context.Context, *QuestionResquestByUserId) (*QuestionCountResponse, error)
	FilterQuestion(*FilterQuestionRequest, QuestionService_FilterQuestionServer) error
	GetQuestionCountAll(context.Context, *GetQuestionCountRequest) (*QuestionCountResponse, error)
	GetFilteredQuestionCount(context.Context, *FilterQuestionRequest) (*QuestionCountResponse, error)
	mustEmbedUnimplementedQuestionServiceServer()
}

// UnimplementedQuestionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionServiceServer struct {
}

func (UnimplementedQuestionServiceServer) CreateQuestion(context.Context, *CreateQuestionRequest) (*QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestions(*GetQuestionsRequest, QuestionService_GetQuestionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetQuestions not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestion(context.Context, *QuestionRequest) (*QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) UpdateQuestion(context.Context, *UpdateQuestionRequest) (*QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) DeleteQuestion(context.Context, *QuestionRequest) (*DeleteQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestionByUserId(*QuestionResquestByUserId, QuestionService_GetQuestionByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetQuestionByUserId not implemented")
}
func (UnimplementedQuestionServiceServer) GetAnswersByUserId(*QuestionResquestByUserId, QuestionService_GetAnswersByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAnswersByUserId not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestionCount(context.Context, *QuestionResquestByUserId) (*QuestionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionCount not implemented")
}
func (UnimplementedQuestionServiceServer) FilterQuestion(*FilterQuestionRequest, QuestionService_FilterQuestionServer) error {
	return status.Errorf(codes.Unimplemented, "method FilterQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestionCountAll(context.Context, *GetQuestionCountRequest) (*QuestionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionCountAll not implemented")
}
func (UnimplementedQuestionServiceServer) GetFilteredQuestionCount(context.Context, *FilterQuestionRequest) (*QuestionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilteredQuestionCount not implemented")
}
func (UnimplementedQuestionServiceServer) mustEmbedUnimplementedQuestionServiceServer() {}

// UnsafeQuestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionServiceServer will
// result in compilation errors.
type UnsafeQuestionServiceServer interface {
	mustEmbedUnimplementedQuestionServiceServer()
}

func RegisterQuestionServiceServer(s grpc.ServiceRegistrar, srv QuestionServiceServer) {
	s.RegisterService(&QuestionService_ServiceDesc, srv)
}

func _QuestionService_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_CreateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).CreateQuestion(ctx, req.(*CreateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetQuestions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetQuestionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuestionServiceServer).GetQuestions(m, &questionServiceGetQuestionsServer{stream})
}

type QuestionService_GetQuestionsServer interface {
	Send(*Question) error
	grpc.ServerStream
}

type questionServiceGetQuestionsServer struct {
	grpc.ServerStream
}

func (x *questionServiceGetQuestionsServer) Send(m *Question) error {
	return x.ServerStream.SendMsg(m)
}

func _QuestionService_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestion(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_UpdateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).UpdateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_UpdateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).UpdateQuestion(ctx, req.(*UpdateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_DeleteQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).DeleteQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_DeleteQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).DeleteQuestion(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetQuestionByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QuestionResquestByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuestionServiceServer).GetQuestionByUserId(m, &questionServiceGetQuestionByUserIdServer{stream})
}

type QuestionService_GetQuestionByUserIdServer interface {
	Send(*Question) error
	grpc.ServerStream
}

type questionServiceGetQuestionByUserIdServer struct {
	grpc.ServerStream
}

func (x *questionServiceGetQuestionByUserIdServer) Send(m *Question) error {
	return x.ServerStream.SendMsg(m)
}

func _QuestionService_GetAnswersByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QuestionResquestByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuestionServiceServer).GetAnswersByUserId(m, &questionServiceGetAnswersByUserIdServer{stream})
}

type QuestionService_GetAnswersByUserIdServer interface {
	Send(*Question) error
	grpc.ServerStream
}

type questionServiceGetAnswersByUserIdServer struct {
	grpc.ServerStream
}

func (x *questionServiceGetAnswersByUserIdServer) Send(m *Question) error {
	return x.ServerStream.SendMsg(m)
}

func _QuestionService_GetQuestionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionResquestByUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionCount(ctx, req.(*QuestionResquestByUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_FilterQuestion_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FilterQuestionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuestionServiceServer).FilterQuestion(m, &questionServiceFilterQuestionServer{stream})
}

type QuestionService_FilterQuestionServer interface {
	Send(*Question) error
	grpc.ServerStream
}

type questionServiceFilterQuestionServer struct {
	grpc.ServerStream
}

func (x *questionServiceFilterQuestionServer) Send(m *Question) error {
	return x.ServerStream.SendMsg(m)
}

func _QuestionService_GetQuestionCountAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionCountAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionCountAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionCountAll(ctx, req.(*GetQuestionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetFilteredQuestionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetFilteredQuestionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetFilteredQuestionCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetFilteredQuestionCount(ctx, req.(*FilterQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionService_ServiceDesc is the grpc.ServiceDesc for QuestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.QuestionService",
	HandlerType: (*QuestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestion",
			Handler:    _QuestionService_CreateQuestion_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _QuestionService_GetQuestion_Handler,
		},
		{
			MethodName: "UpdateQuestion",
			Handler:    _QuestionService_UpdateQuestion_Handler,
		},
		{
			MethodName: "DeleteQuestion",
			Handler:    _QuestionService_DeleteQuestion_Handler,
		},
		{
			MethodName: "GetQuestionCount",
			Handler:    _QuestionService_GetQuestionCount_Handler,
		},
		{
			MethodName: "GetQuestionCountAll",
			Handler:    _QuestionService_GetQuestionCountAll_Handler,
		},
		{
			MethodName: "GetFilteredQuestionCount",
			Handler:    _QuestionService_GetFilteredQuestionCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetQuestions",
			Handler:       _QuestionService_GetQuestions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetQuestionByUserId",
			Handler:       _QuestionService_GetQuestionByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAnswersByUserId",
			Handler:       _QuestionService_GetAnswersByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FilterQuestion",
			Handler:       _QuestionService_FilterQuestion_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "question.proto",
}
