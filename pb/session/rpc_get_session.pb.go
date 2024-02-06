// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: session/rpc_get_session.proto

package session

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

type GetSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *Uuid `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSessionRequest) Reset() {
	*x = GetSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_rpc_get_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionRequest) ProtoMessage() {}

func (x *GetSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_rpc_get_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionRequest.ProtoReflect.Descriptor instead.
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_rpc_get_session_proto_rawDescGZIP(), []int{0}
}

func (x *GetSessionRequest) GetId() *Uuid {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *GetSessionResponse) Reset() {
	*x = GetSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_rpc_get_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionResponse) ProtoMessage() {}

func (x *GetSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_rpc_get_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionResponse.ProtoReflect.Descriptor instead.
func (*GetSessionResponse) Descriptor() ([]byte, []int) {
	return file_session_rpc_get_session_proto_rawDescGZIP(), []int{1}
}

func (x *GetSessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_session_rpc_get_session_proto protoreflect.FileDescriptor

var file_session_rpc_get_session_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x15, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_rpc_get_session_proto_rawDescOnce sync.Once
	file_session_rpc_get_session_proto_rawDescData = file_session_rpc_get_session_proto_rawDesc
)

func file_session_rpc_get_session_proto_rawDescGZIP() []byte {
	file_session_rpc_get_session_proto_rawDescOnce.Do(func() {
		file_session_rpc_get_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_rpc_get_session_proto_rawDescData)
	})
	return file_session_rpc_get_session_proto_rawDescData
}

var file_session_rpc_get_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_session_rpc_get_session_proto_goTypes = []interface{}{
	(*GetSessionRequest)(nil),  // 0: pb.GetSessionRequest
	(*GetSessionResponse)(nil), // 1: pb.GetSessionResponse
	(*Uuid)(nil),               // 2: pb.Uuid
	(*Session)(nil),            // 3: pb.Session
}
var file_session_rpc_get_session_proto_depIdxs = []int32{
	2, // 0: pb.GetSessionRequest.id:type_name -> pb.Uuid
	3, // 1: pb.GetSessionResponse.session:type_name -> pb.Session
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_session_rpc_get_session_proto_init() }
func file_session_rpc_get_session_proto_init() {
	if File_session_rpc_get_session_proto != nil {
		return
	}
	file_session_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_session_rpc_get_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSessionRequest); i {
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
		file_session_rpc_get_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSessionResponse); i {
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
			RawDescriptor: file_session_rpc_get_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_session_rpc_get_session_proto_goTypes,
		DependencyIndexes: file_session_rpc_get_session_proto_depIdxs,
		MessageInfos:      file_session_rpc_get_session_proto_msgTypes,
	}.Build()
	File_session_rpc_get_session_proto = out.File
	file_session_rpc_get_session_proto_rawDesc = nil
	file_session_rpc_get_session_proto_goTypes = nil
	file_session_rpc_get_session_proto_depIdxs = nil
}
