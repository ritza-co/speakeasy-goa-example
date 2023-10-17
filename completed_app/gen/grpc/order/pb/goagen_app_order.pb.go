// Code generated with goa v3.13.2, DO NOT EDIT.
//
// order protocol buffer definition
//
// Command:
// $ goa gen app/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: goagen_app_order.proto

package orderpb

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

type TeaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to have green tea instead of normal.
	IsGreen *bool `protobuf:"varint,1,opt,name=is_green,json=isGreen,proto3,oneof" json:"is_green,omitempty"`
	// Number of spoons of sugar.
	NumberSugars *int32 `protobuf:"zigzag32,2,opt,name=number_sugars,json=numberSugars,proto3,oneof" json:"number_sugars,omitempty"`
	// Whether to have milk.
	IncludeMilk *bool `protobuf:"varint,3,opt,name=include_milk,json=includeMilk,proto3,oneof" json:"include_milk,omitempty"`
}

func (x *TeaRequest) Reset() {
	*x = TeaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_app_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeaRequest) ProtoMessage() {}

func (x *TeaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_app_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeaRequest.ProtoReflect.Descriptor instead.
func (*TeaRequest) Descriptor() ([]byte, []int) {
	return file_goagen_app_order_proto_rawDescGZIP(), []int{0}
}

func (x *TeaRequest) GetIsGreen() bool {
	if x != nil && x.IsGreen != nil {
		return *x.IsGreen
	}
	return false
}

func (x *TeaRequest) GetNumberSugars() int32 {
	if x != nil && x.NumberSugars != nil {
		return *x.NumberSugars
	}
	return 0
}

func (x *TeaRequest) GetIncludeMilk() bool {
	if x != nil && x.IncludeMilk != nil {
		return *x.IncludeMilk
	}
	return false
}

type TeaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *TeaResponse) Reset() {
	*x = TeaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_app_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeaResponse) ProtoMessage() {}

func (x *TeaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_app_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeaResponse.ProtoReflect.Descriptor instead.
func (*TeaResponse) Descriptor() ([]byte, []int) {
	return file_goagen_app_order_proto_rawDescGZIP(), []int{1}
}

func (x *TeaResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_goagen_app_order_proto protoreflect.FileDescriptor

var file_goagen_app_order_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x6f, 0x61, 0x67, 0x65, 0x6e, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0xae, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x07, 0x69, 0x73, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x28,
	0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x67, 0x61, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x11, 0x48, 0x01, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x53,
	0x75, 0x67, 0x61, 0x72, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02,
	0x52, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4d, 0x69, 0x6c, 0x6b, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x5f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x67, 0x61, 0x72, 0x73, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6b,
	0x22, 0x23, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x35, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x03, 0x54, 0x65, 0x61, 0x12, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x65,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x54, 0x65, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goagen_app_order_proto_rawDescOnce sync.Once
	file_goagen_app_order_proto_rawDescData = file_goagen_app_order_proto_rawDesc
)

func file_goagen_app_order_proto_rawDescGZIP() []byte {
	file_goagen_app_order_proto_rawDescOnce.Do(func() {
		file_goagen_app_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_goagen_app_order_proto_rawDescData)
	})
	return file_goagen_app_order_proto_rawDescData
}

var file_goagen_app_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goagen_app_order_proto_goTypes = []interface{}{
	(*TeaRequest)(nil),  // 0: order.TeaRequest
	(*TeaResponse)(nil), // 1: order.TeaResponse
}
var file_goagen_app_order_proto_depIdxs = []int32{
	0, // 0: order.Order.Tea:input_type -> order.TeaRequest
	1, // 1: order.Order.Tea:output_type -> order.TeaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goagen_app_order_proto_init() }
func file_goagen_app_order_proto_init() {
	if File_goagen_app_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goagen_app_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeaRequest); i {
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
		file_goagen_app_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeaResponse); i {
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
	file_goagen_app_order_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goagen_app_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_app_order_proto_goTypes,
		DependencyIndexes: file_goagen_app_order_proto_depIdxs,
		MessageInfos:      file_goagen_app_order_proto_msgTypes,
	}.Build()
	File_goagen_app_order_proto = out.File
	file_goagen_app_order_proto_rawDesc = nil
	file_goagen_app_order_proto_goTypes = nil
	file_goagen_app_order_proto_depIdxs = nil
}
