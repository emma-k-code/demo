// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

/*
Golang package
*/

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

// UserGetRequest 取會員資料參數
type UserGetRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetRequest) Reset()         { *m = UserGetRequest{} }
func (m *UserGetRequest) String() string { return proto.CompactTextString(m) }
func (*UserGetRequest) ProtoMessage()    {}
func (*UserGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45dc9fa912b34814, []int{0}
}
func (m *UserGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetRequest.Unmarshal(m, b)
}
func (m *UserGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetRequest.Marshal(b, m, deterministic)
}
func (dst *UserGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetRequest.Merge(dst, src)
}
func (m *UserGetRequest) XXX_Size() int {
	return xxx_messageInfo_UserGetRequest.Size(m)
}
func (m *UserGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetRequest proto.InternalMessageInfo

func (m *UserGetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// UserResponse 會員資料
type UserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45dc9fa912b34814, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*UserGetRequest)(nil), "pb.UserGetRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// Get 取會員資料
	Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// Get 取會員資料
	Get(context.Context, *UserGetRequest) (*UserResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*UserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_45dc9fa912b34814) }

var fileDescriptor_user_45dc9fa912b34814 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe0, 0xe2, 0x0b, 0x2d,
	0x4e, 0x2d, 0x72, 0x4f, 0x2d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62,
	0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x51, 0xb2, 0xe2, 0xe2,
	0x01, 0xa9, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x45, 0x97, 0x17, 0x92, 0xe2, 0xe2,
	0x00, 0x99, 0x99, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x1b,
	0xd9, 0x70, 0x71, 0x83, 0xf4, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xe9, 0x72, 0x31,
	0xbb, 0xa7, 0x96, 0x08, 0x09, 0xe9, 0x15, 0x24, 0xe9, 0xa1, 0xda, 0x2a, 0x25, 0x00, 0x13, 0x83,
	0xd9, 0xa3, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x49, 0xd5,
	0x80, 0xb3, 0xb4, 0x00, 0x00, 0x00,
}