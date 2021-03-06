// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SendImageRequest struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendImageRequest) Reset()         { *m = SendImageRequest{} }
func (m *SendImageRequest) String() string { return proto.CompactTextString(m) }
func (*SendImageRequest) ProtoMessage()    {}
func (*SendImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *SendImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendImageRequest.Unmarshal(m, b)
}
func (m *SendImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendImageRequest.Marshal(b, m, deterministic)
}
func (m *SendImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendImageRequest.Merge(m, src)
}
func (m *SendImageRequest) XXX_Size() int {
	return xxx_messageInfo_SendImageRequest.Size(m)
}
func (m *SendImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendImageRequest proto.InternalMessageInfo

func (m *SendImageRequest) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type SendImageResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendImageResponse) Reset()         { *m = SendImageResponse{} }
func (m *SendImageResponse) String() string { return proto.CompactTextString(m) }
func (*SendImageResponse) ProtoMessage()    {}
func (*SendImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *SendImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendImageResponse.Unmarshal(m, b)
}
func (m *SendImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendImageResponse.Marshal(b, m, deterministic)
}
func (m *SendImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendImageResponse.Merge(m, src)
}
func (m *SendImageResponse) XXX_Size() int {
	return xxx_messageInfo_SendImageResponse.Size(m)
}
func (m *SendImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendImageResponse proto.InternalMessageInfo

func (m *SendImageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SendImageRequest)(nil), "proto.SendImageRequest")
	proto.RegisterType((*SendImageResponse)(nil), "proto.SendImageResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x1a, 0x5c,
	0x02, 0xc1, 0xa9, 0x79, 0x29, 0x9e, 0xb9, 0x89, 0xe9, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x20, 0xbe, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x84,
	0xa3, 0xa4, 0xcb, 0x25, 0x88, 0xa4, 0xb2, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b,
	0x3d, 0x37, 0xb5, 0xb8, 0x18, 0xa6, 0x98, 0x33, 0x08, 0xc6, 0x35, 0x0a, 0xe0, 0xe2, 0x01, 0x2b,
	0x0d, 0x86, 0xd8, 0x2a, 0xe4, 0xc0, 0xc5, 0x09, 0xd7, 0x2e, 0x24, 0x0e, 0x71, 0x84, 0x1e, 0xba,
	0xd5, 0x52, 0x12, 0x98, 0x12, 0x10, 0x9b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x52, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x94, 0x6a, 0x93, 0x98, 0xc9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageServiceClient interface {
	SendImage(ctx context.Context, in *SendImageRequest, opts ...grpc.CallOption) (*SendImageResponse, error)
}

type imageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageServiceClient(cc *grpc.ClientConn) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) SendImage(ctx context.Context, in *SendImageRequest, opts ...grpc.CallOption) (*SendImageResponse, error) {
	out := new(SendImageResponse)
	err := c.cc.Invoke(ctx, "/proto.ImageService/SendImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
type ImageServiceServer interface {
	SendImage(context.Context, *SendImageRequest) (*SendImageResponse, error)
}

// UnimplementedImageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (*UnimplementedImageServiceServer) SendImage(ctx context.Context, req *SendImageRequest) (*SendImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendImage not implemented")
}

func RegisterImageServiceServer(s *grpc.Server, srv ImageServiceServer) {
	s.RegisterService(&_ImageService_serviceDesc, srv)
}

func _ImageService_SendImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).SendImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImageService/SendImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).SendImage(ctx, req.(*SendImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendImage",
			Handler:    _ImageService_SendImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
