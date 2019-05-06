// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package api

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

type Status int32

const (
	Status_OK    Status = 0
	Status_ERROR Status = 1
)

var Status_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}

var Status_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

type ImportParameter struct {
	ContentPassword      string   `protobuf:"bytes,1,opt,name=content_password,json=contentPassword,proto3" json:"content_password,omitempty"`
	ImportPsd            string   `protobuf:"bytes,2,opt,name=import_psd,json=importPsd,proto3" json:"import_psd,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportParameter) Reset()         { *m = ImportParameter{} }
func (m *ImportParameter) String() string { return proto.CompactTextString(m) }
func (*ImportParameter) ProtoMessage()    {}
func (*ImportParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *ImportParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportParameter.Unmarshal(m, b)
}
func (m *ImportParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportParameter.Marshal(b, m, deterministic)
}
func (m *ImportParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportParameter.Merge(m, src)
}
func (m *ImportParameter) XXX_Size() int {
	return xxx_messageInfo_ImportParameter.Size(m)
}
func (m *ImportParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ImportParameter proto.InternalMessageInfo

func (m *ImportParameter) GetContentPassword() string {
	if m != nil {
		return m.ContentPassword
	}
	return ""
}

func (m *ImportParameter) GetImportPsd() string {
	if m != nil {
		return m.ImportPsd
	}
	return ""
}

func (m *ImportParameter) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type AddressParameter struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressParameter) Reset()         { *m = AddressParameter{} }
func (m *AddressParameter) String() string { return proto.CompactTextString(m) }
func (*AddressParameter) ProtoMessage()    {}
func (*AddressParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AddressParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressParameter.Unmarshal(m, b)
}
func (m *AddressParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressParameter.Marshal(b, m, deterministic)
}
func (m *AddressParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressParameter.Merge(m, src)
}
func (m *AddressParameter) XXX_Size() int {
	return xxx_messageInfo_AddressParameter.Size(m)
}
func (m *AddressParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressParameter.DiscardUnknown(m)
}

var xxx_messageInfo_AddressParameter proto.InternalMessageInfo

func (m *AddressParameter) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddressParameter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AddressInfo struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressInfo) Reset()         { *m = AddressInfo{} }
func (m *AddressInfo) String() string { return proto.CompactTextString(m) }
func (*AddressInfo) ProtoMessage()    {}
func (*AddressInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *AddressInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressInfo.Unmarshal(m, b)
}
func (m *AddressInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressInfo.Marshal(b, m, deterministic)
}
func (m *AddressInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressInfo.Merge(m, src)
}
func (m *AddressInfo) XXX_Size() int {
	return xxx_messageInfo_AddressInfo.Size(m)
}
func (m *AddressInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddressInfo proto.InternalMessageInfo

func (m *AddressInfo) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *AddressInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressInfo) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type CipherParameter struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Message              []byte   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherParameter) Reset()         { *m = CipherParameter{} }
func (m *CipherParameter) String() string { return proto.CompactTextString(m) }
func (*CipherParameter) ProtoMessage()    {}
func (*CipherParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *CipherParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherParameter.Unmarshal(m, b)
}
func (m *CipherParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherParameter.Marshal(b, m, deterministic)
}
func (m *CipherParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherParameter.Merge(m, src)
}
func (m *CipherParameter) XXX_Size() int {
	return xxx_messageInfo_CipherParameter.Size(m)
}
func (m *CipherParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherParameter.DiscardUnknown(m)
}

var xxx_messageInfo_CipherParameter proto.InternalMessageInfo

func (m *CipherParameter) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CipherParameter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CipherParameter) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type CipherText struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherText) Reset()         { *m = CipherText{} }
func (m *CipherText) String() string { return proto.CompactTextString(m) }
func (*CipherText) ProtoMessage()    {}
func (*CipherText) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *CipherText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherText.Unmarshal(m, b)
}
func (m *CipherText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherText.Marshal(b, m, deterministic)
}
func (m *CipherText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherText.Merge(m, src)
}
func (m *CipherText) XXX_Size() int {
	return xxx_messageInfo_CipherText.Size(m)
}
func (m *CipherText) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherText.DiscardUnknown(m)
}

var xxx_messageInfo_CipherText proto.InternalMessageInfo

func (m *CipherText) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *CipherText) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CipherText) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.Status", Status_name, Status_value)
	proto.RegisterType((*ImportParameter)(nil), "api.ImportParameter")
	proto.RegisterType((*AddressParameter)(nil), "api.AddressParameter")
	proto.RegisterType((*AddressInfo)(nil), "api.AddressInfo")
	proto.RegisterType((*CipherParameter)(nil), "api.CipherParameter")
	proto.RegisterType((*CipherText)(nil), "api.CipherText")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0xed, 0xee, 0xc2, 0x42, 0xa6, 0xa5, 0x89, 0x2c, 0x90, 0xa2, 0x45, 0x48, 0x55, 0x10, 0xa2,
	0x70, 0x48, 0xa4, 0xc2, 0x85, 0x8f, 0x0b, 0x94, 0x0a, 0xaa, 0x1e, 0xba, 0xf2, 0x22, 0x90, 0x38,
	0xb0, 0x72, 0x93, 0x69, 0xd6, 0x42, 0x89, 0x2d, 0x7b, 0x02, 0xe4, 0x57, 0xf3, 0x17, 0x50, 0x1c,
	0x6f, 0x29, 0x01, 0xa4, 0x05, 0x6e, 0x33, 0x6f, 0xe6, 0xcd, 0x1b, 0x3f, 0xdb, 0x00, 0xa2, 0xa1,
	0x55, 0xaa, 0x8d, 0x22, 0xc5, 0x26, 0x42, 0xcb, 0xa4, 0x81, 0xf0, 0xb8, 0xd2, 0xca, 0xd0, 0x5c,
	0x18, 0x51, 0x21, 0xa1, 0x61, 0x0f, 0x20, 0xca, 0x55, 0x4d, 0x58, 0xd3, 0x52, 0x0b, 0x6b, 0xbf,
	0x28, 0x53, 0xc4, 0xa3, 0xbd, 0xd1, 0x7e, 0xc0, 0x43, 0x8f, 0xcf, 0x3d, 0xcc, 0xee, 0x00, 0x48,
	0xc7, 0x5e, 0x6a, 0x5b, 0xc4, 0x63, 0xd7, 0x14, 0xf4, 0xc8, 0xdc, 0x16, 0x2c, 0x86, 0x6b, 0x9e,
	0x11, 0x4f, 0xf6, 0x46, 0xfb, 0x3b, 0x7c, 0x9d, 0x26, 0x6f, 0x20, 0x7a, 0x51, 0x14, 0x06, 0xad,
	0xfd, 0xa1, 0x3b, 0x83, 0xeb, 0x03, 0xbd, 0x8b, 0xbc, 0x9b, 0x24, 0xfa, 0x7e, 0xaf, 0xb2, 0x4e,
	0x93, 0x8f, 0xb0, 0xed, 0x27, 0x1d, 0xd7, 0xe7, 0x8a, 0xdd, 0x85, 0xa9, 0x25, 0x41, 0x8d, 0x75,
	0x23, 0x76, 0x0f, 0xb6, 0x53, 0xa1, 0x65, 0xba, 0x70, 0x10, 0xf7, 0xa5, 0x3f, 0x4f, 0x63, 0x11,
	0x4c, 0x2a, 0x5b, 0xba, 0x6d, 0x03, 0xde, 0x85, 0x89, 0x80, 0xf0, 0x50, 0xea, 0x15, 0x9a, 0xff,
	0x5c, 0xb4, 0xab, 0x54, 0x68, 0xad, 0x28, 0x71, 0x6d, 0x86, 0x4f, 0x93, 0xf7, 0x00, 0xbd, 0xc4,
	0x5b, 0xfc, 0x4a, 0x9b, 0x9d, 0x80, 0xc1, 0x95, 0x42, 0x90, 0x70, 0x1a, 0x3b, 0xdc, 0xc5, 0xbf,
	0xee, 0xfe, 0xf0, 0x36, 0x4c, 0x7b, 0x1e, 0x9b, 0xc2, 0xf8, 0xf4, 0x24, 0xda, 0x62, 0x01, 0x5c,
	0x3d, 0xe2, 0xfc, 0x94, 0x47, 0xa3, 0x83, 0x6f, 0x63, 0x80, 0x13, 0x6c, 0x17, 0x68, 0x3e, 0xcb,
	0x1c, 0xd9, 0x73, 0x08, 0x5f, 0x63, 0x8d, 0x46, 0x10, 0x7a, 0x3f, 0xd9, 0x2d, 0xa7, 0x3c, 0xbc,
	0xa7, 0x59, 0x74, 0x19, 0xee, 0x4c, 0x4f, 0xb6, 0xd8, 0x53, 0xb8, 0xf1, 0x0e, 0x8d, 0x3c, 0x6f,
	0xff, 0x81, 0xfb, 0x04, 0x76, 0x0f, 0xfb, 0x67, 0x71, 0x54, 0xe7, 0xa6, 0xd5, 0xc4, 0x6e, 0xba,
	0xae, 0x81, 0xed, 0xb3, 0xf0, 0x12, 0xda, 0x39, 0xf5, 0x13, 0xf5, 0x15, 0xfe, 0x25, 0xf5, 0x31,
	0x04, 0x0b, 0x59, 0xd6, 0x82, 0x1a, 0x83, 0x9b, 0xb3, 0x9e, 0x41, 0xe8, 0x1f, 0xfc, 0x27, 0x6c,
	0x2d, 0xa9, 0x0b, 0xee, 0xe0, 0x13, 0xfd, 0xee, 0xa0, 0x2f, 0xef, 0x7f, 0xb8, 0x57, 0x4a, 0x5a,
	0x35, 0x67, 0x69, 0xae, 0xaa, 0xcc, 0xe6, 0xa6, 0xed, 0x0a, 0x59, 0xa1, 0x33, 0xa1, 0x65, 0x56,
	0xaa, 0xa5, 0x6a, 0xa8, 0x0b, 0xcf, 0xa6, 0xee, 0x83, 0x3e, 0xfa, 0x1e, 0x00, 0x00, 0xff, 0xff,
	0xed, 0x42, 0x10, 0x38, 0xae, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyServiceClient is the client API for KeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyServiceClient interface {
	//生成地址
	GenerateAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error)
	//校验地址
	VerifyAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error)
	//内容加密
	ContentEncrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error)
	//内容解密
	ContentDecrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error)
	//消息签名
	Signature(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error)
	//导入keystore文件内容
	ImportKeystore(ctx context.Context, in *ImportParameter, opts ...grpc.CallOption) (*AddressInfo, error)
}

type keyServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyServiceClient(cc *grpc.ClientConn) KeyServiceClient {
	return &keyServiceClient{cc}
}

func (c *keyServiceClient) GenerateAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error) {
	out := new(AddressInfo)
	err := c.cc.Invoke(ctx, "/api.KeyService/GenerateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) VerifyAddress(ctx context.Context, in *AddressParameter, opts ...grpc.CallOption) (*AddressInfo, error) {
	out := new(AddressInfo)
	err := c.cc.Invoke(ctx, "/api.KeyService/VerifyAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) ContentEncrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error) {
	out := new(CipherText)
	err := c.cc.Invoke(ctx, "/api.KeyService/ContentEncrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) ContentDecrypt(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error) {
	out := new(CipherText)
	err := c.cc.Invoke(ctx, "/api.KeyService/ContentDecrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) Signature(ctx context.Context, in *CipherParameter, opts ...grpc.CallOption) (*CipherText, error) {
	out := new(CipherText)
	err := c.cc.Invoke(ctx, "/api.KeyService/Signature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) ImportKeystore(ctx context.Context, in *ImportParameter, opts ...grpc.CallOption) (*AddressInfo, error) {
	out := new(AddressInfo)
	err := c.cc.Invoke(ctx, "/api.KeyService/import_keystore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyServiceServer is the server API for KeyService service.
type KeyServiceServer interface {
	//生成地址
	GenerateAddress(context.Context, *AddressParameter) (*AddressInfo, error)
	//校验地址
	VerifyAddress(context.Context, *AddressParameter) (*AddressInfo, error)
	//内容加密
	ContentEncrypt(context.Context, *CipherParameter) (*CipherText, error)
	//内容解密
	ContentDecrypt(context.Context, *CipherParameter) (*CipherText, error)
	//消息签名
	Signature(context.Context, *CipherParameter) (*CipherText, error)
	//导入keystore文件内容
	ImportKeystore(context.Context, *ImportParameter) (*AddressInfo, error)
}

func RegisterKeyServiceServer(s *grpc.Server, srv KeyServiceServer) {
	s.RegisterService(&_KeyService_serviceDesc, srv)
}

func _KeyService_GenerateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GenerateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyService/GenerateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GenerateAddress(ctx, req.(*AddressParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_VerifyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).VerifyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyService/VerifyAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).VerifyAddress(ctx, req.(*AddressParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_ContentEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CipherParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).ContentEncrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyService/ContentEncrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).ContentEncrypt(ctx, req.(*CipherParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_ContentDecrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CipherParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).ContentDecrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyService/ContentDecrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).ContentDecrypt(ctx, req.(*CipherParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_Signature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CipherParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).Signature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyService/Signature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).Signature(ctx, req.(*CipherParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_ImportKeystore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).ImportKeystore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyService/ImportKeystore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).ImportKeystore(ctx, req.(*ImportParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.KeyService",
	HandlerType: (*KeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAddress",
			Handler:    _KeyService_GenerateAddress_Handler,
		},
		{
			MethodName: "VerifyAddress",
			Handler:    _KeyService_VerifyAddress_Handler,
		},
		{
			MethodName: "ContentEncrypt",
			Handler:    _KeyService_ContentEncrypt_Handler,
		},
		{
			MethodName: "ContentDecrypt",
			Handler:    _KeyService_ContentDecrypt_Handler,
		},
		{
			MethodName: "Signature",
			Handler:    _KeyService_Signature_Handler,
		},
		{
			MethodName: "import_keystore",
			Handler:    _KeyService_ImportKeystore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
