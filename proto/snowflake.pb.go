// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snowflake.proto

package delz_snowflake

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 生成Snowflake输入参数
type SnowflakeRequest struct {
	// 业务Id
	ServiceId int64 `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	// 返回数据格式
	Format               string   `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnowflakeRequest) Reset()         { *m = SnowflakeRequest{} }
func (m *SnowflakeRequest) String() string { return proto.CompactTextString(m) }
func (*SnowflakeRequest) ProtoMessage()    {}
func (*SnowflakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_snowflake_88301cb484b58db7, []int{0}
}
func (m *SnowflakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnowflakeRequest.Unmarshal(m, b)
}
func (m *SnowflakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnowflakeRequest.Marshal(b, m, deterministic)
}
func (dst *SnowflakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnowflakeRequest.Merge(dst, src)
}
func (m *SnowflakeRequest) XXX_Size() int {
	return xxx_messageInfo_SnowflakeRequest.Size(m)
}
func (m *SnowflakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnowflakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SnowflakeRequest proto.InternalMessageInfo

func (m *SnowflakeRequest) GetServiceId() int64 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *SnowflakeRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

// 生成Snowflake返回参数
type SnowflakeReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnowflakeReply) Reset()         { *m = SnowflakeReply{} }
func (m *SnowflakeReply) String() string { return proto.CompactTextString(m) }
func (*SnowflakeReply) ProtoMessage()    {}
func (*SnowflakeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_snowflake_88301cb484b58db7, []int{1}
}
func (m *SnowflakeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnowflakeReply.Unmarshal(m, b)
}
func (m *SnowflakeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnowflakeReply.Marshal(b, m, deterministic)
}
func (dst *SnowflakeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnowflakeReply.Merge(dst, src)
}
func (m *SnowflakeReply) XXX_Size() int {
	return xxx_messageInfo_SnowflakeReply.Size(m)
}
func (m *SnowflakeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SnowflakeReply.DiscardUnknown(m)
}

var xxx_messageInfo_SnowflakeReply proto.InternalMessageInfo

func (m *SnowflakeReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*SnowflakeRequest)(nil), "delz.snowflake.SnowflakeRequest")
	proto.RegisterType((*SnowflakeReply)(nil), "delz.snowflake.SnowflakeReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SnowflakeClient is the client API for Snowflake service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SnowflakeClient interface {
	// 生成id
	Gen(ctx context.Context, in *SnowflakeRequest, opts ...grpc.CallOption) (*SnowflakeReply, error)
}

type snowflakeClient struct {
	cc *grpc.ClientConn
}

func NewSnowflakeClient(cc *grpc.ClientConn) SnowflakeClient {
	return &snowflakeClient{cc}
}

func (c *snowflakeClient) Gen(ctx context.Context, in *SnowflakeRequest, opts ...grpc.CallOption) (*SnowflakeReply, error) {
	out := new(SnowflakeReply)
	err := c.cc.Invoke(ctx, "/delz.snowflake.Snowflake/Gen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnowflakeServer is the server API for Snowflake service.
type SnowflakeServer interface {
	// 生成id
	Gen(context.Context, *SnowflakeRequest) (*SnowflakeReply, error)
}

func RegisterSnowflakeServer(s *grpc.Server, srv SnowflakeServer) {
	s.RegisterService(&_Snowflake_serviceDesc, srv)
}

func _Snowflake_Gen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnowflakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowflakeServer).Gen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delz.snowflake.Snowflake/Gen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowflakeServer).Gen(ctx, req.(*SnowflakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Snowflake_serviceDesc = grpc.ServiceDesc{
	ServiceName: "delz.snowflake.Snowflake",
	HandlerType: (*SnowflakeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Gen",
			Handler:    _Snowflake_Gen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snowflake.proto",
}

func init() { proto.RegisterFile("snowflake.proto", fileDescriptor_snowflake_88301cb484b58db7) }

var fileDescriptor_snowflake_88301cb484b58db7 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0xcb, 0x2f,
	0x4f, 0xcb, 0x49, 0xcc, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0x49, 0xcd,
	0xa9, 0xd2, 0x83, 0x8b, 0x2a, 0x79, 0x70, 0x09, 0x04, 0xc3, 0x38, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x32, 0x5c, 0x9c, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x9e, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x08, 0x01, 0x21, 0x31, 0x2e, 0xb6, 0xb4, 0xfc, 0xa2, 0xdc,
	0xc4, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x49, 0x81, 0x8b, 0x0f, 0xc9,
	0xa4, 0x82, 0x9c, 0x4a, 0x21, 0x3e, 0x2e, 0xa6, 0x4c, 0x88, 0x01, 0x9c, 0x41, 0x4c, 0x99, 0x29,
	0x46, 0x61, 0x5c, 0x9c, 0x70, 0x15, 0x42, 0x9e, 0x5c, 0xcc, 0xee, 0xa9, 0x79, 0x42, 0x0a, 0x7a,
	0xa8, 0x0e, 0xd2, 0x43, 0x77, 0x8d, 0x94, 0x1c, 0x1e, 0x15, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49,
	0x6c, 0x60, 0xaf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xd4, 0x66, 0xbe, 0xed, 0x00,
	0x00, 0x00,
}