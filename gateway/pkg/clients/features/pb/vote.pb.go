// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: vote.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId        string                 `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	UserId     string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CommentId  string                 `protobuf:"bytes,3,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	QuestionId string                 `protobuf:"bytes,4,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{0}
}

func (x *Vote) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Vote) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Vote) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *Vote) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *Vote) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Vote) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId        string                 `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	UserId     string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CommentId  string                 `protobuf:"bytes,3,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	QuestionId string                 `protobuf:"bytes,4,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{1}
}

func (x *VoteResponse) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *VoteResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VoteResponse) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *VoteResponse) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *VoteResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *VoteResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetVotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *int64 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetVotesRequest) Reset() {
	*x = GetVotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesRequest) ProtoMessage() {}

func (x *GetVotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesRequest.ProtoReflect.Descriptor instead.
func (*GetVotesRequest) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{2}
}

func (x *GetVotesRequest) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *GetVotesRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type CreateVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CommentId  string `protobuf:"bytes,2,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	QuestionId string `protobuf:"bytes,3,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
}

func (x *CreateVoteRequest) Reset() {
	*x = CreateVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteRequest) ProtoMessage() {}

func (x *CreateVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoteRequest.ProtoReflect.Descriptor instead.
func (*CreateVoteRequest) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVoteRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateVoteRequest) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *CreateVoteRequest) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

type UpdateVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId        string  `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	UserId     *string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	CommentId  *string `protobuf:"bytes,3,opt,name=comment_id,json=commentId,proto3,oneof" json:"comment_id,omitempty"`
	QuestionId *string `protobuf:"bytes,4,opt,name=question_id,json=questionId,proto3,oneof" json:"question_id,omitempty"`
}

func (x *UpdateVoteRequest) Reset() {
	*x = UpdateVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoteRequest) ProtoMessage() {}

func (x *UpdateVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateVoteRequest) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateVoteRequest) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *UpdateVoteRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *UpdateVoteRequest) GetCommentId() string {
	if x != nil && x.CommentId != nil {
		return *x.CommentId
	}
	return ""
}

func (x *UpdateVoteRequest) GetQuestionId() string {
	if x != nil && x.QuestionId != nil {
		return *x.QuestionId
	}
	return ""
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{5}
}

func (x *VoteRequest) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

type DeleteVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteVoteResponse) Reset() {
	*x = DeleteVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVoteResponse) ProtoMessage() {}

func (x *DeleteVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVoteResponse.ProtoReflect.Descriptor instead.
func (*DeleteVoteResponse) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteVoteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VoteRequestByUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   *int64 `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit  *int64 `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *VoteRequestByUserId) Reset() {
	*x = VoteRequestByUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequestByUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequestByUserId) ProtoMessage() {}

func (x *VoteRequestByUserId) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequestByUserId.ProtoReflect.Descriptor instead.
func (*VoteRequestByUserId) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{7}
}

func (x *VoteRequestByUserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VoteRequestByUserId) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *VoteRequestByUserId) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type VoteRequestByUserIdQuestionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	QuestionId string `protobuf:"bytes,2,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	Page       *int64 `protobuf:"varint,3,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit      *int64 `protobuf:"varint,4,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *VoteRequestByUserIdQuestionId) Reset() {
	*x = VoteRequestByUserIdQuestionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequestByUserIdQuestionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequestByUserIdQuestionId) ProtoMessage() {}

func (x *VoteRequestByUserIdQuestionId) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequestByUserIdQuestionId.ProtoReflect.Descriptor instead.
func (*VoteRequestByUserIdQuestionId) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{8}
}

func (x *VoteRequestByUserIdQuestionId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VoteRequestByUserIdQuestionId) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *VoteRequestByUserIdQuestionId) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *VoteRequestByUserIdQuestionId) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type VoteCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *VoteCountResponse) Reset() {
	*x = VoteCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteCountResponse) ProtoMessage() {}

func (x *VoteCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteCountResponse.ProtoReflect.Descriptor instead.
func (*VoteCountResponse) Descriptor() ([]byte, []int) {
	return file_vote_proto_rawDescGZIP(), []int{9}
}

func (x *VoteCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_vote_proto protoreflect.FileDescriptor

var file_vote_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xee, 0x01, 0x0a, 0x0c, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x58, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x1e, 0x0a,
	0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0f, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x2e, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x75, 0x0a,
	0x13, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x1d, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x56, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xe3, 0x03, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74,
	0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56,
	0x6f, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x6d, 0x2d, 0x64, 0x65, 0x76, 0x31, 0x64,
	0x65, 0x76, 0x35, 0x2f, 0x6d, 0x74, 0x6d, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2d, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x5f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vote_proto_rawDescOnce sync.Once
	file_vote_proto_rawDescData = file_vote_proto_rawDesc
)

func file_vote_proto_rawDescGZIP() []byte {
	file_vote_proto_rawDescOnce.Do(func() {
		file_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_vote_proto_rawDescData)
	})
	return file_vote_proto_rawDescData
}

var file_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_vote_proto_goTypes = []interface{}{
	(*Vote)(nil),                          // 0: pb.Vote
	(*VoteResponse)(nil),                  // 1: pb.VoteResponse
	(*GetVotesRequest)(nil),               // 2: pb.GetVotesRequest
	(*CreateVoteRequest)(nil),             // 3: pb.CreateVoteRequest
	(*UpdateVoteRequest)(nil),             // 4: pb.UpdateVoteRequest
	(*VoteRequest)(nil),                   // 5: pb.VoteRequest
	(*DeleteVoteResponse)(nil),            // 6: pb.DeleteVoteResponse
	(*VoteRequestByUserId)(nil),           // 7: pb.VoteRequestByUserId
	(*VoteRequestByUserIdQuestionId)(nil), // 8: pb.VoteRequestByUserIdQuestionId
	(*VoteCountResponse)(nil),             // 9: pb.VoteCountResponse
	(*timestamppb.Timestamp)(nil),         // 10: google.protobuf.Timestamp
}
var file_vote_proto_depIdxs = []int32{
	10, // 0: pb.Vote.created_at:type_name -> google.protobuf.Timestamp
	10, // 1: pb.Vote.updated_at:type_name -> google.protobuf.Timestamp
	10, // 2: pb.VoteResponse.created_at:type_name -> google.protobuf.Timestamp
	10, // 3: pb.VoteResponse.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 4: pb.VoteService.CreateVote:input_type -> pb.CreateVoteRequest
	2,  // 5: pb.VoteService.GetVotes:input_type -> pb.GetVotesRequest
	5,  // 6: pb.VoteService.GetVote:input_type -> pb.VoteRequest
	4,  // 7: pb.VoteService.UpdateVote:input_type -> pb.UpdateVoteRequest
	5,  // 8: pb.VoteService.DeleteVote:input_type -> pb.VoteRequest
	7,  // 9: pb.VoteService.GetVotesByUserId:input_type -> pb.VoteRequestByUserId
	7,  // 10: pb.VoteService.GetVoteCount:input_type -> pb.VoteRequestByUserId
	8,  // 11: pb.VoteService.GetVotesByUserIdQuestionId:input_type -> pb.VoteRequestByUserIdQuestionId
	1,  // 12: pb.VoteService.CreateVote:output_type -> pb.VoteResponse
	0,  // 13: pb.VoteService.GetVotes:output_type -> pb.Vote
	1,  // 14: pb.VoteService.GetVote:output_type -> pb.VoteResponse
	1,  // 15: pb.VoteService.UpdateVote:output_type -> pb.VoteResponse
	6,  // 16: pb.VoteService.DeleteVote:output_type -> pb.DeleteVoteResponse
	0,  // 17: pb.VoteService.GetVotesByUserId:output_type -> pb.Vote
	9,  // 18: pb.VoteService.GetVoteCount:output_type -> pb.VoteCountResponse
	0,  // 19: pb.VoteService.GetVotesByUserIdQuestionId:output_type -> pb.Vote
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_vote_proto_init() }
func file_vote_proto_init() {
	if File_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVotesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVoteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequestByUserId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequestByUserIdQuestionId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vote_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteCountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_vote_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_vote_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_vote_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_vote_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vote_proto_goTypes,
		DependencyIndexes: file_vote_proto_depIdxs,
		MessageInfos:      file_vote_proto_msgTypes,
	}.Build()
	File_vote_proto = out.File
	file_vote_proto_rawDesc = nil
	file_vote_proto_goTypes = nil
	file_vote_proto_depIdxs = nil
}
