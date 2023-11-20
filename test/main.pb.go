// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: main.proto

package main

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

type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *NewUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewUser) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age       int32      `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Id        int32      `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Followers *Followers `protobuf:"bytes,4,opt,name=followers,proto3" json:"followers,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFollowers() *Followers {
	if x != nil {
		return x.Followers
	}
	return nil
}

type Followers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Youtube int32 `protobuf:"varint,1,opt,name=youtube,proto3" json:"youtube,omitempty"`
	Vk      int32 `protobuf:"varint,2,opt,name=vk,proto3" json:"vk,omitempty"`
}

func (x *Followers) Reset() {
	*x = Followers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Followers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Followers) ProtoMessage() {}

func (x *Followers) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Followers.ProtoReflect.Descriptor instead.
func (*Followers) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{2}
}

func (x *Followers) GetYoutube() int32 {
	if x != nil {
		return x.Youtube
	}
	return 0
}

func (x *Followers) GetVk() int32 {
	if x != nil {
		return x.Vk
	}
	return 0
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x2f, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x09, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x35, 0x0a, 0x09, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x76, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x76, 0x6b, 0x32,
	0x46, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_main_proto_goTypes = []interface{}{
	(*NewUser)(nil),   // 0: usermgmt.NewUser
	(*User)(nil),      // 1: usermgmt.User
	(*Followers)(nil), // 2: usermgmt.Followers
}
var file_main_proto_depIdxs = []int32{
	2, // 0: usermgmt.User.followers:type_name -> usermgmt.Followers
	0, // 1: usermgmt.UserManagement.CreateNewUser:input_type -> usermgmt.NewUser
	1, // 2: usermgmt.UserManagement.CreateNewUser:output_type -> usermgmt.User
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
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
		file_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Followers); i {
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
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}