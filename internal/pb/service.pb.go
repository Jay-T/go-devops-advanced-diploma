// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x1a, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xd9, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x28, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x5f, 0x64,
	0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64,
	0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32,
	0xc2, 0x04, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x73, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x5f,
	0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f,
	0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f,
	0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x73, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x2f, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73,
	0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d,
	0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70,
	0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f,
	0x6d, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c,
	0x6f, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73,
	0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f,
	0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0xad, 0x04, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x6f, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2d, 0x2e, 0x67, 0x6f,
	0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x5f,
	0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f,
	0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x71,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x2d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x66, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x6f,
	0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76,
	0x6f, 0x70, 0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70,
	0x6c, 0x6f, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70,
	0x73, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f,
	0x6d, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f,
	0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x79, 0x2d, 0x54, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x2d, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2d, 0x64, 0x69, 0x70, 0x6c,
	0x6f, 0x6d, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),         // 0: go_devops_advanced_diploma.LoginRequest
	(*RegisterRequest)(nil),      // 1: go_devops_advanced_diploma.RegisterRequest
	(*CreateSecretRequest)(nil),  // 2: go_devops_advanced_diploma.CreateSecretRequest
	(*UpdateSecretRequest)(nil),  // 3: go_devops_advanced_diploma.UpdateSecretRequest
	(*DeleteSecretRequest)(nil),  // 4: go_devops_advanced_diploma.DeleteSecretRequest
	(*GetSecretRequest)(nil),     // 5: go_devops_advanced_diploma.GetSecretRequest
	(*ListSecretRequest)(nil),    // 6: go_devops_advanced_diploma.ListSecretRequest
	(*CreateFileRequest)(nil),    // 7: go_devops_advanced_diploma.CreateFileRequest
	(*UpdateFileRequest)(nil),    // 8: go_devops_advanced_diploma.UpdateFileRequest
	(*DeleteFileRequest)(nil),    // 9: go_devops_advanced_diploma.DeleteFileRequest
	(*GetFileRequest)(nil),       // 10: go_devops_advanced_diploma.GetFileRequest
	(*ListFilesRequest)(nil),     // 11: go_devops_advanced_diploma.ListFilesRequest
	(*LoginResponse)(nil),        // 12: go_devops_advanced_diploma.LoginResponse
	(*RegisterResponse)(nil),     // 13: go_devops_advanced_diploma.RegisterResponse
	(*CreateSecretResponse)(nil), // 14: go_devops_advanced_diploma.CreateSecretResponse
	(*UpdateSecretResponse)(nil), // 15: go_devops_advanced_diploma.UpdateSecretResponse
	(*DeleteSecretResponse)(nil), // 16: go_devops_advanced_diploma.DeleteSecretResponse
	(*GetSecretResponse)(nil),    // 17: go_devops_advanced_diploma.GetSecretResponse
	(*ListSecretResponse)(nil),   // 18: go_devops_advanced_diploma.ListSecretResponse
	(*CreateFileResponse)(nil),   // 19: go_devops_advanced_diploma.CreateFileResponse
	(*UpdateFileResponse)(nil),   // 20: go_devops_advanced_diploma.UpdateFileResponse
	(*DeleteFileResponse)(nil),   // 21: go_devops_advanced_diploma.DeleteFileResponse
	(*GetFileResponse)(nil),      // 22: go_devops_advanced_diploma.GetFileResponse
	(*ListFilesResponse)(nil),    // 23: go_devops_advanced_diploma.ListFilesResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: go_devops_advanced_diploma.Authentication.Login:input_type -> go_devops_advanced_diploma.LoginRequest
	1,  // 1: go_devops_advanced_diploma.Authentication.Register:input_type -> go_devops_advanced_diploma.RegisterRequest
	2,  // 2: go_devops_advanced_diploma.Secret.CreateSecret:input_type -> go_devops_advanced_diploma.CreateSecretRequest
	3,  // 3: go_devops_advanced_diploma.Secret.UpdateSecret:input_type -> go_devops_advanced_diploma.UpdateSecretRequest
	4,  // 4: go_devops_advanced_diploma.Secret.DeleteSecret:input_type -> go_devops_advanced_diploma.DeleteSecretRequest
	5,  // 5: go_devops_advanced_diploma.Secret.GetSecret:input_type -> go_devops_advanced_diploma.GetSecretRequest
	6,  // 6: go_devops_advanced_diploma.Secret.ListSecret:input_type -> go_devops_advanced_diploma.ListSecretRequest
	7,  // 7: go_devops_advanced_diploma.File.CreateFile:input_type -> go_devops_advanced_diploma.CreateFileRequest
	8,  // 8: go_devops_advanced_diploma.File.UpdateFileName:input_type -> go_devops_advanced_diploma.UpdateFileRequest
	9,  // 9: go_devops_advanced_diploma.File.DeleteFile:input_type -> go_devops_advanced_diploma.DeleteFileRequest
	10, // 10: go_devops_advanced_diploma.File.GetFile:input_type -> go_devops_advanced_diploma.GetFileRequest
	11, // 11: go_devops_advanced_diploma.File.ListFiles:input_type -> go_devops_advanced_diploma.ListFilesRequest
	12, // 12: go_devops_advanced_diploma.Authentication.Login:output_type -> go_devops_advanced_diploma.LoginResponse
	13, // 13: go_devops_advanced_diploma.Authentication.Register:output_type -> go_devops_advanced_diploma.RegisterResponse
	14, // 14: go_devops_advanced_diploma.Secret.CreateSecret:output_type -> go_devops_advanced_diploma.CreateSecretResponse
	15, // 15: go_devops_advanced_diploma.Secret.UpdateSecret:output_type -> go_devops_advanced_diploma.UpdateSecretResponse
	16, // 16: go_devops_advanced_diploma.Secret.DeleteSecret:output_type -> go_devops_advanced_diploma.DeleteSecretResponse
	17, // 17: go_devops_advanced_diploma.Secret.GetSecret:output_type -> go_devops_advanced_diploma.GetSecretResponse
	18, // 18: go_devops_advanced_diploma.Secret.ListSecret:output_type -> go_devops_advanced_diploma.ListSecretResponse
	19, // 19: go_devops_advanced_diploma.File.CreateFile:output_type -> go_devops_advanced_diploma.CreateFileResponse
	20, // 20: go_devops_advanced_diploma.File.UpdateFileName:output_type -> go_devops_advanced_diploma.UpdateFileResponse
	21, // 21: go_devops_advanced_diploma.File.DeleteFile:output_type -> go_devops_advanced_diploma.DeleteFileResponse
	22, // 22: go_devops_advanced_diploma.File.GetFile:output_type -> go_devops_advanced_diploma.GetFileResponse
	23, // 23: go_devops_advanced_diploma.File.ListFiles:output_type -> go_devops_advanced_diploma.ListFilesResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_auth_proto_init()
	file_secrets_proto_init()
	file_files_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
