// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: vote.proto

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
	VoteService_CreateVote_FullMethodName                 = "/pb.VoteService/CreateVote"
	VoteService_GetVotes_FullMethodName                   = "/pb.VoteService/GetVotes"
	VoteService_GetVote_FullMethodName                    = "/pb.VoteService/GetVote"
	VoteService_UpdateVote_FullMethodName                 = "/pb.VoteService/UpdateVote"
	VoteService_DeleteVote_FullMethodName                 = "/pb.VoteService/DeleteVote"
	VoteService_GetVotesByUserId_FullMethodName           = "/pb.VoteService/GetVotesByUserId"
	VoteService_GetVoteCount_FullMethodName               = "/pb.VoteService/GetVoteCount"
	VoteService_GetVotesByUserIdQuestionId_FullMethodName = "/pb.VoteService/GetVotesByUserIdQuestionId"
)

// VoteServiceClient is the client API for VoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoteServiceClient interface {
	CreateVote(ctx context.Context, in *CreateVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	GetVotes(ctx context.Context, in *GetVotesRequest, opts ...grpc.CallOption) (VoteService_GetVotesClient, error)
	GetVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	UpdateVote(ctx context.Context, in *UpdateVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	DeleteVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*DeleteVoteResponse, error)
	GetVotesByUserId(ctx context.Context, in *VoteRequestByUserId, opts ...grpc.CallOption) (VoteService_GetVotesByUserIdClient, error)
	GetVoteCount(ctx context.Context, in *VoteRequestByUserId, opts ...grpc.CallOption) (*VoteCountResponse, error)
	GetVotesByUserIdQuestionId(ctx context.Context, in *VoteRequestByUserIdQuestionId, opts ...grpc.CallOption) (VoteService_GetVotesByUserIdQuestionIdClient, error)
}

type voteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteServiceClient(cc grpc.ClientConnInterface) VoteServiceClient {
	return &voteServiceClient{cc}
}

func (c *voteServiceClient) CreateVote(ctx context.Context, in *CreateVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, VoteService_CreateVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) GetVotes(ctx context.Context, in *GetVotesRequest, opts ...grpc.CallOption) (VoteService_GetVotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &VoteService_ServiceDesc.Streams[0], VoteService_GetVotes_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &voteServiceGetVotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VoteService_GetVotesClient interface {
	Recv() (*Vote, error)
	grpc.ClientStream
}

type voteServiceGetVotesClient struct {
	grpc.ClientStream
}

func (x *voteServiceGetVotesClient) Recv() (*Vote, error) {
	m := new(Vote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *voteServiceClient) GetVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, VoteService_GetVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) UpdateVote(ctx context.Context, in *UpdateVoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, VoteService_UpdateVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) DeleteVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*DeleteVoteResponse, error) {
	out := new(DeleteVoteResponse)
	err := c.cc.Invoke(ctx, VoteService_DeleteVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) GetVotesByUserId(ctx context.Context, in *VoteRequestByUserId, opts ...grpc.CallOption) (VoteService_GetVotesByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &VoteService_ServiceDesc.Streams[1], VoteService_GetVotesByUserId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &voteServiceGetVotesByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VoteService_GetVotesByUserIdClient interface {
	Recv() (*Vote, error)
	grpc.ClientStream
}

type voteServiceGetVotesByUserIdClient struct {
	grpc.ClientStream
}

func (x *voteServiceGetVotesByUserIdClient) Recv() (*Vote, error) {
	m := new(Vote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *voteServiceClient) GetVoteCount(ctx context.Context, in *VoteRequestByUserId, opts ...grpc.CallOption) (*VoteCountResponse, error) {
	out := new(VoteCountResponse)
	err := c.cc.Invoke(ctx, VoteService_GetVoteCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) GetVotesByUserIdQuestionId(ctx context.Context, in *VoteRequestByUserIdQuestionId, opts ...grpc.CallOption) (VoteService_GetVotesByUserIdQuestionIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &VoteService_ServiceDesc.Streams[2], VoteService_GetVotesByUserIdQuestionId_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &voteServiceGetVotesByUserIdQuestionIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VoteService_GetVotesByUserIdQuestionIdClient interface {
	Recv() (*Vote, error)
	grpc.ClientStream
}

type voteServiceGetVotesByUserIdQuestionIdClient struct {
	grpc.ClientStream
}

func (x *voteServiceGetVotesByUserIdQuestionIdClient) Recv() (*Vote, error) {
	m := new(Vote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VoteServiceServer is the server API for VoteService service.
// All implementations must embed UnimplementedVoteServiceServer
// for forward compatibility
type VoteServiceServer interface {
	CreateVote(context.Context, *CreateVoteRequest) (*VoteResponse, error)
	GetVotes(*GetVotesRequest, VoteService_GetVotesServer) error
	GetVote(context.Context, *VoteRequest) (*VoteResponse, error)
	UpdateVote(context.Context, *UpdateVoteRequest) (*VoteResponse, error)
	DeleteVote(context.Context, *VoteRequest) (*DeleteVoteResponse, error)
	GetVotesByUserId(*VoteRequestByUserId, VoteService_GetVotesByUserIdServer) error
	GetVoteCount(context.Context, *VoteRequestByUserId) (*VoteCountResponse, error)
	GetVotesByUserIdQuestionId(*VoteRequestByUserIdQuestionId, VoteService_GetVotesByUserIdQuestionIdServer) error
	mustEmbedUnimplementedVoteServiceServer()
}

// UnimplementedVoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVoteServiceServer struct {
}

func (UnimplementedVoteServiceServer) CreateVote(context.Context, *CreateVoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVote not implemented")
}
func (UnimplementedVoteServiceServer) GetVotes(*GetVotesRequest, VoteService_GetVotesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVotes not implemented")
}
func (UnimplementedVoteServiceServer) GetVote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVote not implemented")
}
func (UnimplementedVoteServiceServer) UpdateVote(context.Context, *UpdateVoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVote not implemented")
}
func (UnimplementedVoteServiceServer) DeleteVote(context.Context, *VoteRequest) (*DeleteVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVote not implemented")
}
func (UnimplementedVoteServiceServer) GetVotesByUserId(*VoteRequestByUserId, VoteService_GetVotesByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVotesByUserId not implemented")
}
func (UnimplementedVoteServiceServer) GetVoteCount(context.Context, *VoteRequestByUserId) (*VoteCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoteCount not implemented")
}
func (UnimplementedVoteServiceServer) GetVotesByUserIdQuestionId(*VoteRequestByUserIdQuestionId, VoteService_GetVotesByUserIdQuestionIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVotesByUserIdQuestionId not implemented")
}
func (UnimplementedVoteServiceServer) mustEmbedUnimplementedVoteServiceServer() {}

// UnsafeVoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoteServiceServer will
// result in compilation errors.
type UnsafeVoteServiceServer interface {
	mustEmbedUnimplementedVoteServiceServer()
}

func RegisterVoteServiceServer(s grpc.ServiceRegistrar, srv VoteServiceServer) {
	s.RegisterService(&VoteService_ServiceDesc, srv)
}

func _VoteService_CreateVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).CreateVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_CreateVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).CreateVote(ctx, req.(*CreateVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_GetVotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetVotesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VoteServiceServer).GetVotes(m, &voteServiceGetVotesServer{stream})
}

type VoteService_GetVotesServer interface {
	Send(*Vote) error
	grpc.ServerStream
}

type voteServiceGetVotesServer struct {
	grpc.ServerStream
}

func (x *voteServiceGetVotesServer) Send(m *Vote) error {
	return x.ServerStream.SendMsg(m)
}

func _VoteService_GetVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).GetVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_GetVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).GetVote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_UpdateVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).UpdateVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_UpdateVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).UpdateVote(ctx, req.(*UpdateVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_DeleteVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).DeleteVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_DeleteVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).DeleteVote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_GetVotesByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VoteRequestByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VoteServiceServer).GetVotesByUserId(m, &voteServiceGetVotesByUserIdServer{stream})
}

