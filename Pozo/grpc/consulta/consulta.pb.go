// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: consulta.proto

package __

import (
	context "context"
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

// Formato esperado desde clientes, Lider.go
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consulta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_consulta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_consulta_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_consulta_proto protoreflect.FileDescriptor

var file_consulta_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x49, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0c,
	0x50, 0x65, 0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x7a, 0x6f, 0x12, 0x11, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_consulta_proto_rawDescOnce sync.Once
	file_consulta_proto_rawDescData = file_consulta_proto_rawDesc
)

func file_consulta_proto_rawDescGZIP() []byte {
	file_consulta_proto_rawDescOnce.Do(func() {
		file_consulta_proto_rawDescData = protoimpl.X.CompressGZIP(file_consulta_proto_rawDescData)
	})
	return file_consulta_proto_rawDescData
}

var file_consulta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_consulta_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: consulta.Message
}
var file_consulta_proto_depIdxs = []int32{
	0, // 0: consulta.ConsultaService.PeticionPozo:input_type -> consulta.Message
	0, // 1: consulta.ConsultaService.PeticionPozo:output_type -> consulta.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_consulta_proto_init() }
func file_consulta_proto_init() {
	if File_consulta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consulta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_consulta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consulta_proto_goTypes,
		DependencyIndexes: file_consulta_proto_depIdxs,
		MessageInfos:      file_consulta_proto_msgTypes,
	}.Build()
	File_consulta_proto = out.File
	file_consulta_proto_rawDesc = nil
	file_consulta_proto_goTypes = nil
	file_consulta_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsultaServiceClient is the client API for ConsultaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsultaServiceClient interface {
	PeticionPozo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type consultaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsultaServiceClient(cc grpc.ClientConnInterface) ConsultaServiceClient {
	return &consultaServiceClient{cc}
}

func (c *consultaServiceClient) PeticionPozo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/consulta.ConsultaService/PeticionPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsultaServiceServer is the server API for ConsultaService service.
type ConsultaServiceServer interface {
	PeticionPozo(context.Context, *Message) (*Message, error)
}

// UnimplementedConsultaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsultaServiceServer struct {
}

func (*UnimplementedConsultaServiceServer) PeticionPozo(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeticionPozo not implemented")
}

func RegisterConsultaServiceServer(s *grpc.Server, srv ConsultaServiceServer) {
	s.RegisterService(&_ConsultaService_serviceDesc, srv)
}

func _ConsultaService_PeticionPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsultaServiceServer).PeticionPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consulta.ConsultaService/PeticionPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsultaServiceServer).PeticionPozo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsultaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consulta.ConsultaService",
	HandlerType: (*ConsultaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PeticionPozo",
			Handler:    _ConsultaService_PeticionPozo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consulta.proto",
}
