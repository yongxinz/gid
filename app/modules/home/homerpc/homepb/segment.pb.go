// Code generated by protoc-gen-go. DO NOT EDIT.
// source: segment.proto

package app_homepb

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

type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_0426eb73ea5ac81d, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0426eb73ea5ac81d, []int{1}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingReply struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0426eb73ea5ac81d, []int{2}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PingReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type IdRequest struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0426eb73ea5ac81d, []int{3}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type IdReply struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReply) Reset()         { *m = IdReply{} }
func (m *IdReply) String() string { return proto.CompactTextString(m) }
func (*IdReply) ProtoMessage()    {}
func (*IdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0426eb73ea5ac81d, []int{4}
}

func (m *IdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdReply.Unmarshal(m, b)
}
func (m *IdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdReply.Marshal(b, m, deterministic)
}
func (m *IdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReply.Merge(m, src)
}
func (m *IdReply) XXX_Size() int {
	return xxx_messageInfo_IdReply.Size(m)
}
func (m *IdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReply.DiscardUnknown(m)
}

var xxx_messageInfo_IdReply proto.InternalMessageInfo

func (m *IdReply) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IdReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SnowIdRequest struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnowIdRequest) Reset()         { *m = SnowIdRequest{} }
func (m *SnowIdRequest) String() string { return proto.CompactTextString(m) }
func (*SnowIdRequest) ProtoMessage()    {}
func (*SnowIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0426eb73ea5ac81d, []int{5}
}

func (m *SnowIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnowIdRequest.Unmarshal(m, b)
}
func (m *SnowIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnowIdRequest.Marshal(b, m, deterministic)
}
func (m *SnowIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnowIdRequest.Merge(m, src)
}
func (m *SnowIdRequest) XXX_Size() int {
	return xxx_messageInfo_SnowIdRequest.Size(m)
}
func (m *SnowIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnowIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SnowIdRequest proto.InternalMessageInfo

func (m *SnowIdRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type SnowIdReply struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnowIdReply) Reset()         { *m = SnowIdReply{} }
func (m *SnowIdReply) String() string { return proto.CompactTextString(m) }
func (*SnowIdReply) ProtoMessage()    {}
func (*SnowIdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0426eb73ea5ac81d, []int{6}
}

func (m *SnowIdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnowIdReply.Unmarshal(m, b)
}
func (m *SnowIdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnowIdReply.Marshal(b, m, deterministic)
}
func (m *SnowIdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnowIdReply.Merge(m, src)
}
func (m *SnowIdReply) XXX_Size() int {
	return xxx_messageInfo_SnowIdReply.Size(m)
}
func (m *SnowIdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SnowIdReply.DiscardUnknown(m)
}

var xxx_messageInfo_SnowIdReply proto.InternalMessageInfo

func (m *SnowIdReply) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SnowIdReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "app.homepb.Status")
	proto.RegisterType((*PingRequest)(nil), "app.homepb.PingRequest")
	proto.RegisterType((*PingReply)(nil), "app.homepb.PingReply")
	proto.RegisterType((*IdRequest)(nil), "app.homepb.IdRequest")
	proto.RegisterType((*IdReply)(nil), "app.homepb.IdReply")
	proto.RegisterType((*SnowIdRequest)(nil), "app.homepb.SnowIdRequest")
	proto.RegisterType((*SnowIdReply)(nil), "app.homepb.SnowIdReply")
}

func init() { proto.RegisterFile("segment.proto", fileDescriptor_0426eb73ea5ac81d) }

