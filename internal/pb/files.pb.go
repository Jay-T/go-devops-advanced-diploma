// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: files.proto

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

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Ready    *bool  `protobuf:"varint,3,opt,name=ready,proto3,oneof" json:"ready,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileInfo) GetReady() bool {
	if x != nil && x.Ready != nil {
		return *x.Ready
	}
	return false
}

type CreateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*CreateFileRequest_Info
	//	*CreateFileRequest_ChunkData
	Data isCreateFileRequest_Data `protobuf_oneof:"data"`
}

func (x *CreateFileRequest) Reset() {
	*x = CreateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileRequest) ProtoMessage() {}

func (x *CreateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileRequest.ProtoReflect.Descriptor instead.
func (*CreateFileRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{1}
}

func (m *CreateFileRequest) GetData() isCreateFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CreateFileRequest) GetInfo() *FileInfo {
	if x, ok := x.GetData().(*CreateFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *CreateFileRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*CreateFileRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isCreateFileRequest_Data interface {
	isCreateFileRequest_Data()
}

type CreateFileRequest_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type CreateFileRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*CreateFileRequest_Info) isCreateFileRequest_Data() {}

func (*CreateFileRequest_ChunkData) isCreateFileRequest_Data() {}

type CreateFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Size uint32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CreateFileResponse) Reset() {
	*x = CreateFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileResponse) ProtoMessage() {}

func (x *CreateFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileResponse.ProtoReflect.Descriptor instead.
func (*CreateFileResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFileResponse) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *CreateFileResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type UpdateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UpdateFileRequest_Info
	//	*UpdateFileRequest_ChunkData
	Data isUpdateFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UpdateFileRequest) Reset() {
	*x = UpdateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequest) ProtoMessage() {}

func (x *UpdateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{3}
}

func (m *UpdateFileRequest) GetData() isUpdateFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UpdateFileRequest) GetInfo() *FileInfo {
	if x, ok := x.GetData().(*UpdateFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UpdateFileRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UpdateFileRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUpdateFileRequest_Data interface {
	isUpdateFileRequest_Data()
}

type UpdateFileRequest_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UpdateFileRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UpdateFileRequest_Info) isUpdateFileRequest_Data() {}

func (*UpdateFileRequest_ChunkData) isUpdateFileRequest_Data() {}

type UpdateFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *UpdateFileResponse) Reset() {
	*x = UpdateFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileResponse) ProtoMessage() {}

func (x *UpdateFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileResponse.ProtoReflect.Descriptor instead.
func (*UpdateFileResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFileResponse) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFileRequest) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFileResponse) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *FileInfo `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{7}
}

