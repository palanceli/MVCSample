// Code generated by protoc-gen-go. DO NOT EDIT.
// source: receiver/receiver.proto

// option go_package = "github.com/palanceli/MVCSample/receiver;receiverpb";

package receiver_api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type ReceiveDataRequest struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveDataRequest) Reset()         { *m = ReceiveDataRequest{} }
func (m *ReceiveDataRequest) String() string { return proto.CompactTextString(m) }
func (*ReceiveDataRequest) ProtoMessage()    {}
func (*ReceiveDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c71b9929f9f99114, []int{0}
}

func (m *ReceiveDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveDataRequest.Unmarshal(m, b)
}
func (m *ReceiveDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveDataRequest.Marshal(b, m, deterministic)
}
func (m *ReceiveDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveDataRequest.Merge(m, src)
}
func (m *ReceiveDataRequest) XXX_Size() int {
	return xxx_messageInfo_ReceiveDataRequest.Size(m)
}
func (m *ReceiveDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveDataRequest proto.InternalMessageInfo

func (m *ReceiveDataRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReceiveDataRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ReceiveDataReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveDataReply) Reset()         { *m = ReceiveDataReply{} }
func (m *ReceiveDataReply) String() string { return proto.CompactTextString(m) }
func (*ReceiveDataReply) ProtoMessage()    {}
func (*ReceiveDataReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c71b9929f9f99114, []int{1}
}

func (m *ReceiveDataReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveDataReply.Unmarshal(m, b)
}
func (m *ReceiveDataReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveDataReply.Marshal(b, m, deterministic)
}
func (m *ReceiveDataReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveDataReply.Merge(m, src)
}
func (m *ReceiveDataReply) XXX_Size() int {
	return xxx_messageInfo_ReceiveDataReply.Size(m)
}
func (m *ReceiveDataReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveDataReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveDataReply proto.InternalMessageInfo

func (m *ReceiveDataReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReceiveDataReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReceiveDataRequest)(nil), "receiver.api.ReceiveDataRequest")
	proto.RegisterType((*ReceiveDataReply)(nil), "receiver.api.ReceiveDataReply")
}

func init() { proto.RegisterFile("receiver/receiver.proto", fileDescriptor_c71b9929f9f99114) }

var fileDescriptor_c71b9929f9f99114 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4a, 0x4d, 0x4e,
	0xcd, 0x2c, 0x4b, 0x2d, 0xd2, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0xe0,
	0xfc, 0xc4, 0x82, 0x4c, 0x25, 0x27, 0x2e, 0xa1, 0x20, 0x08, 0xdf, 0x25, 0xb1, 0x24, 0x31, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x35, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcd,
	0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0x6c, 0xb8, 0x04, 0x50, 0xcc,
	0x28, 0xc8, 0xa9, 0x14, 0x12, 0xe3, 0x62, 0x2b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x86, 0x9a, 0x01,
	0xe5, 0x09, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x43, 0x4d, 0x00, 0x31, 0x8d, 0x52, 0xb8, 0xf8,
	0xa1, 0xba, 0x8b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x02, 0xb9, 0xb8, 0x91, 0x0c,
	0x14, 0x52, 0xd0, 0x43, 0x76, 0xb2, 0x1e, 0xa6, 0x7b, 0xa5, 0xe4, 0xf0, 0xa8, 0x28, 0xc8, 0xa9,
	0x54, 0x62, 0x48, 0x62, 0x03, 0x7b, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xca, 0x77, 0xa7,
	0x21, 0x17, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReceiverServiceClient is the client API for ReceiverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReceiverServiceClient interface {
	ReceiveData(ctx context.Context, in *ReceiveDataRequest, opts ...grpc.CallOption) (*ReceiveDataReply, error)
}

type receiverServiceClient struct {
	cc *grpc.ClientConn
}

func NewReceiverServiceClient(cc *grpc.ClientConn) ReceiverServiceClient {
	return &receiverServiceClient{cc}
}

func (c *receiverServiceClient) ReceiveData(ctx context.Context, in *ReceiveDataRequest, opts ...grpc.CallOption) (*ReceiveDataReply, error) {
	out := new(ReceiveDataReply)
	err := c.cc.Invoke(ctx, "/receiver.api.ReceiverService/ReceiveData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiverServiceServer is the server API for ReceiverService service.
type ReceiverServiceServer interface {
	ReceiveData(context.Context, *ReceiveDataRequest) (*ReceiveDataReply, error)
}

func RegisterReceiverServiceServer(s *grpc.Server, srv ReceiverServiceServer) {
	s.RegisterService(&_ReceiverService_serviceDesc, srv)
}

func _ReceiverService_ReceiveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServiceServer).ReceiveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/receiver.api.ReceiverService/ReceiveData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServiceServer).ReceiveData(ctx, req.(*ReceiveDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReceiverService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "receiver.api.ReceiverService",
	HandlerType: (*ReceiverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveData",
			Handler:    _ReceiverService_ReceiveData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiver/receiver.proto",
}