var fileDescriptor_0426eb73ea5ac81d = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xba, 0x55, 0xf2, 0x4a, 0x45, 0x9e, 0xc8, 0x66, 0x41, 0x98, 0x39, 0x0d, 0x0f,
	0x39, 0x6c, 0xe0, 0xd5, 0x93, 0x94, 0xe2, 0x45, 0xb2, 0x4f, 0x90, 0x99, 0x50, 0x0b, 0x6b, 0x13,
	0x4d, 0x86, 0xec, 0x73, 0xf9, 0x05, 0x25, 0x69, 0xd5, 0x4e, 0x45, 0x10, 0x6f, 0xff, 0xf6, 0xff,
	0xde, 0xff, 0xf7, 0xde, 0x23, 0x90, 0x59, 0x55, 0x35, 0xaa, 0x75, 0xcc, 0x3c, 0x6b, 0xa7, 0x11,
	0x84, 0x31, 0xec, 0x51, 0x37, 0xca, 0x6c, 0x28, 0x83, 0x64, 0xed, 0x84, 0xdb, 0x59, 0x44, 0x18,
	0x3f, 0x68, 0xa9, 0x66, 0xd1, 0x3c, 0x5a, 0x4c, 0x78, 0xd0, 0x78, 0x02, 0x71, 0x63, 0xab, 0xd9,
	0x68, 0x1e, 0x2d, 0x08, 0xf7, 0x92, 0x66, 0x90, 0xde, 0xd7, 0x6d, 0xc5, 0xd5, 0xd3, 0x4e, 0x59,
	0x47, 0xef, 0x80, 0x74, 0x9f, 0x66, 0xbb, 0xc7, 0x2b, 0x48, 0x6c, 0xc8, 0x0a, 0x19, 0xe9, 0x12,
	0xd9, 0x27, 0x88, 0x75, 0x14, 0xde, 0x57, 0x78, 0x9a, 0x14, 0x4e, 0xf4, 0xd1, 0x41, 0xd3, 0x0b,
	0x20, 0xa5, 0xec, 0x93, 0x3d, 0xda, 0x89, 0x2a, 0x24, 0x11, 0xee, 0x25, 0xbd, 0x85, 0x23, 0x6f,
	0xff, 0x95, 0x74, 0x0c, 0xa3, 0x5a, 0x06, 0x4e, 0xcc, 0x47, 0xb5, 0xa4, 0x97, 0x90, 0xad, 0x5b,
	0xfd, 0xf2, 0x1b, 0xa9, 0x84, 0xf4, 0xbd, 0xe4, 0x9f, 0xb4, 0xe5, 0x6b, 0x04, 0x71, 0x51, 0x4b,
	0xbc, 0x86, 0xb1, 0x3f, 0x14, 0x4e, 0x87, 0xbd, 0x83, 0x4b, 0xe6, 0x67, 0xdf, 0x0d, 0xcf, 0x5e,
	0xc1, 0xa4, 0x50, 0xae, 0x94, 0x78, 0xe0, 0x7f, 0x0c, 0x9f, 0x9f, 0x7e, 0xfd, 0xed, 0x9b, 0x6e,
	0x80, 0x14, 0xca, 0x75, 0x2b, 0xe0, 0xf9, 0xc1, 0xb4, 0xc3, 0xcd, 0xf3, 0xe9, 0x4f, 0x96, 0xd9,
	0xee, 0x37, 0x49, 0x78, 0x28, 0xab, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xf8, 0x36, 0xb5,
	0x39, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GidClient is the client API for Gid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GidClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdReply, error)
	GetSnowId(ctx context.Context, in *SnowIdRequest, opts ...grpc.CallOption) (*SnowIdReply, error)
}

type gidClient struct {
	cc *grpc.ClientConn
}

func NewGidClient(cc *grpc.ClientConn) GidClient {
	return &gidClient{cc}
}

func (c *gidClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/app.homepb.Gid/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gidClient) GetId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdReply, error) {
	out := new(IdReply)
	err := c.cc.Invoke(ctx, "/app.homepb.Gid/GetId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gidClient) GetSnowId(ctx context.Context, in *SnowIdRequest, opts ...grpc.CallOption) (*SnowIdReply, error) {
	out := new(SnowIdReply)
	err := c.cc.Invoke(ctx, "/app.homepb.Gid/GetSnowId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GidServer is the server API for Gid service.
type GidServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetId(context.Context, *IdRequest) (*IdReply, error)
	GetSnowId(context.Context, *SnowIdRequest) (*SnowIdReply, error)
}

// UnimplementedGidServer can be embedded to have forward compatible implementations.
type UnimplementedGidServer struct {
}

func (*UnimplementedGidServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedGidServer) GetId(ctx context.Context, req *IdRequest) (*IdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetId not implemented")
}
func (*UnimplementedGidServer) GetSnowId(ctx context.Context, req *SnowIdRequest) (*SnowIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnowId not implemented")
}

func RegisterGidServer(s *grpc.Server, srv GidServer) {
	s.RegisterService(&_Gid_serviceDesc, srv)
}

func _Gid_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GidServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.homepb.Gid/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GidServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gid_GetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GidServer).GetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.homepb.Gid/GetId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GidServer).GetId(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gid_GetSnowId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnowIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GidServer).GetSnowId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.homepb.Gid/GetSnowId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GidServer).GetSnowId(ctx, req.(*SnowIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gid_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.homepb.Gid",
	HandlerType: (*GidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Gid_Ping_Handler,
		},
		{
			MethodName: "GetId",
			Handler:    _Gid_GetId_Handler,
		},
		{
			MethodName: "GetSnowId",
			Handler:    _Gid_GetSnowId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "segment.proto",
}
