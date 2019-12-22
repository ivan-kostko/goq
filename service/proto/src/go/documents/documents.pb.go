// Code generated by protoc-gen-go. DO NOT EDIT.
// source: documents.proto

package service

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d3647adf041676, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Document struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d3647adf041676, []int{1}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Document) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type DocumentStoreReuest struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DocumentStoreReuest) Reset()         { *m = DocumentStoreReuest{} }
func (m *DocumentStoreReuest) String() string { return proto.CompactTextString(m) }
func (*DocumentStoreReuest) ProtoMessage()    {}
func (*DocumentStoreReuest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d3647adf041676, []int{2}
}

func (m *DocumentStoreReuest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentStoreReuest.Unmarshal(m, b)
}
func (m *DocumentStoreReuest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentStoreReuest.Marshal(b, m, deterministic)
}
func (m *DocumentStoreReuest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentStoreReuest.Merge(m, src)
}
func (m *DocumentStoreReuest) XXX_Size() int {
	return xxx_messageInfo_DocumentStoreReuest.Size(m)
}
func (m *DocumentStoreReuest) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentStoreReuest.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentStoreReuest proto.InternalMessageInfo

func (m *DocumentStoreReuest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type DocumentGetReuest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentGetReuest) Reset()         { *m = DocumentGetReuest{} }
func (m *DocumentGetReuest) String() string { return proto.CompactTextString(m) }
func (*DocumentGetReuest) ProtoMessage()    {}
func (*DocumentGetReuest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d3647adf041676, []int{3}
}

func (m *DocumentGetReuest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentGetReuest.Unmarshal(m, b)
}
func (m *DocumentGetReuest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentGetReuest.Marshal(b, m, deterministic)
}
func (m *DocumentGetReuest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentGetReuest.Merge(m, src)
}
func (m *DocumentGetReuest) XXX_Size() int {
	return xxx_messageInfo_DocumentGetReuest.Size(m)
}
func (m *DocumentGetReuest) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentGetReuest.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentGetReuest proto.InternalMessageInfo

func (m *DocumentGetReuest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type DocumentGetResponse struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DocumentGetResponse) Reset()         { *m = DocumentGetResponse{} }
func (m *DocumentGetResponse) String() string { return proto.CompactTextString(m) }
func (*DocumentGetResponse) ProtoMessage()    {}
func (*DocumentGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d3647adf041676, []int{4}
}

func (m *DocumentGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentGetResponse.Unmarshal(m, b)
}
func (m *DocumentGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentGetResponse.Marshal(b, m, deterministic)
}
func (m *DocumentGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentGetResponse.Merge(m, src)
}
func (m *DocumentGetResponse) XXX_Size() int {
	return xxx_messageInfo_DocumentGetResponse.Size(m)
}
func (m *DocumentGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentGetResponse proto.InternalMessageInfo

func (m *DocumentGetResponse) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "service.Empty")
	proto.RegisterType((*Document)(nil), "service.Document")
	proto.RegisterType((*DocumentStoreReuest)(nil), "service.DocumentStoreReuest")
	proto.RegisterType((*DocumentGetReuest)(nil), "service.DocumentGetReuest")
	proto.RegisterType((*DocumentGetResponse)(nil), "service.DocumentGetResponse")
}

func init() { proto.RegisterFile("documents.proto", fileDescriptor_37d3647adf041676) }

