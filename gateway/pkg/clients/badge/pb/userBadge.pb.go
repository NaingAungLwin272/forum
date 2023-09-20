// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: userBadge.proto

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

type UserBadge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId       string                 `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	UserId    string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BadgeId   string                 `protobuf:"bytes,3,opt,name=badge_id,json=badgeId,proto3" json:"badge_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserBadge) Reset() {
	*x = UserBadge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBadge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBadge) ProtoMessage() {}

func (x *UserBadge) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBadge.ProtoReflect.Descriptor instead.
func (*UserBadge) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{0}
}

func (x *UserBadge) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *UserBadge) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserBadge) GetBadgeId() string {
	if x != nil {
		return x.BadgeId
	}
	return ""
}

func (x *UserBadge) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserBadge) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateUserBadgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BadgeId string `protobuf:"bytes,2,opt,name=badge_id,json=badgeId,proto3" json:"badge_id,omitempty"`
}

func (x *CreateUserBadgeRequest) Reset() {
	*x = CreateUserBadgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserBadgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserBadgeRequest) ProtoMessage() {}

func (x *CreateUserBadgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserBadgeRequest.ProtoReflect.Descriptor instead.
func (*CreateUserBadgeRequest) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserBadgeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateUserBadgeRequest) GetBadgeId() string {
	if x != nil {
		return x.BadgeId
	}
	return ""
}

type UserBadgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BadgeId string `protobuf:"bytes,2,opt,name=badge_id,json=badgeId,proto3" json:"badge_id,omitempty"`
}

func (x *UserBadgeRequest) Reset() {
	*x = UserBadgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBadgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBadgeRequest) ProtoMessage() {}

func (x *UserBadgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBadgeRequest.ProtoReflect.Descriptor instead.
func (*UserBadgeRequest) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{2}
}

func (x *UserBadgeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserBadgeRequest) GetBadgeId() string {
	if x != nil {
		return x.BadgeId
	}
	return ""
}

type UpdateUserBadgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BadgeId *string `protobuf:"bytes,2,opt,name=badge_id,json=badgeId,proto3,oneof" json:"badge_id,omitempty"`
}

func (x *UpdateUserBadgeRequest) Reset() {
	*x = UpdateUserBadgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserBadgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserBadgeRequest) ProtoMessage() {}

func (x *UpdateUserBadgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserBadgeRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserBadgeRequest) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserBadgeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserBadgeRequest) GetBadgeId() string {
	if x != nil && x.BadgeId != nil {
		return *x.BadgeId
	}
	return ""
}

type GetUserBadgesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *int64 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetUserBadgesRequest) Reset() {
	*x = GetUserBadgesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBadgesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBadgesRequest) ProtoMessage() {}

func (x *GetUserBadgesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBadgesRequest.ProtoReflect.Descriptor instead.
func (*GetUserBadgesRequest) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserBadgesRequest) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *GetUserBadgesRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetUserBadgesOfUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserBadgesOfUserRequest) Reset() {
	*x = GetUserBadgesOfUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBadgesOfUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBadgesOfUserRequest) ProtoMessage() {}

func (x *GetUserBadgesOfUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBadgesOfUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserBadgesOfUserRequest) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserBadgesOfUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserBadgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBadge *UserBadge `protobuf:"bytes,1,opt,name=user_badge,json=userBadge,proto3" json:"user_badge,omitempty"`
}

func (x *UserBadgeResponse) Reset() {
	*x = UserBadgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBadgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBadgeResponse) ProtoMessage() {}

func (x *UserBadgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBadgeResponse.ProtoReflect.Descriptor instead.
func (*UserBadgeResponse) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{6}
}

func (x *UserBadgeResponse) GetUserBadge() *UserBadge {
	if x != nil {
		return x.UserBadge
	}
	return nil
}

type UserBadgeResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBadges []*UserBadge `protobuf:"bytes,1,rep,name=user_badges,json=userBadges,proto3" json:"user_badges,omitempty"`
}

func (x *UserBadgeResponseList) Reset() {
	*x = UserBadgeResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBadgeResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBadgeResponseList) ProtoMessage() {}

func (x *UserBadgeResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBadgeResponseList.ProtoReflect.Descriptor instead.
func (*UserBadgeResponseList) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{7}
}

func (x *UserBadgeResponseList) GetUserBadges() []*UserBadge {
	if x != nil {
		return x.UserBadges
	}
	return nil
}

type DeleteUserBadgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteUserBadgeResponse) Reset() {
	*x = DeleteUserBadgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserBadgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserBadgeResponse) ProtoMessage() {}

func (x *DeleteUserBadgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserBadgeResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserBadgeResponse) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserBadgeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BadgeRequestByUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BadgeRequestByUserId) Reset() {
	*x = BadgeRequestByUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadgeRequestByUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadgeRequestByUserId) ProtoMessage() {}

func (x *BadgeRequestByUserId) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadgeRequestByUserId.ProtoReflect.Descriptor instead.
func (*BadgeRequestByUserId) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{9}
}

func (x *BadgeRequestByUserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BadgeCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BadgeCountResponse) Reset() {
	*x = BadgeCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userBadge_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadgeCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadgeCountResponse) ProtoMessage() {}

