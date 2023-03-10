// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: secrets.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SecretMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SecretMessage) Reset() {
	*x = SecretMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretMessage) ProtoMessage() {}

func (x *SecretMessage) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretMessage.ProtoReflect.Descriptor instead.
func (*SecretMessage) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *SecretMessage) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SecretMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CreateSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SecretMessage `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateSecretRequest) Reset() {
	*x = CreateSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretRequest) ProtoMessage() {}

func (x *CreateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretRequest.ProtoReflect.Descriptor instead.
func (*CreateSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSecretRequest) GetData() *SecretMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SecretMessage `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateSecretResponse) Reset() {
	*x = CreateSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretResponse) ProtoMessage() {}

func (x *CreateSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretResponse.ProtoReflect.Descriptor instead.
func (*CreateSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSecretResponse) GetData() *SecretMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SecretMessage `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateSecretRequest) Reset() {
	*x = UpdateSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSecretRequest) ProtoMessage() {}

func (x *UpdateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSecretRequest.ProtoReflect.Descriptor instead.
func (*UpdateSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSecretRequest) GetData() *SecretMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SecretMessage `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateSecretResponse) Reset() {
	*x = UpdateSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSecretResponse) ProtoMessage() {}

func (x *UpdateSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSecretResponse.ProtoReflect.Descriptor instead.
func (*UpdateSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSecretResponse) GetData() *SecretMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteSecretRequest) Reset() {
	*x = DeleteSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretRequest) ProtoMessage() {}

func (x *DeleteSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretRequest.ProtoReflect.Descriptor instead.
func (*DeleteSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteSecretResponse) Reset() {
	*x = DeleteSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretResponse) ProtoMessage() {}

func (x *DeleteSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretResponse.ProtoReflect.Descriptor instead.
func (*DeleteSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSecretResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetSecretRequest) Reset() {
	*x = GetSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRequest) ProtoMessage() {}

func (x *GetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRequest.ProtoReflect.Descriptor instead.
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{7}
}

func (x *GetSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SecretMessage `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSecretResponse) Reset() {
	*x = GetSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponse) ProtoMessage() {}

func (x *GetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponse.ProtoReflect.Descriptor instead.
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{8}
}

func (x *GetSecretResponse) GetData() *SecretMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ListSecretRequest) Reset() {
	*x = ListSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSecretRequest) ProtoMessage() {}

func (x *ListSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSecretRequest.ProtoReflect.Descriptor instead.
func (*ListSecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{9}
}

func (x *ListSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ListSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*SecretMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListSecretResponse) Reset() {
	*x = ListSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSecretResponse) ProtoMessage() {}

func (x *ListSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSecretResponse.ProtoReflect.Descriptor instead.
func (*ListSecretResponse) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{10}
}

func (x *ListSecretResponse) GetData() []*SecretMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_secrets_proto protoreflect.FileDescriptor

var file_secrets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x22, 0x37, 0x0a, 0x0d, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x54, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x5f, 0x64,
	0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64,
	0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x54, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c,
	0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x28, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x5f,
	0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f,
	0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x53, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c,
	0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x79, 0x2d, 0x54, 0x2f, 0x67, 0x6f, 0x2d, 0x64,
	0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2d, 0x64,
	0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_secrets_proto_rawDescOnce sync.Once
	file_secrets_proto_rawDescData = file_secrets_proto_rawDesc
)

func file_secrets_proto_rawDescGZIP() []byte {
	file_secrets_proto_rawDescOnce.Do(func() {
		file_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_secrets_proto_rawDescData)
	})
	return file_secrets_proto_rawDescData
}

var file_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_secrets_proto_goTypes = []interface{}{
	(*SecretMessage)(nil),        // 0: go_devops_advanced_diploma.SecretMessage
	(*CreateSecretRequest)(nil),  // 1: go_devops_advanced_diploma.CreateSecretRequest
	(*CreateSecretResponse)(nil), // 2: go_devops_advanced_diploma.CreateSecretResponse
	(*UpdateSecretRequest)(nil),  // 3: go_devops_advanced_diploma.UpdateSecretRequest
	(*UpdateSecretResponse)(nil), // 4: go_devops_advanced_diploma.UpdateSecretResponse
	(*DeleteSecretRequest)(nil),  // 5: go_devops_advanced_diploma.DeleteSecretRequest
	(*DeleteSecretResponse)(nil), // 6: go_devops_advanced_diploma.DeleteSecretResponse
	(*GetSecretRequest)(nil),     // 7: go_devops_advanced_diploma.GetSecretRequest
	(*GetSecretResponse)(nil),    // 8: go_devops_advanced_diploma.GetSecretResponse
	(*ListSecretRequest)(nil),    // 9: go_devops_advanced_diploma.ListSecretRequest
	(*ListSecretResponse)(nil),   // 10: go_devops_advanced_diploma.ListSecretResponse
}
var file_secrets_proto_depIdxs = []int32{
	0, // 0: go_devops_advanced_diploma.CreateSecretRequest.data:type_name -> go_devops_advanced_diploma.SecretMessage
	0, // 1: go_devops_advanced_diploma.CreateSecretResponse.data:type_name -> go_devops_advanced_diploma.SecretMessage
	0, // 2: go_devops_advanced_diploma.UpdateSecretRequest.data:type_name -> go_devops_advanced_diploma.SecretMessage
	0, // 3: go_devops_advanced_diploma.UpdateSecretResponse.data:type_name -> go_devops_advanced_diploma.SecretMessage
	0, // 4: go_devops_advanced_diploma.GetSecretResponse.data:type_name -> go_devops_advanced_diploma.SecretMessage
	0, // 5: go_devops_advanced_diploma.ListSecretResponse.data:type_name -> go_devops_advanced_diploma.SecretMessage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_secrets_proto_init() }
func file_secrets_proto_init() {
	if File_secrets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_secrets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretMessage); i {
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
		file_secrets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSecretRequest); i {
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
		file_secrets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSecretResponse); i {
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
		file_secrets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSecretRequest); i {
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
		file_secrets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSecretResponse); i {
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
		file_secrets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSecretRequest); i {
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
		file_secrets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSecretResponse); i {
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
		file_secrets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretRequest); i {
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
		file_secrets_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretResponse); i {
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
		file_secrets_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSecretRequest); i {
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
		file_secrets_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSecretResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_secrets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_secrets_proto_goTypes,
		DependencyIndexes: file_secrets_proto_depIdxs,
		MessageInfos:      file_secrets_proto_msgTypes,
	}.Build()
	File_secrets_proto = out.File
	file_secrets_proto_rawDesc = nil
	file_secrets_proto_goTypes = nil
	file_secrets_proto_depIdxs = nil
}