var fileDescriptor_37d3647adf041676 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x3d, 0x6b, 0x84, 0x40,
	0x10, 0xc5, 0x04, 0xa3, 0x4e, 0x20, 0xc1, 0x49, 0x23, 0x22, 0x21, 0x6c, 0x95, 0x26, 0x16, 0x4a,
	0xea, 0x34, 0x06, 0x7b, 0xf3, 0x0f, 0xa2, 0x13, 0xd8, 0xc2, 0x5d, 0xd9, 0x1d, 0x03, 0x57, 0xdf,
	0x1f, 0x3f, 0xd8, 0x73, 0x45, 0xf0, 0xe0, 0xb8, 0xd2, 0xf7, 0x35, 0xef, 0xb9, 0xf0, 0x3c, 0xe8,
	0x7e, 0x1e, 0x49, 0xb1, 0x2d, 0x27, 0xa3, 0x59, 0x63, 0x64, 0xc9, 0xfc, 0xcb, 0x9e, 0x44, 0x04,
	0xe1, 0xf7, 0x38, 0xf1, 0x41, 0x34, 0x10, 0x37, 0x8b, 0x08, 0x5f, 0x01, 0xe4, 0x40, 0x8a, 0xe5,
	0x9f, 0x24, 0x93, 0x05, 0x6f, 0xc1, 0x7b, 0xd2, 0x6d, 0x10, 0xcc, 0x20, 0xea, 0xb5, 0x62, 0x52,
	0x9c, 0xdd, 0x39, 0xd2, 0x7f, 0x8a, 0x06, 0x5e, 0x7c, 0xca, 0x0f, 0x6b, 0x43, 0x1d, 0xcd, 0x64,
	0x19, 0x3f, 0x20, 0xf6, 0x0d, 0x5c, 0xdc, 0x63, 0x95, 0x96, 0x4b, 0x83, 0xd2, 0xeb, 0xbb, 0x55,
	0x22, 0x6a, 0x48, 0x3d, 0xda, 0x12, 0x2f, 0x19, 0x57, 0x4a, 0x6d, 0x4f, 0x3b, 0x93, 0x9d, 0xb4,
	0xb2, 0x74, 0xe3, 0xe9, 0xea, 0x18, 0x40, 0xe2, 0x61, 0x8b, 0x9f, 0x10, 0xba, 0x19, 0x58, 0xec,
	0x3c, 0x9b, 0x79, 0xf9, 0xd3, 0xca, 0xba, 0x7f, 0x89, 0x5f, 0x70, 0xdf, 0x12, 0x63, 0xbe, 0x33,
	0xad, 0x6b, 0xf2, 0xe2, 0x32, 0x77, 0x2e, 0xfd, 0xfb, 0xe0, 0x5e, 0xa9, 0x3e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xc6, 0xef, 0x7e, 0x19, 0xb8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DocumentsClient is the client API for Documents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DocumentsClient interface {
	Store(ctx context.Context, in *DocumentStoreReuest, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *DocumentGetReuest, opts ...grpc.CallOption) (*DocumentGetResponse, error)
}

type documentsClient struct {
	cc *grpc.ClientConn
}

func NewDocumentsClient(cc *grpc.ClientConn) DocumentsClient {
	return &documentsClient{cc}
}

func (c *documentsClient) Store(ctx context.Context, in *DocumentStoreReuest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.Documents/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsClient) Get(ctx context.Context, in *DocumentGetReuest, opts ...grpc.CallOption) (*DocumentGetResponse, error) {
	out := new(DocumentGetResponse)
	err := c.cc.Invoke(ctx, "/service.Documents/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentsServer is the server API for Documents service.
type DocumentsServer interface {
	Store(context.Context, *DocumentStoreReuest) (*Empty, error)
	Get(context.Context, *DocumentGetReuest) (*DocumentGetResponse, error)
}

// UnimplementedDocumentsServer can be embedded to have forward compatible implementations.
type UnimplementedDocumentsServer struct {
}

func (*UnimplementedDocumentsServer) Store(ctx context.Context, req *DocumentStoreReuest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedDocumentsServer) Get(ctx context.Context, req *DocumentGetReuest) (*DocumentGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterDocumentsServer(s *grpc.Server, srv DocumentsServer) {
	s.RegisterService(&_Documents_serviceDesc, srv)
}

func _Documents_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentStoreReuest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Documents/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).Store(ctx, req.(*DocumentStoreReuest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Documents_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentGetReuest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Documents/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).Get(ctx, req.(*DocumentGetReuest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Documents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Documents",
	HandlerType: (*DocumentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Documents_Store_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Documents_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "documents.proto",
}
