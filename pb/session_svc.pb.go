// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: session_svc.proto

package pb

import (
	session "github.com/Streamfair/streamfair_session_svc/pb/session"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_session_svc_proto protoreflect.FileDescriptor

var file_session_svc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbd, 0x07, 0x0a, 0x0e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99,
	0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64, 0x20,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61, 0x64,
	0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65,
	0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0xa0, 0x01, 0x0a, 0x0d, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5a, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x32, 0x24, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x91, 0x01,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x92, 0x41, 0x28,
	0x12, 0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64,
	0x7d, 0x12, 0x9a, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x57, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64, 0x20,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61, 0x64,
	0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65,
	0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x2a, 0x24, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x9a,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x57, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x32, 0x24, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61,
	0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x9d, 0x01, 0x0a, 0x0d,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x57, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61,
	0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x42, 0x95, 0x01, 0x92, 0x41,
	0x61, 0x12, 0x5f, 0x0a, 0x1a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x20,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x3c, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x12, 0x1d, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x1a, 0x0f, 0x6e, 0x65,
	0x6c, 0x69, 0x78, 0x40, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x6f, 0x2e, 0x64, 0x65, 0x32, 0x03, 0x31,
	0x2e, 0x30, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x66, 0x61, 0x69, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_session_svc_proto_goTypes = []interface{}{
	(*session.CreateSessionRequest)(nil),  // 0: pb.CreateSessionRequest
	(*session.ExtendSessionRequest)(nil),  // 1: pb.ExtendSessionRequest
	(*session.GetSessionRequest)(nil),     // 2: pb.GetSessionRequest
	(*session.DeleteSessionRequest)(nil),  // 3: pb.DeleteSessionRequest
	(*session.RevokeSessionRequest)(nil),  // 4: pb.RevokeSessionRequest
	(*session.VerifySessionRequest)(nil),  // 5: pb.VerifySessionRequest
	(*session.CreateSessionResponse)(nil), // 6: pb.CreateSessionResponse
	(*session.ExtendSessionResponse)(nil), // 7: pb.ExtendSessionResponse
	(*session.GetSessionResponse)(nil),    // 8: pb.GetSessionResponse
	(*emptypb.Empty)(nil),                 // 9: google.protobuf.Empty
	(*session.VerifySessionResponse)(nil), // 10: pb.VerifySessionResponse
}
var file_session_svc_proto_depIdxs = []int32{
	0,  // 0: pb.SessionService.CreateSession:input_type -> pb.CreateSessionRequest
	1,  // 1: pb.SessionService.ExtendSession:input_type -> pb.ExtendSessionRequest
	2,  // 2: pb.SessionService.GetSession:input_type -> pb.GetSessionRequest
	3,  // 3: pb.SessionService.DeleteSession:input_type -> pb.DeleteSessionRequest
	4,  // 4: pb.SessionService.RevokeSession:input_type -> pb.RevokeSessionRequest
	5,  // 5: pb.SessionService.VerifySession:input_type -> pb.VerifySessionRequest
	6,  // 6: pb.SessionService.CreateSession:output_type -> pb.CreateSessionResponse
	7,  // 7: pb.SessionService.ExtendSession:output_type -> pb.ExtendSessionResponse
	8,  // 8: pb.SessionService.GetSession:output_type -> pb.GetSessionResponse
	9,  // 9: pb.SessionService.DeleteSession:output_type -> google.protobuf.Empty
	9,  // 10: pb.SessionService.RevokeSession:output_type -> google.protobuf.Empty
	10, // 11: pb.SessionService.VerifySession:output_type -> pb.VerifySessionResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_session_svc_proto_init() }
func file_session_svc_proto_init() {
	if File_session_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_session_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_session_svc_proto_goTypes,
		DependencyIndexes: file_session_svc_proto_depIdxs,
	}.Build()
	File_session_svc_proto = out.File
	file_session_svc_proto_rawDesc = nil
	file_session_svc_proto_goTypes = nil
	file_session_svc_proto_depIdxs = nil
}
