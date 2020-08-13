// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote/ethbackend.proto

package remote

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TxRequest struct {
	Signedtx             []byte   `protobuf:"bytes,1,opt,name=signedtx,proto3" json:"signedtx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxRequest) Reset()         { *m = TxRequest{} }
func (m *TxRequest) String() string { return proto.CompactTextString(m) }
func (*TxRequest) ProtoMessage()    {}
func (*TxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4527a6acc4e6161e, []int{0}
}

func (m *TxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxRequest.Unmarshal(m, b)
}
func (m *TxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxRequest.Marshal(b, m, deterministic)
}
func (m *TxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRequest.Merge(m, src)
}
func (m *TxRequest) XXX_Size() int {
	return xxx_messageInfo_TxRequest.Size(m)
}
func (m *TxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxRequest proto.InternalMessageInfo

func (m *TxRequest) GetSignedtx() []byte {
	if m != nil {
		return m.Signedtx
	}
	return nil
}

type AddReply struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}
func (*AddReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4527a6acc4e6161e, []int{1}
}

func (m *AddReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReply.Unmarshal(m, b)
}
func (m *AddReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReply.Marshal(b, m, deterministic)
}
func (m *AddReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReply.Merge(m, src)
}
func (m *AddReply) XXX_Size() int {
	return xxx_messageInfo_AddReply.Size(m)
}
func (m *AddReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddReply proto.InternalMessageInfo

func (m *AddReply) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type EtherbaseRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EtherbaseRequest) Reset()         { *m = EtherbaseRequest{} }
func (m *EtherbaseRequest) String() string { return proto.CompactTextString(m) }
func (*EtherbaseRequest) ProtoMessage()    {}
func (*EtherbaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4527a6acc4e6161e, []int{2}
}

func (m *EtherbaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EtherbaseRequest.Unmarshal(m, b)
}
func (m *EtherbaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EtherbaseRequest.Marshal(b, m, deterministic)
}
func (m *EtherbaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EtherbaseRequest.Merge(m, src)
}
func (m *EtherbaseRequest) XXX_Size() int {
	return xxx_messageInfo_EtherbaseRequest.Size(m)
}
func (m *EtherbaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EtherbaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EtherbaseRequest proto.InternalMessageInfo

type EtherbaseReply struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EtherbaseReply) Reset()         { *m = EtherbaseReply{} }
func (m *EtherbaseReply) String() string { return proto.CompactTextString(m) }
func (*EtherbaseReply) ProtoMessage()    {}
func (*EtherbaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4527a6acc4e6161e, []int{3}
}

func (m *EtherbaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EtherbaseReply.Unmarshal(m, b)
}
func (m *EtherbaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EtherbaseReply.Marshal(b, m, deterministic)
}
func (m *EtherbaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EtherbaseReply.Merge(m, src)
}
func (m *EtherbaseReply) XXX_Size() int {
	return xxx_messageInfo_EtherbaseReply.Size(m)
}
func (m *EtherbaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EtherbaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_EtherbaseReply proto.InternalMessageInfo

func (m *EtherbaseReply) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type NetVersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetVersionRequest) Reset()         { *m = NetVersionRequest{} }
func (m *NetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*NetVersionRequest) ProtoMessage()    {}
func (*NetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4527a6acc4e6161e, []int{4}
}

func (m *NetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetVersionRequest.Unmarshal(m, b)
}
func (m *NetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetVersionRequest.Marshal(b, m, deterministic)
}
func (m *NetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetVersionRequest.Merge(m, src)
}
func (m *NetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_NetVersionRequest.Size(m)
}
func (m *NetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetVersionRequest proto.InternalMessageInfo

type NetVersionReply struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetVersionReply) Reset()         { *m = NetVersionReply{} }
func (m *NetVersionReply) String() string { return proto.CompactTextString(m) }
func (*NetVersionReply) ProtoMessage()    {}
func (*NetVersionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4527a6acc4e6161e, []int{5}
}

func (m *NetVersionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetVersionReply.Unmarshal(m, b)
}
func (m *NetVersionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetVersionReply.Marshal(b, m, deterministic)
}
func (m *NetVersionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetVersionReply.Merge(m, src)
}
func (m *NetVersionReply) XXX_Size() int {
	return xxx_messageInfo_NetVersionReply.Size(m)
}
func (m *NetVersionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NetVersionReply.DiscardUnknown(m)
}

var xxx_messageInfo_NetVersionReply proto.InternalMessageInfo

func (m *NetVersionReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*TxRequest)(nil), "remote.TxRequest")
	proto.RegisterType((*AddReply)(nil), "remote.AddReply")
	proto.RegisterType((*EtherbaseRequest)(nil), "remote.EtherbaseRequest")
	proto.RegisterType((*EtherbaseReply)(nil), "remote.EtherbaseReply")
	proto.RegisterType((*NetVersionRequest)(nil), "remote.NetVersionRequest")
	proto.RegisterType((*NetVersionReply)(nil), "remote.NetVersionReply")
}

func init() { proto.RegisterFile("remote/ethbackend.proto", fileDescriptor_4527a6acc4e6161e) }

var fileDescriptor_4527a6acc4e6161e = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xe9, 0x1c, 0x63, 0xbb, 0xc8, 0xd6, 0x45, 0x70, 0xb3, 0x0f, 0xa2, 0x45, 0x50, 0x04,
	0x33, 0xd4, 0x47, 0x11, 0x6c, 0xb5, 0x20, 0x08, 0x43, 0xca, 0xf0, 0xc1, 0xb7, 0xc6, 0x5c, 0x96,
	0xe0, 0x6c, 0x6a, 0x92, 0xc1, 0xf6, 0xd3, 0xfc, 0x77, 0xb2, 0x76, 0xad, 0x45, 0xeb, 0x53, 0x92,
	0x73, 0xee, 0x17, 0xce, 0xe1, 0xc2, 0x48, 0xe3, 0x87, 0xb2, 0x38, 0x41, 0x2b, 0x58, 0xf2, 0xf6,
	0x8e, 0x29, 0xa7, 0x99, 0x56, 0x56, 0x91, 0x4e, 0x61, 0xf8, 0xa7, 0xd0, 0x9b, 0xad, 0x62, 0xfc,
	0x5c, 0xa2, 0xb1, 0xc4, 0x83, 0xae, 0x91, 0xf3, 0x14, 0xb9, 0x5d, 0x8d, 0x9d, 0x23, 0xe7, 0x6c,
	0x37, 0xae, 0xde, 0xfe, 0x21, 0x74, 0x03, 0xce, 0x63, 0xcc, 0x16, 0x6b, 0x42, 0xa0, 0x2d, 0x12,
	0x23, 0xb6, 0x33, 0xf9, 0xdd, 0x27, 0xe0, 0x46, 0x56, 0xa0, 0x66, 0x89, 0xc1, 0xed, 0x7f, 0xfe,
	0x09, 0xf4, 0x6b, 0xda, 0x7f, 0xe4, 0x1e, 0x0c, 0xa7, 0x68, 0x5f, 0x50, 0x1b, 0xa9, 0xd2, 0x12,
	0x3d, 0x86, 0x41, 0x5d, 0xdc, 0xb0, 0x7d, 0x68, 0x49, 0x9e, 0x93, 0xed, 0xb8, 0x25, 0xf9, 0xd5,
	0x97, 0x03, 0x10, 0xcd, 0x1e, 0xc3, 0xe0, 0xfe, 0x29, 0x9a, 0x3e, 0x90, 0x73, 0xd8, 0x09, 0x38,
	0x27, 0x43, 0x5a, 0x34, 0xa3, 0x55, 0x2d, 0xcf, 0x2d, 0xa5, 0xaa, 0xc0, 0x2d, 0xf4, 0xaa, 0x60,
	0x64, 0x5c, 0xda, 0xbf, 0xf3, 0x7b, 0xfb, 0x0d, 0xce, 0x06, 0xbf, 0x03, 0xf8, 0x09, 0x47, 0x0e,
	0xca, 0xa9, 0x3f, 0x2d, 0xbc, 0x51, 0x93, 0x95, 0x2d, 0xd6, 0xe1, 0x25, 0xb8, 0x52, 0x51, 0xbb,
	0xd4, 0x4c, 0x5d, 0xcc, 0xd1, 0x0a, 0xca, 0x59, 0x58, 0x2b, 0xf3, 0xec, 0xbc, 0x0e, 0xe8, 0xa4,
	0x60, 0x6f, 0x8a, 0x83, 0x75, 0xf2, 0xc5, 0x5d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x74, 0xfd,
	0x74, 0x87, 0xd3, 0x01, 0x00, 0x00,
}
