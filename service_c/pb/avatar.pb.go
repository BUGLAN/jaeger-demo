// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: avatar.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetAvatarURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetAvatarURLRequest) Reset() {
	*x = GetAvatarURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarURLRequest) ProtoMessage() {}

func (x *GetAvatarURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarURLRequest.ProtoReflect.Descriptor instead.
func (*GetAvatarURLRequest) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{0}
}

func (x *GetAvatarURLRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetAvatarURLReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetAvatarURLReply) Reset() {
	*x = GetAvatarURLReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarURLReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarURLReply) ProtoMessage() {}

func (x *GetAvatarURLReply) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarURLReply.ProtoReflect.Descriptor instead.
func (*GetAvatarURLReply) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{1}
}

func (x *GetAvatarURLReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_avatar_proto protoreflect.FileDescriptor

var file_avatar_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x4a, 0x0a, 0x06,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_avatar_proto_rawDescOnce sync.Once
	file_avatar_proto_rawDescData = file_avatar_proto_rawDesc
)

func file_avatar_proto_rawDescGZIP() []byte {
	file_avatar_proto_rawDescOnce.Do(func() {
		file_avatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_avatar_proto_rawDescData)
	})
	return file_avatar_proto_rawDescData
}

var file_avatar_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_avatar_proto_goTypes = []interface{}{
	(*GetAvatarURLRequest)(nil), // 0: pb.GetAvatarURLRequest
	(*GetAvatarURLReply)(nil),   // 1: pb.GetAvatarURLReply
}
var file_avatar_proto_depIdxs = []int32{
	0, // 0: pb.Avatar.GetAvatarURL:input_type -> pb.GetAvatarURLRequest
	1, // 1: pb.Avatar.GetAvatarURL:output_type -> pb.GetAvatarURLReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_avatar_proto_init() }
func file_avatar_proto_init() {
	if File_avatar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_avatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvatarURLRequest); i {
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
		file_avatar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvatarURLReply); i {
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
			RawDescriptor: file_avatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avatar_proto_goTypes,
		DependencyIndexes: file_avatar_proto_depIdxs,
		MessageInfos:      file_avatar_proto_msgTypes,
	}.Build()
	File_avatar_proto = out.File
	file_avatar_proto_rawDesc = nil
	file_avatar_proto_goTypes = nil
	file_avatar_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AvatarClient is the client API for Avatar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AvatarClient interface {
	GetAvatarURL(ctx context.Context, in *GetAvatarURLRequest, opts ...grpc.CallOption) (*GetAvatarURLReply, error)
}

type avatarClient struct {
	cc grpc.ClientConnInterface
}

func NewAvatarClient(cc grpc.ClientConnInterface) AvatarClient {
	return &avatarClient{cc}
}

func (c *avatarClient) GetAvatarURL(ctx context.Context, in *GetAvatarURLRequest, opts ...grpc.CallOption) (*GetAvatarURLReply, error) {
	out := new(GetAvatarURLReply)
	err := c.cc.Invoke(ctx, "/pb.Avatar/GetAvatarURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvatarServer is the server API for Avatar service.
type AvatarServer interface {
	GetAvatarURL(context.Context, *GetAvatarURLRequest) (*GetAvatarURLReply, error)
}

// UnimplementedAvatarServer can be embedded to have forward compatible implementations.
type UnimplementedAvatarServer struct {
}

func (*UnimplementedAvatarServer) GetAvatarURL(context.Context, *GetAvatarURLRequest) (*GetAvatarURLReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatarURL not implemented")
}

func RegisterAvatarServer(s *grpc.Server, srv AvatarServer) {
	s.RegisterService(&_Avatar_serviceDesc, srv)
}

func _Avatar_GetAvatarURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvatarURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServer).GetAvatarURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Avatar/GetAvatarURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServer).GetAvatarURL(ctx, req.(*GetAvatarURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Avatar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Avatar",
	HandlerType: (*AvatarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvatarURL",
			Handler:    _Avatar_GetAvatarURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "avatar.proto",
}
