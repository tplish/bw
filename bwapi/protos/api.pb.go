// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: protos/api.proto

package protos

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type ListItemsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListItemsReq) Reset() {
	*x = ListItemsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsReq) ProtoMessage() {}

func (x *ListItemsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsReq.ProtoReflect.Descriptor instead.
func (*ListItemsReq) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{1}
}

type ListItemsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListItemsResp) Reset() {
	*x = ListItemsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsResp) ProtoMessage() {}

func (x *ListItemsResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsResp.ProtoReflect.Descriptor instead.
func (*ListItemsResp) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListItemsResp) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_protos_api_proto protoreflect.FileDescriptor

var file_protos_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x62, 0x77, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x22, 0x32, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x77, 0x61, 0x70, 0x69,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x5b, 0x0a, 0x05,
	0x42, 0x77, 0x41, 0x70, 0x69, 0x12, 0x52, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x13, 0x2e, 0x62, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x62, 0x77, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_api_proto_rawDescOnce sync.Once
	file_protos_api_proto_rawDescData = file_protos_api_proto_rawDesc
)

func file_protos_api_proto_rawDescGZIP() []byte {
	file_protos_api_proto_rawDescOnce.Do(func() {
		file_protos_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_api_proto_rawDescData)
	})
	return file_protos_api_proto_rawDescData
}

var file_protos_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_api_proto_goTypes = []interface{}{
	(*Item)(nil),          // 0: bwapi.Item
	(*ListItemsReq)(nil),  // 1: bwapi.ListItemsReq
	(*ListItemsResp)(nil), // 2: bwapi.ListItemsResp
}
var file_protos_api_proto_depIdxs = []int32{
	0, // 0: bwapi.ListItemsResp.items:type_name -> bwapi.Item
	1, // 1: bwapi.BwApi.ListItems:input_type -> bwapi.ListItemsReq
	2, // 2: bwapi.BwApi.ListItems:output_type -> bwapi.ListItemsResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_api_proto_init() }
func file_protos_api_proto_init() {
	if File_protos_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_protos_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsReq); i {
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
		file_protos_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsResp); i {
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
			RawDescriptor: file_protos_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_api_proto_goTypes,
		DependencyIndexes: file_protos_api_proto_depIdxs,
		MessageInfos:      file_protos_api_proto_msgTypes,
	}.Build()
	File_protos_api_proto = out.File
	file_protos_api_proto_rawDesc = nil
	file_protos_api_proto_goTypes = nil
	file_protos_api_proto_depIdxs = nil
}
