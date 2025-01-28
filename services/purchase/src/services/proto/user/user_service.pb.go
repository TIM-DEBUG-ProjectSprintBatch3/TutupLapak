// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: proto/user_service.proto

// Definisi package

package user

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
type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserIds       []string               `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds,omitempty"` // Larik string untuk menerima beberapa ID pengguna
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_proto_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// Response Payload
type UserResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Email             string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`                         // Email pengguna
	Phone             string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`                         // Nomor telepon pengguna
	BankAccountName   string                 `protobuf:"bytes,3,opt,name=bankAccountName,proto3" json:"bankAccountName,omitempty"`     // Nama bank pengguna
	BankAccountHolder string                 `protobuf:"bytes,4,opt,name=bankAccountHolder,proto3" json:"bankAccountHolder,omitempty"` // Nama pemilik akun bank
	BankAccountNumber string                 `protobuf:"bytes,5,opt,name=bankAccountNumber,proto3" json:"bankAccountNumber,omitempty"` // Nomor akun bank
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_proto_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserResponse) GetBankAccountName() string {
	if x != nil {
		return x.BankAccountName
	}
	return ""
}

func (x *UserResponse) GetBankAccountHolder() string {
	if x != nil {
		return x.BankAccountHolder
	}
	return ""
}

func (x *UserResponse) GetBankAccountNumber() string {
	if x != nil {
		return x.BankAccountNumber
	}
	return ""
}

type UserWithIdResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	UserId            string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`                       // Id User
	Email             string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`                         // Email pengguna
	Phone             string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`                         // Nomor telepon pengguna
	BankAccountName   string                 `protobuf:"bytes,4,opt,name=bankAccountName,proto3" json:"bankAccountName,omitempty"`     // Nama bank pengguna
	BankAccountHolder string                 `protobuf:"bytes,5,opt,name=bankAccountHolder,proto3" json:"bankAccountHolder,omitempty"` // Nama pemilik akun bank
	BankAccountNumber string                 `protobuf:"bytes,6,opt,name=bankAccountNumber,proto3" json:"bankAccountNumber,omitempty"` // Nomor akun bank
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UserWithIdResponse) Reset() {
	*x = UserWithIdResponse{}
	mi := &file_proto_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserWithIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWithIdResponse) ProtoMessage() {}

func (x *UserWithIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWithIdResponse.ProtoReflect.Descriptor instead.
func (*UserWithIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserWithIdResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserWithIdResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserWithIdResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserWithIdResponse) GetBankAccountName() string {
	if x != nil {
		return x.BankAccountName
	}
	return ""
}

func (x *UserWithIdResponse) GetBankAccountHolder() string {
	if x != nil {
		return x.BankAccountHolder
	}
	return ""
}

func (x *UserWithIdResponse) GetBankAccountNumber() string {
	if x != nil {
		return x.BankAccountNumber
	}
	return ""
}

// Response untuk banyak pengguna
type UsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserResponse        `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"` // Larik UserResponse untuk beberapa pengguna
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersResponse) Reset() {
	*x = UsersResponse{}
	mi := &file_proto_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersResponse) ProtoMessage() {}

func (x *UsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersResponse.ProtoReflect.Descriptor instead.
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *UsersResponse) GetUsers() []*UserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

type UsersWithIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserWithIdResponse  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"` // Larik UserWithIdResponse untuk beberapa pengguna
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersWithIdResponse) Reset() {
	*x = UsersWithIdResponse{}
	mi := &file_proto_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersWithIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersWithIdResponse) ProtoMessage() {}

func (x *UsersWithIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersWithIdResponse.ProtoReflect.Descriptor instead.
func (*UsersWithIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *UsersWithIdResponse) GetUsers() []*UserWithIdResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_proto_user_service_proto protoreflect.FileDescriptor

var file_proto_user_service_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x27, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x11, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x61, 0x6e, 0x6b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xde, 0x01, 0x0a,
	0x12, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x61,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x11, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x61, 0x6e, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x39, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32,
	0x8d, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x64, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x19, 0x5a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_proto_user_service_proto_rawDescOnce sync.Once
	file_proto_user_service_proto_rawDescData []byte
)

func file_proto_user_service_proto_rawDescGZIP() []byte {
	file_proto_user_service_proto_rawDescOnce.Do(func() {
		file_proto_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_service_proto_rawDesc), len(file_proto_user_service_proto_rawDesc)))
	})
	return file_proto_user_service_proto_rawDescData
}

var file_proto_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_user_service_proto_goTypes = []any{
	(*UserRequest)(nil),         // 0: user.UserRequest
	(*UserResponse)(nil),        // 1: user.UserResponse
	(*UserWithIdResponse)(nil),  // 2: user.UserWithIdResponse
	(*UsersResponse)(nil),       // 3: user.UsersResponse
	(*UsersWithIdResponse)(nil), // 4: user.UsersWithIdResponse
}
var file_proto_user_service_proto_depIdxs = []int32{
	1, // 0: user.UsersResponse.users:type_name -> user.UserResponse
	2, // 1: user.UsersWithIdResponse.users:type_name -> user.UserWithIdResponse
	0, // 2: user.UserService.GetUserDetails:input_type -> user.UserRequest
	0, // 3: user.UserService.GetUserDetailsWithId:input_type -> user.UserRequest
	3, // 4: user.UserService.GetUserDetails:output_type -> user.UsersResponse
	4, // 5: user.UserService.GetUserDetailsWithId:output_type -> user.UsersWithIdResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_user_service_proto_init() }
func file_proto_user_service_proto_init() {
	if File_proto_user_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_service_proto_rawDesc), len(file_proto_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_service_proto_goTypes,
		DependencyIndexes: file_proto_user_service_proto_depIdxs,
		MessageInfos:      file_proto_user_service_proto_msgTypes,
	}.Build()
	File_proto_user_service_proto = out.File
	file_proto_user_service_proto_goTypes = nil
	file_proto_user_service_proto_depIdxs = nil
}
