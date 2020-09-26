// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/v1/cat.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetMyCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetMyCatRequest) Reset() {
	*x = GetMyCatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_cat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyCatRequest) ProtoMessage() {}

func (x *GetMyCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_cat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyCatRequest.ProtoReflect.Descriptor instead.
func (*GetMyCatRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_cat_proto_rawDescGZIP(), []int{0}
}

func (x *GetMyCatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MyCatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *MyCatResponse) Reset() {
	*x = MyCatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_cat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyCatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyCatResponse) ProtoMessage() {}

func (x *MyCatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_cat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyCatResponse.ProtoReflect.Descriptor instead.
func (*MyCatResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_cat_proto_rawDescGZIP(), []int{1}
}

func (x *MyCatResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MyCatResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_proto_v1_cat_proto protoreflect.FileDescriptor

var file_proto_v1_cat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4d, 0x79, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x4d, 0x79, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x32, 0x58, 0x0a, 0x03, 0x43,
	0x61, 0x74, 0x12, 0x51, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x61, 0x74, 0x12, 0x18,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x4d, 0x79, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x63, 0x61, 0x74, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x1a, 0x5a, 0x18, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75,
	0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_cat_proto_rawDescOnce sync.Once
	file_proto_v1_cat_proto_rawDescData = file_proto_v1_cat_proto_rawDesc
)

func file_proto_v1_cat_proto_rawDescGZIP() []byte {
	file_proto_v1_cat_proto_rawDescOnce.Do(func() {
		file_proto_v1_cat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_cat_proto_rawDescData)
	})
	return file_proto_v1_cat_proto_rawDescData
}

var file_proto_v1_cat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1_cat_proto_goTypes = []interface{}{
	(*GetMyCatRequest)(nil), // 0: example.GetMyCatRequest
	(*MyCatResponse)(nil),   // 1: example.MyCatResponse
}
var file_proto_v1_cat_proto_depIdxs = []int32{
	0, // 0: example.Cat.GetMyCat:input_type -> example.GetMyCatRequest
	1, // 1: example.Cat.GetMyCat:output_type -> example.MyCatResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_cat_proto_init() }
func file_proto_v1_cat_proto_init() {
	if File_proto_v1_cat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_cat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyCatRequest); i {
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
		file_proto_v1_cat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyCatResponse); i {
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
			RawDescriptor: file_proto_v1_cat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_cat_proto_goTypes,
		DependencyIndexes: file_proto_v1_cat_proto_depIdxs,
		MessageInfos:      file_proto_v1_cat_proto_msgTypes,
	}.Build()
	File_proto_v1_cat_proto = out.File
	file_proto_v1_cat_proto_rawDesc = nil
	file_proto_v1_cat_proto_goTypes = nil
	file_proto_v1_cat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatClient is the client API for Cat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatClient interface {
	GetMyCat(ctx context.Context, in *GetMyCatRequest, opts ...grpc.CallOption) (*MyCatResponse, error)
}

type catClient struct {
	cc grpc.ClientConnInterface
}

func NewCatClient(cc grpc.ClientConnInterface) CatClient {
	return &catClient{cc}
}

func (c *catClient) GetMyCat(ctx context.Context, in *GetMyCatRequest, opts ...grpc.CallOption) (*MyCatResponse, error) {
	out := new(MyCatResponse)
	err := c.cc.Invoke(ctx, "/example.Cat/GetMyCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatServer is the server API for Cat service.
type CatServer interface {
	GetMyCat(context.Context, *GetMyCatRequest) (*MyCatResponse, error)
}

// UnimplementedCatServer can be embedded to have forward compatible implementations.
type UnimplementedCatServer struct {
}

func (*UnimplementedCatServer) GetMyCat(context.Context, *GetMyCatRequest) (*MyCatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCat not implemented")
}

func RegisterCatServer(s *grpc.Server, srv CatServer) {
	s.RegisterService(&_Cat_serviceDesc, srv)
}

func _Cat_GetMyCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).GetMyCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Cat/GetMyCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).GetMyCat(ctx, req.(*GetMyCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Cat",
	HandlerType: (*CatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyCat",
			Handler:    _Cat_GetMyCat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/cat.proto",
}
