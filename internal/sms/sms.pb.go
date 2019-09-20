// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sms.proto

// SMS Service
//
// SMS Service API consists of 2 services to send a single and multiple SMSs
// This service is Idempotent:
//  It is safe to retry sending the same SMS and will be processed only once.
//  The client has to attach idempotency key with every single sms.

package sms

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SMS struct {
	IdempotencyKey       string   `protobuf:"bytes,1,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	FromPhoneNumber      string   `protobuf:"bytes,2,opt,name=from_phone_number,json=fromPhoneNumber,proto3" json:"from_phone_number,omitempty"`
	ToPhoneNumber        string   `protobuf:"bytes,3,opt,name=to_phone_number,json=toPhoneNumber,proto3" json:"to_phone_number,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMS) Reset()         { *m = SMS{} }
func (m *SMS) String() string { return proto.CompactTextString(m) }
func (*SMS) ProtoMessage()    {}
func (*SMS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{0}
}

func (m *SMS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMS.Unmarshal(m, b)
}
func (m *SMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMS.Marshal(b, m, deterministic)
}
func (m *SMS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMS.Merge(m, src)
}
func (m *SMS) XXX_Size() int {
	return xxx_messageInfo_SMS.Size(m)
}
func (m *SMS) XXX_DiscardUnknown() {
	xxx_messageInfo_SMS.DiscardUnknown(m)
}

var xxx_messageInfo_SMS proto.InternalMessageInfo

func (m *SMS) GetIdempotencyKey() string {
	if m != nil {
		return m.IdempotencyKey
	}
	return ""
}

func (m *SMS) GetFromPhoneNumber() string {
	if m != nil {
		return m.FromPhoneNumber
	}
	return ""
}

func (m *SMS) GetToPhoneNumber() string {
	if m != nil {
		return m.ToPhoneNumber
	}
	return ""
}

func (m *SMS) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SendOneRequest struct {
	Sms                  *SMS     `protobuf:"bytes,1,opt,name=sms,proto3" json:"sms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendOneRequest) Reset()         { *m = SendOneRequest{} }
func (m *SendOneRequest) String() string { return proto.CompactTextString(m) }
func (*SendOneRequest) ProtoMessage()    {}
func (*SendOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{1}
}

func (m *SendOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendOneRequest.Unmarshal(m, b)
}
func (m *SendOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendOneRequest.Marshal(b, m, deterministic)
}
func (m *SendOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendOneRequest.Merge(m, src)
}
func (m *SendOneRequest) XXX_Size() int {
	return xxx_messageInfo_SendOneRequest.Size(m)
}
func (m *SendOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendOneRequest proto.InternalMessageInfo

func (m *SendOneRequest) GetSms() *SMS {
	if m != nil {
		return m.Sms
	}
	return nil
}

type SendOneResponse struct {
	Sent                 bool     `protobuf:"varint,1,opt,name=sent,proto3" json:"sent,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendOneResponse) Reset()         { *m = SendOneResponse{} }
func (m *SendOneResponse) String() string { return proto.CompactTextString(m) }
func (*SendOneResponse) ProtoMessage()    {}
func (*SendOneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{2}
}

func (m *SendOneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendOneResponse.Unmarshal(m, b)
}
func (m *SendOneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendOneResponse.Marshal(b, m, deterministic)
}
func (m *SendOneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendOneResponse.Merge(m, src)
}
func (m *SendOneResponse) XXX_Size() int {
	return xxx_messageInfo_SendOneResponse.Size(m)
}
func (m *SendOneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendOneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendOneResponse proto.InternalMessageInfo

func (m *SendOneResponse) GetSent() bool {
	if m != nil {
		return m.Sent
	}
	return false
}

func (m *SendOneResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendManyRequest struct {
	Sms                  *SMS     `protobuf:"bytes,1,opt,name=sms,proto3" json:"sms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendManyRequest) Reset()         { *m = SendManyRequest{} }
func (m *SendManyRequest) String() string { return proto.CompactTextString(m) }
func (*SendManyRequest) ProtoMessage()    {}
func (*SendManyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{3}
}

func (m *SendManyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendManyRequest.Unmarshal(m, b)
}
func (m *SendManyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendManyRequest.Marshal(b, m, deterministic)
}
func (m *SendManyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendManyRequest.Merge(m, src)
}
func (m *SendManyRequest) XXX_Size() int {
	return xxx_messageInfo_SendManyRequest.Size(m)
}
func (m *SendManyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendManyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendManyRequest proto.InternalMessageInfo

func (m *SendManyRequest) GetSms() *SMS {
	if m != nil {
		return m.Sms
	}
	return nil
}

type SendManyResponse struct {
	// An array of errors occured (if any) for each sms
	Errors               []string `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendManyResponse) Reset()         { *m = SendManyResponse{} }
func (m *SendManyResponse) String() string { return proto.CompactTextString(m) }
func (*SendManyResponse) ProtoMessage()    {}
func (*SendManyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{4}
}

func (m *SendManyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendManyResponse.Unmarshal(m, b)
}
func (m *SendManyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendManyResponse.Marshal(b, m, deterministic)
}
func (m *SendManyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendManyResponse.Merge(m, src)
}
func (m *SendManyResponse) XXX_Size() int {
	return xxx_messageInfo_SendManyResponse.Size(m)
}
func (m *SendManyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendManyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendManyResponse proto.InternalMessageInfo

func (m *SendManyResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*SMS)(nil), "sms.SMS")
	proto.RegisterType((*SendOneRequest)(nil), "sms.SendOneRequest")
	proto.RegisterType((*SendOneResponse)(nil), "sms.SendOneResponse")
	proto.RegisterType((*SendManyRequest)(nil), "sms.SendManyRequest")
	proto.RegisterType((*SendManyResponse)(nil), "sms.SendManyResponse")
}

func init() { proto.RegisterFile("sms.proto", fileDescriptor_c8d8bdc537111860) }

var fileDescriptor_c8d8bdc537111860 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8a, 0x13, 0x41,
	0x14, 0xa6, 0x92, 0x21, 0x3f, 0x4f, 0x26, 0xd1, 0x1a, 0x95, 0xb6, 0x11, 0x1c, 0x7a, 0x15, 0x82,
	0x49, 0x41, 0x04, 0x17, 0xb3, 0x11, 0xdc, 0x4a, 0x46, 0xe9, 0x72, 0xe1, 0x2e, 0x54, 0x92, 0x67,
	0x4f, 0x33, 0x53, 0xf5, 0xda, 0xaa, 0xca, 0x0c, 0xbd, 0xf5, 0x0a, 0xde, 0xc4, 0xb5, 0xb7, 0xf0,
	0x00, 0x82, 0x78, 0x10, 0xa9, 0xee, 0xce, 0x24, 0xed, 0xca, 0x5d, 0xbd, 0xef, 0xef, 0xfd, 0x50,
	0x30, 0x74, 0xda, 0xcd, 0x0b, 0x4b, 0x9e, 0x78, 0xd7, 0x69, 0x17, 0x3f, 0xcf, 0x88, 0xb2, 0x1b,
	0x14, 0xaa, 0xc8, 0x85, 0x32, 0x86, 0xbc, 0xf2, 0x39, 0x99, 0x46, 0x12, 0xbf, 0xce, 0x72, 0x7f,
	0xb5, 0x5b, 0xcf, 0x37, 0xa4, 0x85, 0xbe, 0xcb, 0xfd, 0x35, 0xdd, 0x89, 0x8c, 0x66, 0x15, 0x39,
	0xbb, 0x55, 0x37, 0xf9, 0x56, 0x79, 0xb2, 0x4e, 0xdc, 0x3f, 0x6b, 0x5f, 0xf2, 0x83, 0x41, 0x57,
	0x2e, 0x25, 0x17, 0x30, 0xce, 0xb7, 0xa8, 0x0b, 0xf2, 0x68, 0x36, 0xe5, 0xea, 0x1a, 0xcb, 0x88,
	0x9d, 0xb3, 0xc9, 0xf0, 0x6d, 0xef, 0xf7, 0xaf, 0x17, 0x9d, 0x4f, 0x2c, 0x1d, 0x1d, 0xd1, 0xef,
	0xb0, 0xe4, 0x0b, 0x78, 0xf4, 0xd9, 0x92, 0x5e, 0x15, 0x57, 0x64, 0x70, 0x65, 0x76, 0x7a, 0x8d,
	0x36, 0xea, 0xb4, 0x2c, 0xe3, 0x20, 0xf8, 0x10, 0xf8, 0xcb, 0x8a, 0xe6, 0x73, 0x18, 0x7b, 0x6a,
	0x3b, 0xba, 0x2d, 0xc7, 0xa9, 0xa7, 0x63, 0xfd, 0x39, 0xf4, 0x37, 0x64, 0x3c, 0x1a, 0x1f, 0x9d,
	0xb4, 0x74, 0x7b, 0x38, 0x79, 0x09, 0x23, 0x89, 0x66, 0xfb, 0xde, 0x60, 0x8a, 0x5f, 0x76, 0xe8,
	0x3c, 0x8f, 0x21, 0x5c, 0xab, 0x1a, 0xfe, 0xc1, 0x62, 0x30, 0x0f, 0x47, 0x94, 0x4b, 0x99, 0x06,
	0x30, 0x79, 0x03, 0xe3, 0x7b, 0xb5, 0x2b, 0xc8, 0x38, 0xe4, 0x1c, 0x4e, 0x5c, 0xc8, 0x0f, 0xfa,
	0x41, 0x5a, 0xbd, 0x79, 0x04, 0x7d, 0x8d, 0xce, 0xa9, 0x0c, 0xeb, 0x85, 0xd2, 0x7d, 0x99, 0xcc,
	0xea, 0x80, 0xa5, 0x32, 0xe5, 0xff, 0xf4, 0x9b, 0xc2, 0xc3, 0x83, 0xbc, 0x69, 0xf8, 0x14, 0x7a,
	0x68, 0x2d, 0xd9, 0x60, 0xe9, 0x4e, 0x86, 0x69, 0x53, 0x2d, 0xbe, 0x33, 0x00, 0xb9, 0x94, 0x12,
	0xed, 0x6d, 0xbe, 0x41, 0x7e, 0x09, 0xfd, 0x66, 0x54, 0x7e, 0x56, 0x87, 0xb6, 0xd6, 0x8c, 0x1f,
	0xb7, 0xc1, 0x3a, 0x3c, 0x89, 0xbe, 0xfe, 0xfc, 0xf3, 0xad, 0xc3, 0x93, 0x53, 0xe1, 0xb4, 0x13,
	0x0e, 0xcd, 0x56, 0x90, 0xc1, 0x0b, 0x36, 0xe5, 0x1f, 0x61, 0xb0, 0x1f, 0x85, 0x1f, 0xbc, 0x47,
	0x8b, 0xc4, 0x4f, 0xfe, 0x41, 0x9b, 0xc8, 0x67, 0x55, 0xe4, 0x59, 0x32, 0x3a, 0x44, 0x6a, 0x65,
	0xca, 0x0b, 0x36, 0x9d, 0xb0, 0x75, 0xaf, 0xfa, 0x44, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x37, 0xec, 0x97, 0xf3, 0xac, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SMSServiceClient is the client API for SMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SMSServiceClient interface {
	// SendOne method sends a single sms
	SendOne(ctx context.Context, in *SendOneRequest, opts ...grpc.CallOption) (*SendOneResponse, error)
	// SendMany method sends many SMSs in one request.
	// It relies on calling SendOne method for each sms.
	//
	// For HTTP API, newline-delimited JSON is used for streaming.
	SendMany(ctx context.Context, opts ...grpc.CallOption) (SMSService_SendManyClient, error)
}

type sMSServiceClient struct {
	cc *grpc.ClientConn
}

func NewSMSServiceClient(cc *grpc.ClientConn) SMSServiceClient {
	return &sMSServiceClient{cc}
}

func (c *sMSServiceClient) SendOne(ctx context.Context, in *SendOneRequest, opts ...grpc.CallOption) (*SendOneResponse, error) {
	out := new(SendOneResponse)
	err := c.cc.Invoke(ctx, "/sms.SMSService/SendOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSServiceClient) SendMany(ctx context.Context, opts ...grpc.CallOption) (SMSService_SendManyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SMSService_serviceDesc.Streams[0], "/sms.SMSService/SendMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &sMSServiceSendManyClient{stream}
	return x, nil
}

type SMSService_SendManyClient interface {
	Send(*SendManyRequest) error
	CloseAndRecv() (*SendManyResponse, error)
	grpc.ClientStream
}

type sMSServiceSendManyClient struct {
	grpc.ClientStream
}

func (x *sMSServiceSendManyClient) Send(m *SendManyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sMSServiceSendManyClient) CloseAndRecv() (*SendManyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendManyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SMSServiceServer is the server API for SMSService service.
type SMSServiceServer interface {
	// SendOne method sends a single sms
	SendOne(context.Context, *SendOneRequest) (*SendOneResponse, error)
	// SendMany method sends many SMSs in one request.
	// It relies on calling SendOne method for each sms.
	//
	// For HTTP API, newline-delimited JSON is used for streaming.
	SendMany(SMSService_SendManyServer) error
}

// UnimplementedSMSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSMSServiceServer struct {
}

func (*UnimplementedSMSServiceServer) SendOne(ctx context.Context, req *SendOneRequest) (*SendOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOne not implemented")
}
func (*UnimplementedSMSServiceServer) SendMany(srv SMSService_SendManyServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMany not implemented")
}

func RegisterSMSServiceServer(s *grpc.Server, srv SMSServiceServer) {
	s.RegisterService(&_SMSService_serviceDesc, srv)
}

func _SMSService_SendOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).SendOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SMSService/SendOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).SendOne(ctx, req.(*SendOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMSService_SendMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SMSServiceServer).SendMany(&sMSServiceSendManyServer{stream})
}

type SMSService_SendManyServer interface {
	SendAndClose(*SendManyResponse) error
	Recv() (*SendManyRequest, error)
	grpc.ServerStream
}

type sMSServiceSendManyServer struct {
	grpc.ServerStream
}

func (x *sMSServiceSendManyServer) SendAndClose(m *SendManyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sMSServiceSendManyServer) Recv() (*SendManyRequest, error) {
	m := new(SendManyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SMSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SMSService",
	HandlerType: (*SMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOne",
			Handler:    _SMSService_SendOne_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMany",
			Handler:       _SMSService_SendMany_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "sms.proto",
}
