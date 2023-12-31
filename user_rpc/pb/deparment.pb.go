// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: deparment.proto

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

type Deparment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId       string                 `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Deparment) Reset() {
	*x = Deparment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deparment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deparment) ProtoMessage() {}

func (x *Deparment) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deparment.ProtoReflect.Descriptor instead.
func (*Deparment) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{0}
}

func (x *Deparment) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Deparment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deparment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Deparment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DeparmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId       string                 `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DeparmentResponse) Reset() {
	*x = DeparmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeparmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeparmentResponse) ProtoMessage() {}

func (x *DeparmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeparmentResponse.ProtoReflect.Descriptor instead.
func (*DeparmentResponse) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{1}
}

func (x *DeparmentResponse) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *DeparmentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeparmentResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DeparmentResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateDeparmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateDeparmentRequest) Reset() {
	*x = CreateDeparmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeparmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeparmentRequest) ProtoMessage() {}

func (x *CreateDeparmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeparmentRequest.ProtoReflect.Descriptor instead.
func (*CreateDeparmentRequest) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeparmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeparmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
}

func (x *DeparmentRequest) Reset() {
	*x = DeparmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeparmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeparmentRequest) ProtoMessage() {}

func (x *DeparmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeparmentRequest.ProtoReflect.Descriptor instead.
func (*DeparmentRequest) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{3}
}

func (x *DeparmentRequest) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

type DeparmentUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId  string  `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *DeparmentUpdateRequest) Reset() {
	*x = DeparmentUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeparmentUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeparmentUpdateRequest) ProtoMessage() {}

func (x *DeparmentUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeparmentUpdateRequest.ProtoReflect.Descriptor instead.
func (*DeparmentUpdateRequest) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{4}
}

func (x *DeparmentUpdateRequest) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *DeparmentUpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type GetDeparmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *int64 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetDeparmentsRequest) Reset() {
	*x = GetDeparmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeparmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeparmentsRequest) ProtoMessage() {}

func (x *GetDeparmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeparmentsRequest.ProtoReflect.Descriptor instead.
func (*GetDeparmentsRequest) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeparmentsRequest) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *GetDeparmentsRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type DeleteDeparmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteDeparmentResponse) Reset() {
	*x = DeleteDeparmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeparmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeparmentResponse) ProtoMessage() {}

func (x *DeleteDeparmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeparmentResponse.ProtoReflect.Descriptor instead.
func (*DeleteDeparmentResponse) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDeparmentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DepartmentCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *DepartmentCountResponse) Reset() {
	*x = DepartmentCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deparment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentCountResponse) ProtoMessage() {}

func (x *DepartmentCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deparment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentCountResponse.ProtoReflect.Descriptor instead.
func (*DepartmentCountResponse) Descriptor() ([]byte, []int) {
	return file_deparment_proto_rawDescGZIP(), []int{7}
}

func (x *DepartmentCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_deparment_proto protoreflect.FileDescriptor

var file_deparment_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xae, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23,
	0x0a, 0x10, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x16, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0f, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb6, 0x03, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x46, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x6d,
	0x2d, 0x64, 0x65, 0x76, 0x31, 0x64, 0x65, 0x76, 0x35, 0x2f, 0x6d, 0x74, 0x6d, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2d, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_deparment_proto_rawDescOnce sync.Once
	file_deparment_proto_rawDescData = file_deparment_proto_rawDesc
)

func file_deparment_proto_rawDescGZIP() []byte {
	file_deparment_proto_rawDescOnce.Do(func() {
		file_deparment_proto_rawDescData = protoimpl.X.CompressGZIP(file_deparment_proto_rawDescData)
	})
	return file_deparment_proto_rawDescData
}

var file_deparment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_deparment_proto_goTypes = []interface{}{
	(*Deparment)(nil),               // 0: pb.Deparment
	(*DeparmentResponse)(nil),       // 1: pb.DeparmentResponse
	(*CreateDeparmentRequest)(nil),  // 2: pb.CreateDeparmentRequest
	(*DeparmentRequest)(nil),        // 3: pb.DeparmentRequest
	(*DeparmentUpdateRequest)(nil),  // 4: pb.DeparmentUpdateRequest
	(*GetDeparmentsRequest)(nil),    // 5: pb.GetDeparmentsRequest
	(*DeleteDeparmentResponse)(nil), // 6: pb.DeleteDeparmentResponse
	(*DepartmentCountResponse)(nil), // 7: pb.DepartmentCountResponse
	(*timestamppb.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_deparment_proto_depIdxs = []int32{
	8,  // 0: pb.Deparment.created_at:type_name -> google.protobuf.Timestamp
	8,  // 1: pb.Deparment.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 2: pb.DeparmentResponse.created_at:type_name -> google.protobuf.Timestamp
	8,  // 3: pb.DeparmentResponse.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 4: pb.DeparmentService.CreateDeparment:input_type -> pb.CreateDeparmentRequest
	3,  // 5: pb.DeparmentService.GetDeparment:input_type -> pb.DeparmentRequest
	5,  // 6: pb.DeparmentService.GetDeparments:input_type -> pb.GetDeparmentsRequest
	4,  // 7: pb.DeparmentService.UpdateDeparment:input_type -> pb.DeparmentUpdateRequest
	3,  // 8: pb.DeparmentService.DeleteDeparment:input_type -> pb.DeparmentRequest
	5,  // 9: pb.DeparmentService.GetDepartmentCount:input_type -> pb.GetDeparmentsRequest
	1,  // 10: pb.DeparmentService.CreateDeparment:output_type -> pb.DeparmentResponse
	1,  // 11: pb.DeparmentService.GetDeparment:output_type -> pb.DeparmentResponse
	0,  // 12: pb.DeparmentService.GetDeparments:output_type -> pb.Deparment
	1,  // 13: pb.DeparmentService.UpdateDeparment:output_type -> pb.DeparmentResponse
	6,  // 14: pb.DeparmentService.DeleteDeparment:output_type -> pb.DeleteDeparmentResponse
	7,  // 15: pb.DeparmentService.GetDepartmentCount:output_type -> pb.DepartmentCountResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_deparment_proto_init() }
func file_deparment_proto_init() {
	if File_deparment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deparment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deparment); i {
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
		file_deparment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeparmentResponse); i {
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
		file_deparment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeparmentRequest); i {
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
		file_deparment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeparmentRequest); i {
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
		file_deparment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeparmentUpdateRequest); i {
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
		file_deparment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeparmentsRequest); i {
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
		file_deparment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeparmentResponse); i {
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
		file_deparment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentCountResponse); i {
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
	file_deparment_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_deparment_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deparment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deparment_proto_goTypes,
		DependencyIndexes: file_deparment_proto_depIdxs,
		MessageInfos:      file_deparment_proto_msgTypes,
	}.Build()
	File_deparment_proto = out.File
	file_deparment_proto_rawDesc = nil
	file_deparment_proto_goTypes = nil
	file_deparment_proto_depIdxs = nil
}