type VoteService_GetVotesByUserIdServer interface {
	Send(*Vote) error
	grpc.ServerStream
}

type voteServiceGetVotesByUserIdServer struct {
	grpc.ServerStream
}

func (x *voteServiceGetVotesByUserIdServer) Send(m *Vote) error {
	return x.ServerStream.SendMsg(m)
}

func _VoteService_GetVoteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequestByUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).GetVoteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_GetVoteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).GetVoteCount(ctx, req.(*VoteRequestByUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_GetVotesByUserIdQuestionId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VoteRequestByUserIdQuestionId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VoteServiceServer).GetVotesByUserIdQuestionId(m, &voteServiceGetVotesByUserIdQuestionIdServer{stream})
}

type VoteService_GetVotesByUserIdQuestionIdServer interface {
	Send(*Vote) error
	grpc.ServerStream
}

type voteServiceGetVotesByUserIdQuestionIdServer struct {
	grpc.ServerStream
}

func (x *voteServiceGetVotesByUserIdQuestionIdServer) Send(m *Vote) error {
	return x.ServerStream.SendMsg(m)
}

// VoteService_ServiceDesc is the grpc.ServiceDesc for VoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VoteService",
	HandlerType: (*VoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVote",
			Handler:    _VoteService_CreateVote_Handler,
		},
		{
			MethodName: "GetVote",
			Handler:    _VoteService_GetVote_Handler,
		},
		{
			MethodName: "UpdateVote",
			Handler:    _VoteService_UpdateVote_Handler,
		},
		{
			MethodName: "DeleteVote",
			Handler:    _VoteService_DeleteVote_Handler,
		},
		{
			MethodName: "GetVoteCount",
			Handler:    _VoteService_GetVoteCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetVotes",
			Handler:       _VoteService_GetVotes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVotesByUserId",
			Handler:       _VoteService_GetVotesByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVotesByUserIdQuestionId",
			Handler:       _VoteService_GetVotesByUserIdQuestionId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vote.proto",
}