func (x *GetFileRequest) GetKey() *FileInfo {
	if x != nil {
		return x.Key
	}
	return nil
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*GetFileResponse_Info
	//	*GetFileResponse_ChunkData
	Data isGetFileResponse_Data `protobuf_oneof:"data"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{8}
}

func (m *GetFileResponse) GetData() isGetFileResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *GetFileResponse) GetInfo() *FileInfo {
	if x, ok := x.GetData().(*GetFileResponse_Info); ok {
		return x.Info
	}
	return nil
}

func (x *GetFileResponse) GetChunkData() []byte {
	if x, ok := x.GetData().(*GetFileResponse_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isGetFileResponse_Data interface {
	isGetFileResponse_Data()
}

type GetFileResponse_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type GetFileResponse_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*GetFileResponse_Info) isGetFileResponse_Data() {}

func (*GetFileResponse_ChunkData) isGetFileResponse_Data() {}

type ListFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ListFileRequest) Reset() {
	*x = ListFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileRequest) ProtoMessage() {}

func (x *ListFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileRequest.ProtoReflect.Descriptor instead.
func (*ListFileRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{9}
}

func (x *ListFileRequest) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ListFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*FileInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *ListFileResponse) Reset() {
	*x = ListFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileResponse) ProtoMessage() {}

func (x *ListFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileResponse.ProtoReflect.Descriptor instead.
func (*ListFileResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{10}
}

func (x *ListFileResponse) GetInfo() []*FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_files_proto protoreflect.FileDescriptor

var file_files_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67,
	0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x22, 0x67, 0x0a, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x05,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x22, 0x78, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70,
	0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f,
	0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x62, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x78, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f,
	0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4e, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65,
	0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69,
	0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x76, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73,
	0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d,
	0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4b, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67,
	0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x5f,
	0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f,
	0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x79, 0x2d, 0x54, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65,
	0x76, 0x6f, 0x70, 0x73, 0x2d, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2d, 0x64, 0x69,
	0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_files_proto_rawDescOnce sync.Once
	file_files_proto_rawDescData = file_files_proto_rawDesc
)

func file_files_proto_rawDescGZIP() []byte {
	file_files_proto_rawDescOnce.Do(func() {
		file_files_proto_rawDescData = protoimpl.X.CompressGZIP(file_files_proto_rawDescData)
	})
	return file_files_proto_rawDescData
}

var file_files_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_files_proto_goTypes = []interface{}{
	(*FileInfo)(nil),           // 0: go_devops_advanced_diploma.FileInfo
	(*CreateFileRequest)(nil),  // 1: go_devops_advanced_diploma.CreateFileRequest
	(*CreateFileResponse)(nil), // 2: go_devops_advanced_diploma.CreateFileResponse
	(*UpdateFileRequest)(nil),  // 3: go_devops_advanced_diploma.UpdateFileRequest
	(*UpdateFileResponse)(nil), // 4: go_devops_advanced_diploma.UpdateFileResponse
	(*DeleteFileRequest)(nil),  // 5: go_devops_advanced_diploma.DeleteFileRequest
	(*DeleteFileResponse)(nil), // 6: go_devops_advanced_diploma.DeleteFileResponse
	(*GetFileRequest)(nil),     // 7: go_devops_advanced_diploma.GetFileRequest
	(*GetFileResponse)(nil),    // 8: go_devops_advanced_diploma.GetFileResponse
	(*ListFileRequest)(nil),    // 9: go_devops_advanced_diploma.ListFileRequest
	(*ListFileResponse)(nil),   // 10: go_devops_advanced_diploma.ListFileResponse
}
var file_files_proto_depIdxs = []int32{
	0,  // 0: go_devops_advanced_diploma.CreateFileRequest.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 1: go_devops_advanced_diploma.CreateFileResponse.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 2: go_devops_advanced_diploma.UpdateFileRequest.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 3: go_devops_advanced_diploma.UpdateFileResponse.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 4: go_devops_advanced_diploma.DeleteFileRequest.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 5: go_devops_advanced_diploma.DeleteFileResponse.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 6: go_devops_advanced_diploma.GetFileRequest.key:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 7: go_devops_advanced_diploma.GetFileResponse.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 8: go_devops_advanced_diploma.ListFileRequest.info:type_name -> go_devops_advanced_diploma.FileInfo
	0,  // 9: go_devops_advanced_diploma.ListFileResponse.info:type_name -> go_devops_advanced_diploma.FileInfo
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_files_proto_init() }
func file_files_proto_init() {
	if File_files_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_files_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_files_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileRequest); i {
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
		file_files_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileResponse); i {
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
		file_files_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileRequest); i {
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
		file_files_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileResponse); i {
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
		file_files_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
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
		file_files_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileResponse); i {
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
		file_files_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_files_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
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
		file_files_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileRequest); i {
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
		file_files_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileResponse); i {
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
	file_files_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_files_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CreateFileRequest_Info)(nil),
		(*CreateFileRequest_ChunkData)(nil),
	}
	file_files_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*UpdateFileRequest_Info)(nil),
		(*UpdateFileRequest_ChunkData)(nil),
	}
	file_files_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*GetFileResponse_Info)(nil),
		(*GetFileResponse_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_files_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_files_proto_goTypes,
		DependencyIndexes: file_files_proto_depIdxs,
		MessageInfos:      file_files_proto_msgTypes,
	}.Build()
	File_files_proto = out.File
	file_files_proto_rawDesc = nil
	file_files_proto_goTypes = nil
	file_files_proto_depIdxs = nil
}