func (x *BadgeCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userBadge_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadgeCountResponse.ProtoReflect.Descriptor instead.
func (*BadgeCountResponse) Descriptor() ([]byte, []int) {
	return file_userBadge_proto_rawDescGZIP(), []int{10}
}

func (x *BadgeCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_userBadge_proto protoreflect.FileDescriptor

var file_userBadge_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x61, 0x64, 0x67, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x61, 0x64, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x64, 0x67, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x4c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x64, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x64, 0x67, 0x65, 0x49, 0x64, 0x22, 0x46, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x62, 0x61, 0x64, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x61, 0x64,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x61, 0x64, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x35, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x61, 0x64, 0x67, 0x65, 0x73, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61,
	0x64, 0x67, 0x65, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x22, 0x47,
	0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x14,
	0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a,
	0x12, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x8a, 0x04, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67,
	0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x61, 0x64, 0x67, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x73, 0x4f, 0x66,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x73, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61,
	0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x61, 0x64, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x6d, 0x2d, 0x64, 0x65, 0x76, 0x31, 0x64, 0x65, 0x76,
	0x35, 0x2f, 0x6d, 0x74, 0x6d, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2d,
	0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x5f, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userBadge_proto_rawDescOnce sync.Once
	file_userBadge_proto_rawDescData = file_userBadge_proto_rawDesc
)

func file_userBadge_proto_rawDescGZIP() []byte {
	file_userBadge_proto_rawDescOnce.Do(func() {
		file_userBadge_proto_rawDescData = protoimpl.X.CompressGZIP(file_userBadge_proto_rawDescData)
	})
	return file_userBadge_proto_rawDescData
}

var file_userBadge_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_userBadge_proto_goTypes = []interface{}{
	(*UserBadge)(nil),                  // 0: pb.UserBadge
	(*CreateUserBadgeRequest)(nil),     // 1: pb.CreateUserBadgeRequest
	(*UserBadgeRequest)(nil),           // 2: pb.UserBadgeRequest
	(*UpdateUserBadgeRequest)(nil),     // 3: pb.UpdateUserBadgeRequest
	(*GetUserBadgesRequest)(nil),       // 4: pb.GetUserBadgesRequest
	(*GetUserBadgesOfUserRequest)(nil), // 5: pb.GetUserBadgesOfUserRequest
	(*UserBadgeResponse)(nil),          // 6: pb.UserBadgeResponse
	(*UserBadgeResponseList)(nil),      // 7: pb.UserBadgeResponseList
	(*DeleteUserBadgeResponse)(nil),    // 8: pb.DeleteUserBadgeResponse
	(*BadgeRequestByUserId)(nil),       // 9: pb.BadgeRequestByUserId
	(*BadgeCountResponse)(nil),         // 10: pb.BadgeCountResponse
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
}
var file_userBadge_proto_depIdxs = []int32{
	11, // 0: pb.UserBadge.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: pb.UserBadge.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: pb.UserBadgeResponse.user_badge:type_name -> pb.UserBadge
	0,  // 3: pb.UserBadgeResponseList.user_badges:type_name -> pb.UserBadge
	1,  // 4: pb.UserBadgeService.CreateUserBadge:input_type -> pb.CreateUserBadgeRequest
	2,  // 5: pb.UserBadgeService.GetUserBadge:input_type -> pb.UserBadgeRequest
	4,  // 6: pb.UserBadgeService.GetUserBadges:input_type -> pb.GetUserBadgesRequest
	5,  // 7: pb.UserBadgeService.GetUserBadgesOfUser:input_type -> pb.GetUserBadgesOfUserRequest
	3,  // 8: pb.UserBadgeService.UpdateUserBadge:input_type -> pb.UpdateUserBadgeRequest
	2,  // 9: pb.UserBadgeService.DeleteUserBadge:input_type -> pb.UserBadgeRequest
	9,  // 10: pb.UserBadgeService.GetBadgeCount:input_type -> pb.BadgeRequestByUserId
	6,  // 11: pb.UserBadgeService.CreateUserBadge:output_type -> pb.UserBadgeResponse
	6,  // 12: pb.UserBadgeService.GetUserBadge:output_type -> pb.UserBadgeResponse
	7,  // 13: pb.UserBadgeService.GetUserBadges:output_type -> pb.UserBadgeResponseList
	7,  // 14: pb.UserBadgeService.GetUserBadgesOfUser:output_type -> pb.UserBadgeResponseList
	6,  // 15: pb.UserBadgeService.UpdateUserBadge:output_type -> pb.UserBadgeResponse
	8,  // 16: pb.UserBadgeService.DeleteUserBadge:output_type -> pb.DeleteUserBadgeResponse
	10, // 17: pb.UserBadgeService.GetBadgeCount:output_type -> pb.BadgeCountResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_userBadge_proto_init() }
func file_userBadge_proto_init() {
	if File_userBadge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userBadge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBadge); i {
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
		file_userBadge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserBadgeRequest); i {
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
		file_userBadge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBadgeRequest); i {
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
		file_userBadge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserBadgeRequest); i {
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
		file_userBadge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBadgesRequest); i {
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
		file_userBadge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBadgesOfUserRequest); i {
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
		file_userBadge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBadgeResponse); i {
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
		file_userBadge_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBadgeResponseList); i {
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
		file_userBadge_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserBadgeResponse); i {
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
		file_userBadge_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadgeRequestByUserId); i {
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
		file_userBadge_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadgeCountResponse); i {
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
	file_userBadge_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_userBadge_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userBadge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userBadge_proto_goTypes,
		DependencyIndexes: file_userBadge_proto_depIdxs,
		MessageInfos:      file_userBadge_proto_msgTypes,
	}.Build()
	File_userBadge_proto = out.File
	file_userBadge_proto_rawDesc = nil
	file_userBadge_proto_goTypes = nil
	file_userBadge_proto_depIdxs = nil
}
