// version proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: src/grpc/proto/product.proto

// definisi package

package product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request Payload
type ProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	mi := &file_src_grpc_proto_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_proto_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_src_grpc_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

// Response Payload
type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Qty           string                 `protobuf:"bytes,3,opt,name=Qty,proto3" json:"Qty,omitempty"`
	Price         string                 `protobuf:"bytes,4,opt,name=Price,proto3" json:"Price,omitempty"`
	Sku           string                 `protobuf:"bytes,5,opt,name=Sku,proto3" json:"Sku,omitempty"`
	FileId        string                 `protobuf:"bytes,6,opt,name=FileId,proto3" json:"FileId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_src_grpc_proto_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_proto_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_src_grpc_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *ProductResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ProductResponse) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ProductResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

var File_src_grpc_proto_product_proto protoreflect.FileDescriptor

var file_src_grpc_proto_product_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x2e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x51, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x51, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x53, 0x6b, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x32,
	0x5b, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_src_grpc_proto_product_proto_rawDescOnce sync.Once
	file_src_grpc_proto_product_proto_rawDescData []byte
)

func file_src_grpc_proto_product_proto_rawDescGZIP() []byte {
	file_src_grpc_proto_product_proto_rawDescOnce.Do(func() {
		file_src_grpc_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_src_grpc_proto_product_proto_rawDesc), len(file_src_grpc_proto_product_proto_rawDesc)))
	})
	return file_src_grpc_proto_product_proto_rawDescData
}

var file_src_grpc_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_grpc_proto_product_proto_goTypes = []any{
	(*ProductRequest)(nil),  // 0: product.ProductRequest
	(*ProductResponse)(nil), // 1: product.ProductResponse
}
var file_src_grpc_proto_product_proto_depIdxs = []int32{
	0, // 0: product.ProductService.GetProductDetailById:input_type -> product.ProductRequest
	1, // 1: product.ProductService.GetProductDetailById:output_type -> product.ProductResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_grpc_proto_product_proto_init() }
func file_src_grpc_proto_product_proto_init() {
	if File_src_grpc_proto_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_src_grpc_proto_product_proto_rawDesc), len(file_src_grpc_proto_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_grpc_proto_product_proto_goTypes,
		DependencyIndexes: file_src_grpc_proto_product_proto_depIdxs,
		MessageInfos:      file_src_grpc_proto_product_proto_msgTypes,
	}.Build()
	File_src_grpc_proto_product_proto = out.File
	file_src_grpc_proto_product_proto_goTypes = nil
	file_src_grpc_proto_product_proto_depIdxs = nil
}
