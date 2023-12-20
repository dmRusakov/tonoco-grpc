// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tonoco/model/v1/file.proto

package tonoco_model_v1

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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Slug      string `protobuf:"bytes,30,opt,name=slug,proto3" json:"slug,omitempty"`
	Active    bool   `protobuf:"varint,40,opt,name=active,proto3" json:"active,omitempty"`
	Order     int32  `protobuf:"varint,50,opt,name=order,proto3" json:"order,omitempty"`
	CreatedAt int64  `protobuf:"varint,60,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy string `protobuf:"bytes,70,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt int64  `protobuf:"varint,80,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy string `protobuf:"bytes,90,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tonoco_model_v1_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_tonoco_model_v1_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_tonoco_model_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *File) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *File) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *File) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *File) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *File) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *File) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

var File_tonoco_model_v1_file_proto protoreflect.FileDescriptor

var file_tonoco_model_v1_file_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x6f, 0x6e, 0x6f, 0x63, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x6f,
	0x6e, 0x6f, 0x63, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0xe8, 0x01,
	0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6d, 0x52, 0x75, 0x73, 0x61, 0x6b, 0x6f, 0x76,
	0x2f, 0x74, 0x6f, 0x6e, 0x6f, 0x63, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x6f, 0x6e, 0x6f, 0x63, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x6f, 0x6e, 0x6f, 0x63, 0x6f, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tonoco_model_v1_file_proto_rawDescOnce sync.Once
	file_tonoco_model_v1_file_proto_rawDescData = file_tonoco_model_v1_file_proto_rawDesc
)

func file_tonoco_model_v1_file_proto_rawDescGZIP() []byte {
	file_tonoco_model_v1_file_proto_rawDescOnce.Do(func() {
		file_tonoco_model_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_tonoco_model_v1_file_proto_rawDescData)
	})
	return file_tonoco_model_v1_file_proto_rawDescData
}

var file_tonoco_model_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tonoco_model_v1_file_proto_goTypes = []interface{}{
	(*File)(nil), // 0: tonoco.model.v1.File
}
var file_tonoco_model_v1_file_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tonoco_model_v1_file_proto_init() }
func file_tonoco_model_v1_file_proto_init() {
	if File_tonoco_model_v1_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tonoco_model_v1_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_tonoco_model_v1_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tonoco_model_v1_file_proto_goTypes,
		DependencyIndexes: file_tonoco_model_v1_file_proto_depIdxs,
		MessageInfos:      file_tonoco_model_v1_file_proto_msgTypes,
	}.Build()
	File_tonoco_model_v1_file_proto = out.File
	file_tonoco_model_v1_file_proto_rawDesc = nil
	file_tonoco_model_v1_file_proto_goTypes = nil
	file_tonoco_model_v1_file_proto_depIdxs = nil
}
