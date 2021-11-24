// Client Service Protobuf file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: clients/client.proto

package clients

import (
	types "github.com/SSH-Management/server/proto/server/types"
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

type CreateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ip    string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateClientRequest) Reset() {
	*x = CreateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRequest) ProtoMessage() {}

func (x *CreateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRequest.ProtoReflect.Descriptor instead.
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return file_clients_client_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClientRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClientRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CreateClientRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type CreateClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Users     []*types.User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	PublicKey string        `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *CreateClientResponse) Reset() {
	*x = CreateClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientResponse) ProtoMessage() {}

func (x *CreateClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clients_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientResponse.ProtoReflect.Descriptor instead.
func (*CreateClientResponse) Descriptor() ([]byte, []int) {
	return file_clients_client_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClientResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateClientResponse) GetUsers() []*types.User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *CreateClientResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type DeleteClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteClientRequest) Reset() {
	*x = DeleteClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientRequest) ProtoMessage() {}

func (x *DeleteClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientRequest.ProtoReflect.Descriptor instead.
func (*DeleteClientRequest) Descriptor() ([]byte, []int) {
	return file_clients_client_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteClientRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClientResponse) Reset() {
	*x = DeleteClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clients_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientResponse) ProtoMessage() {}

func (x *DeleteClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clients_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientResponse.ProtoReflect.Descriptor instead.
func (*DeleteClientResponse) Descriptor() ([]byte, []int) {
	return file_clients_client_proto_rawDescGZIP(), []int{3}
}

var File_clients_client_proto protoreflect.FileDescriptor

var file_clients_client_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x68, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x25, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9d, 0x01, 0x0a, 0x0d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x53, 0x48, 0x2d, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clients_client_proto_rawDescOnce sync.Once
	file_clients_client_proto_rawDescData = file_clients_client_proto_rawDesc
)

func file_clients_client_proto_rawDescGZIP() []byte {
	file_clients_client_proto_rawDescOnce.Do(func() {
		file_clients_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_clients_client_proto_rawDescData)
	})
	return file_clients_client_proto_rawDescData
}

var file_clients_client_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_clients_client_proto_goTypes = []interface{}{
	(*CreateClientRequest)(nil),  // 0: clients.CreateClientRequest
	(*CreateClientResponse)(nil), // 1: clients.CreateClientResponse
	(*DeleteClientRequest)(nil),  // 2: clients.DeleteClientRequest
	(*DeleteClientResponse)(nil), // 3: clients.DeleteClientResponse
	(*types.User)(nil),           // 4: types.User
}
var file_clients_client_proto_depIdxs = []int32{
	4, // 0: clients.CreateClientResponse.users:type_name -> types.User
	0, // 1: clients.ClientService.Create:input_type -> clients.CreateClientRequest
	2, // 2: clients.ClientService.Delete:input_type -> clients.DeleteClientRequest
	1, // 3: clients.ClientService.Create:output_type -> clients.CreateClientResponse
	3, // 4: clients.ClientService.Delete:output_type -> clients.DeleteClientResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_clients_client_proto_init() }
func file_clients_client_proto_init() {
	if File_clients_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clients_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientRequest); i {
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
		file_clients_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientResponse); i {
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
		file_clients_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientRequest); i {
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
		file_clients_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientResponse); i {
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
			RawDescriptor: file_clients_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clients_client_proto_goTypes,
		DependencyIndexes: file_clients_client_proto_depIdxs,
		MessageInfos:      file_clients_client_proto_msgTypes,
	}.Build()
	File_clients_client_proto = out.File
	file_clients_client_proto_rawDesc = nil
	file_clients_client_proto_goTypes = nil
	file_clients_client_proto_depIdxs = nil
}
