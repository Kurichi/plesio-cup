// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: schema/mission.proto

package grpc

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

type GetTreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTreeRequest) Reset() {
	*x = GetTreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreeRequest) ProtoMessage() {}

func (x *GetTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreeRequest.ProtoReflect.Descriptor instead.
func (*GetTreeRequest) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{0}
}

func (x *GetTreeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tree *Tree `protobuf:"bytes,1,opt,name=tree,proto3" json:"tree,omitempty"`
}

func (x *GetTreeResponse) Reset() {
	*x = GetTreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreeResponse) ProtoMessage() {}

func (x *GetTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreeResponse.ProtoReflect.Descriptor instead.
func (*GetTreeResponse) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{1}
}

func (x *GetTreeResponse) GetTree() *Tree {
	if x != nil {
		return x.Tree
	}
	return nil
}

type Tree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Children    []*Tree `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Tree) Reset() {
	*x = Tree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tree) ProtoMessage() {}

func (x *Tree) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tree.ProtoReflect.Descriptor instead.
func (*Tree) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{2}
}

func (x *Tree) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tree) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tree) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tree) GetChildren() []*Tree {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_schema_mission_proto protoreflect.FileDescriptor

var file_schema_mission_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x72, 0x65, 0x65, 0x22, 0x20, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x04, 0x74, 0x72, 0x65,
	0x65, 0x22, 0x74, 0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x32, 0x47, 0x0a, 0x0b, 0x54, 0x72, 0x65, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65,
	0x65, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0a, 0x5a, 0x08, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_mission_proto_rawDescOnce sync.Once
	file_schema_mission_proto_rawDescData = file_schema_mission_proto_rawDesc
)

func file_schema_mission_proto_rawDescGZIP() []byte {
	file_schema_mission_proto_rawDescOnce.Do(func() {
		file_schema_mission_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_mission_proto_rawDescData)
	})
	return file_schema_mission_proto_rawDescData
}

var file_schema_mission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_mission_proto_goTypes = []interface{}{
	(*GetTreeRequest)(nil),  // 0: tree.GetTreeRequest
	(*GetTreeResponse)(nil), // 1: tree.GetTreeResponse
	(*Tree)(nil),            // 2: tree.Tree
}
var file_schema_mission_proto_depIdxs = []int32{
	2, // 0: tree.GetTreeResponse.tree:type_name -> tree.Tree
	2, // 1: tree.Tree.children:type_name -> tree.Tree
	0, // 2: tree.TreeService.GetTree:input_type -> tree.GetTreeRequest
	1, // 3: tree.TreeService.GetTree:output_type -> tree.GetTreeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schema_mission_proto_init() }
func file_schema_mission_proto_init() {
	if File_schema_mission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_mission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTreeRequest); i {
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
		file_schema_mission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTreeResponse); i {
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
		file_schema_mission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tree); i {
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
			RawDescriptor: file_schema_mission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_mission_proto_goTypes,
		DependencyIndexes: file_schema_mission_proto_depIdxs,
		MessageInfos:      file_schema_mission_proto_msgTypes,
	}.Build()
	File_schema_mission_proto = out.File
	file_schema_mission_proto_rawDesc = nil
	file_schema_mission_proto_goTypes = nil
	file_schema_mission_proto_depIdxs = nil
}
