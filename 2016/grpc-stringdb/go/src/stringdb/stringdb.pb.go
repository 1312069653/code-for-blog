// gRPC service definition for a "string database", with string key/value pairs
// kept on the server.
//
// Eli Bendersky [http://eli.thegreenplace.net]
// This code is in the public domain.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: stringdb.proto

package stringdb

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

type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stringdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_stringdb_proto_rawDescGZIP(), []int{0}
}

func (x *GetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Empty string returned when key not found on the server.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetValueReply) Reset() {
	*x = GetValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueReply) ProtoMessage() {}

func (x *GetValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_stringdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueReply.ProtoReflect.Descriptor instead.
func (*GetValueReply) Descriptor() ([]byte, []int) {
	return file_stringdb_proto_rawDescGZIP(), []int{1}
}

func (x *GetValueReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetValueRequest) Reset() {
	*x = SetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetValueRequest) ProtoMessage() {}

func (x *SetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stringdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetValueRequest.ProtoReflect.Descriptor instead.
func (*SetValueRequest) Descriptor() ([]byte, []int) {
	return file_stringdb_proto_rawDescGZIP(), []int{2}
}

func (x *SetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returns the value.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetValueReply) Reset() {
	*x = SetValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetValueReply) ProtoMessage() {}

func (x *SetValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_stringdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetValueReply.ProtoReflect.Descriptor instead.
func (*SetValueReply) Descriptor() ([]byte, []int) {
	return file_stringdb_proto_rawDescGZIP(), []int{3}
}

func (x *SetValueReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CountValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CountValueRequest) Reset() {
	*x = CountValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountValueRequest) ProtoMessage() {}

func (x *CountValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stringdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountValueRequest.ProtoReflect.Descriptor instead.
func (*CountValueRequest) Descriptor() ([]byte, []int) {
	return file_stringdb_proto_rawDescGZIP(), []int{4}
}

func (x *CountValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type CountValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returns the size of the value, in bytes. If key isn't found on the server,
	// returns -1.
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountValueReply) Reset() {
	*x = CountValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountValueReply) ProtoMessage() {}

func (x *CountValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_stringdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountValueReply.ProtoReflect.Descriptor instead.
func (*CountValueReply) Descriptor() ([]byte, []int) {
	return file_stringdb_proto_rawDescGZIP(), []int{5}
}

func (x *CountValueReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_stringdb_proto protoreflect.FileDescriptor

var file_stringdb_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x64, 0x62, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x27, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd6, 0x01, 0x0a, 0x08, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x44, 0x62, 0x12, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x64, 0x62, 0x2e, 0x53,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0a, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x64, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x64, 0x62, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x64, 0x62, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stringdb_proto_rawDescOnce sync.Once
	file_stringdb_proto_rawDescData = file_stringdb_proto_rawDesc
)

func file_stringdb_proto_rawDescGZIP() []byte {
	file_stringdb_proto_rawDescOnce.Do(func() {
		file_stringdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_stringdb_proto_rawDescData)
	})
	return file_stringdb_proto_rawDescData
}

var file_stringdb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stringdb_proto_goTypes = []interface{}{
	(*GetValueRequest)(nil),   // 0: stringdb.GetValueRequest
	(*GetValueReply)(nil),     // 1: stringdb.GetValueReply
	(*SetValueRequest)(nil),   // 2: stringdb.SetValueRequest
	(*SetValueReply)(nil),     // 3: stringdb.SetValueReply
	(*CountValueRequest)(nil), // 4: stringdb.CountValueRequest
	(*CountValueReply)(nil),   // 5: stringdb.CountValueReply
}
var file_stringdb_proto_depIdxs = []int32{
	0, // 0: stringdb.StringDb.GetValue:input_type -> stringdb.GetValueRequest
	2, // 1: stringdb.StringDb.SetValue:input_type -> stringdb.SetValueRequest
	4, // 2: stringdb.StringDb.CountValue:input_type -> stringdb.CountValueRequest
	1, // 3: stringdb.StringDb.GetValue:output_type -> stringdb.GetValueReply
	3, // 4: stringdb.StringDb.SetValue:output_type -> stringdb.SetValueReply
	5, // 5: stringdb.StringDb.CountValue:output_type -> stringdb.CountValueReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stringdb_proto_init() }
func file_stringdb_proto_init() {
	if File_stringdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stringdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
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
		file_stringdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueReply); i {
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
		file_stringdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetValueRequest); i {
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
		file_stringdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetValueReply); i {
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
		file_stringdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountValueRequest); i {
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
		file_stringdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountValueReply); i {
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
			RawDescriptor: file_stringdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stringdb_proto_goTypes,
		DependencyIndexes: file_stringdb_proto_depIdxs,
		MessageInfos:      file_stringdb_proto_msgTypes,
	}.Build()
	File_stringdb_proto = out.File
	file_stringdb_proto_rawDesc = nil
	file_stringdb_proto_goTypes = nil
	file_stringdb_proto_depIdxs = nil
}
