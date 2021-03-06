// Code generated by protoc-gen-go. DO NOT EDIT.
// source: petstore.proto

package grpc2rest

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

type Pet struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}
func (*Pet) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{0}
}
func (m *Pet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pet.Unmarshal(m, b)
}
func (m *Pet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pet.Marshal(b, m, deterministic)
}
func (dst *Pet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pet.Merge(dst, src)
}
func (m *Pet) XXX_Size() int {
	return xxx_messageInfo_Pet.Size(m)
}
func (m *Pet) XXX_DiscardUnknown() {
	xxx_messageInfo_Pet.DiscardUnknown(m)
}

var xxx_messageInfo_Pet proto.InternalMessageInfo

func (m *Pet) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type PetByIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetByIdRequest) Reset()         { *m = PetByIdRequest{} }
func (m *PetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*PetByIdRequest) ProtoMessage()    {}
func (*PetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{2}
}
func (m *PetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetByIdRequest.Unmarshal(m, b)
}
func (m *PetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetByIdRequest.Marshal(b, m, deterministic)
}
func (dst *PetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetByIdRequest.Merge(dst, src)
}
func (m *PetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_PetByIdRequest.Size(m)
}
func (m *PetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PetByIdRequest proto.InternalMessageInfo

func (m *PetByIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserByNameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserByNameRequest) Reset()         { *m = UserByNameRequest{} }
func (m *UserByNameRequest) String() string { return proto.CompactTextString(m) }
func (*UserByNameRequest) ProtoMessage()    {}
func (*UserByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{3}
}
func (m *UserByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserByNameRequest.Unmarshal(m, b)
}
func (m *UserByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserByNameRequest.Marshal(b, m, deterministic)
}
func (dst *UserByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserByNameRequest.Merge(dst, src)
}
func (m *UserByNameRequest) XXX_Size() int {
	return xxx_messageInfo_UserByNameRequest.Size(m)
}
func (m *UserByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserByNameRequest proto.InternalMessageInfo

func (m *UserByNameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type PetResponse struct {
	Pet                  *Pet     `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetResponse) Reset()         { *m = PetResponse{} }
func (m *PetResponse) String() string { return proto.CompactTextString(m) }
func (*PetResponse) ProtoMessage()    {}
func (*PetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{4}
}
func (m *PetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetResponse.Unmarshal(m, b)
}
func (m *PetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetResponse.Marshal(b, m, deterministic)
}
func (dst *PetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetResponse.Merge(dst, src)
}
func (m *PetResponse) XXX_Size() int {
	return xxx_messageInfo_PetResponse.Size(m)
}
func (m *PetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PetResponse proto.InternalMessageInfo

func (m *PetResponse) GetPet() *Pet {
	if m != nil {
		return m.Pet
	}
	return nil
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{5}
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

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type PetRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetRequest) Reset()         { *m = PetRequest{} }
func (m *PetRequest) String() string { return proto.CompactTextString(m) }
func (*PetRequest) ProtoMessage()    {}
func (*PetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{6}
}
func (m *PetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetRequest.Unmarshal(m, b)
}
func (m *PetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetRequest.Marshal(b, m, deterministic)
}
func (dst *PetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetRequest.Merge(dst, src)
}
func (m *PetRequest) XXX_Size() int {
	return xxx_messageInfo_PetRequest.Size(m)
}
func (m *PetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PetRequest proto.InternalMessageInfo

func (m *PetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4fb6ef6ce229512a, []int{7}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*Pet)(nil), "grpc2rest.Pet")
	proto.RegisterType((*User)(nil), "grpc2rest.User")
	proto.RegisterType((*PetByIdRequest)(nil), "grpc2rest.PetByIdRequest")
	proto.RegisterType((*UserByNameRequest)(nil), "grpc2rest.UserByNameRequest")
	proto.RegisterType((*PetResponse)(nil), "grpc2rest.PetResponse")
	proto.RegisterType((*UserResponse)(nil), "grpc2rest.UserResponse")
	proto.RegisterType((*PetRequest)(nil), "grpc2rest.PetRequest")
	proto.RegisterType((*UserRequest)(nil), "grpc2rest.UserRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GRPC2RestPetStoreServiceClient is the client API for GRPC2RestPetStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPC2RestPetStoreServiceClient interface {
	PetById(ctx context.Context, in *PetByIdRequest, opts ...grpc.CallOption) (*PetResponse, error)
	UserByName(ctx context.Context, in *UserByNameRequest, opts ...grpc.CallOption) (*UserResponse, error)
	PetPUT(ctx context.Context, in *PetRequest, opts ...grpc.CallOption) (*PetResponse, error)
	UserPUT(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type gRPC2RestPetStoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewGRPC2RestPetStoreServiceClient(cc *grpc.ClientConn) GRPC2RestPetStoreServiceClient {
	return &gRPC2RestPetStoreServiceClient{cc}
}

func (c *gRPC2RestPetStoreServiceClient) PetById(ctx context.Context, in *PetByIdRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	out := new(PetResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.GRPC2RestPetStoreService/PetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPC2RestPetStoreServiceClient) UserByName(ctx context.Context, in *UserByNameRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.GRPC2RestPetStoreService/UserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPC2RestPetStoreServiceClient) PetPUT(ctx context.Context, in *PetRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	out := new(PetResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.GRPC2RestPetStoreService/PetPUT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPC2RestPetStoreServiceClient) UserPUT(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.GRPC2RestPetStoreService/UserPUT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPC2RestPetStoreServiceServer is the server API for GRPC2RestPetStoreService service.
type GRPC2RestPetStoreServiceServer interface {
	PetById(context.Context, *PetByIdRequest) (*PetResponse, error)
	UserByName(context.Context, *UserByNameRequest) (*UserResponse, error)
	PetPUT(context.Context, *PetRequest) (*PetResponse, error)
	UserPUT(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterGRPC2RestPetStoreServiceServer(s *grpc.Server, srv GRPC2RestPetStoreServiceServer) {
	s.RegisterService(&_GRPC2RestPetStoreService_serviceDesc, srv)
}

func _GRPC2RestPetStoreService_PetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPC2RestPetStoreServiceServer).PetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2rest.GRPC2RestPetStoreService/PetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPC2RestPetStoreServiceServer).PetById(ctx, req.(*PetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC2RestPetStoreService_UserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPC2RestPetStoreServiceServer).UserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2rest.GRPC2RestPetStoreService/UserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPC2RestPetStoreServiceServer).UserByName(ctx, req.(*UserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC2RestPetStoreService_PetPUT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPC2RestPetStoreServiceServer).PetPUT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2rest.GRPC2RestPetStoreService/PetPUT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPC2RestPetStoreServiceServer).PetPUT(ctx, req.(*PetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC2RestPetStoreService_UserPUT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPC2RestPetStoreServiceServer).UserPUT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2rest.GRPC2RestPetStoreService/UserPUT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPC2RestPetStoreServiceServer).UserPUT(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPC2RestPetStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc2rest.GRPC2RestPetStoreService",
	HandlerType: (*GRPC2RestPetStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PetById",
			Handler:    _GRPC2RestPetStoreService_PetById_Handler,
		},
		{
			MethodName: "UserByName",
			Handler:    _GRPC2RestPetStoreService_UserByName_Handler,
		},
		{
			MethodName: "PetPUT",
			Handler:    _GRPC2RestPetStoreService_PetPUT_Handler,
		},
		{
			MethodName: "UserPUT",
			Handler:    _GRPC2RestPetStoreService_UserPUT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petstore.proto",
}

func init() { proto.RegisterFile("petstore.proto", fileDescriptor_petstore_4fb6ef6ce229512a) }

var fileDescriptor_petstore_4fb6ef6ce229512a = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0xd3, 0x52, 0x40, 0x06, 0x53, 0xe3, 0x46, 0xb1, 0x12, 0x0f, 0xcd, 0x7a, 0xc1, 0x0b,
	0x98, 0x72, 0x30, 0x31, 0x9e, 0xe0, 0x60, 0xbc, 0x98, 0xcd, 0x22, 0x57, 0x13, 0x84, 0x89, 0x36,
	0x11, 0xba, 0xee, 0x2e, 0x26, 0x7c, 0x0d, 0x3f, 0xb1, 0xd9, 0x05, 0x6a, 0xff, 0x85, 0x93, 0x37,
	0x66, 0xe7, 0xbd, 0x79, 0x93, 0xdf, 0x50, 0xf0, 0x05, 0x6a, 0xa5, 0x13, 0x89, 0x7d, 0x21, 0x13,
	0x9d, 0x90, 0xd6, 0xbb, 0x14, 0xf3, 0x48, 0xa2, 0xd2, 0xf4, 0x06, 0x6a, 0x0c, 0x35, 0xf1, 0xc1,
	0x8d, 0x17, 0x81, 0x13, 0x3a, 0xbd, 0x3a, 0x77, 0xe3, 0x05, 0x21, 0xe0, 0xad, 0x66, 0x4b, 0x0c,
	0xdc, 0xd0, 0xe9, 0xb5, 0xb8, 0xfd, 0x4d, 0x5f, 0xc1, 0x9b, 0x2a, 0x94, 0x25, 0x6d, 0x17, 0x8e,
	0xd6, 0x0a, 0x65, 0x46, 0x9f, 0xd6, 0xe4, 0x0c, 0xea, 0xb8, 0x9c, 0xc5, 0x9f, 0x41, 0xcd, 0x36,
	0xb6, 0x85, 0x79, 0x15, 0x1f, 0xc9, 0x0a, 0x03, 0x6f, 0xfb, 0x6a, 0x0b, 0x1a, 0x82, 0xcf, 0x50,
	0x8f, 0x36, 0x4f, 0x0b, 0x8e, 0x5f, 0x6b, 0x54, 0xa5, 0xad, 0xe8, 0x00, 0x4e, 0xcd, 0x06, 0xa3,
	0xcd, 0xf3, 0x6c, 0x89, 0x7b, 0x51, 0x36, 0xde, 0xc9, 0xc7, 0xd3, 0x01, 0xb4, 0x19, 0x6a, 0x8e,
	0x4a, 0x24, 0x2b, 0x85, 0x24, 0x84, 0x9a, 0x40, 0x6d, 0x55, 0xed, 0xc8, 0xef, 0xa7, 0x14, 0xfa,
	0x46, 0x64, 0x5a, 0x74, 0x08, 0xc7, 0x26, 0x21, 0x75, 0x5c, 0x83, 0x67, 0x86, 0xed, 0x2c, 0x27,
	0x19, 0x8b, 0x95, 0xd9, 0x26, 0xbd, 0x05, 0xb0, 0x29, 0x95, 0x4b, 0x57, 0xa2, 0x44, 0x68, 0x6f,
	0x63, 0xaa, 0x2d, 0xff, 0x44, 0x34, 0xfa, 0x71, 0x21, 0x78, 0xe4, 0x6c, 0x1c, 0x71, 0x54, 0x9a,
	0xa1, 0x9e, 0x98, 0xff, 0xc0, 0x04, 0xe5, 0x77, 0x3c, 0x47, 0xf2, 0x00, 0xcd, 0x1d, 0x6e, 0x72,
	0x99, 0x47, 0x91, 0x39, 0x41, 0xb7, 0x53, 0xa0, 0xb4, 0x07, 0x33, 0x06, 0xf8, 0x3b, 0x05, 0xb9,
	0x2a, 0x80, 0xc9, 0x5d, 0xa8, 0x7b, 0x51, 0xc4, 0xb6, 0x1f, 0x72, 0x07, 0x0d, 0x86, 0x9a, 0x4d,
	0x5f, 0xc8, 0x79, 0x31, 0xe6, 0x70, 0xfa, 0x3d, 0x34, 0xcd, 0x20, 0xe3, 0xec, 0x94, 0x86, 0x1f,
	0x0e, 0x7d, 0x6b, 0xd8, 0x6f, 0x60, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xff, 0xce, 0x7f,
	0x15, 0x03, 0x00, 0x00,
}
