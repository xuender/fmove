// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: move/target.proto

package move

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

// Targer message.
type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	SubDir    string `protobuf:"bytes,2,opt,name=subDir,proto3" json:"subDir,omitempty"`
	Condition string `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	Note      string `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_move_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_move_target_proto_rawDescGZIP(), []int{0}
}

func (x *Target) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Target) GetSubDir() string {
	if x != nil {
		return x.SubDir
	}
	return ""
}

func (x *Target) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *Target) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

var File_move_target_proto protoreflect.FileDescriptor

var file_move_target_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x76, 0x65, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x06, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x44, 0x69,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x44, 0x69, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_move_target_proto_rawDescOnce sync.Once
	file_move_target_proto_rawDescData = file_move_target_proto_rawDesc
)

func file_move_target_proto_rawDescGZIP() []byte {
	file_move_target_proto_rawDescOnce.Do(func() {
		file_move_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_move_target_proto_rawDescData)
	})
	return file_move_target_proto_rawDescData
}

var file_move_target_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_move_target_proto_goTypes = []interface{}{
	(*Target)(nil), // 0: meta.Target
}
var file_move_target_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_move_target_proto_init() }
func file_move_target_proto_init() {
	if File_move_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_move_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
			RawDescriptor: file_move_target_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_move_target_proto_goTypes,
		DependencyIndexes: file_move_target_proto_depIdxs,
		MessageInfos:      file_move_target_proto_msgTypes,
	}.Build()
	File_move_target_proto = out.File
	file_move_target_proto_rawDesc = nil
	file_move_target_proto_goTypes = nil
	file_move_target_proto_depIdxs = nil
}
