// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lorem.proto

package schemas

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtobufLorem struct {
	Id                   *string          `protobuf:"bytes,1,req,name=Id" json:"Id,omitempty"`
	Data                 map[int32]string `protobuf:"bytes,2,rep,name=Data" json:"Data,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Timestamp            *int64           `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProtobufLorem) Reset()         { *m = ProtobufLorem{} }
func (m *ProtobufLorem) String() string { return proto.CompactTextString(m) }
func (*ProtobufLorem) ProtoMessage()    {}
func (*ProtobufLorem) Descriptor() ([]byte, []int) {
	return fileDescriptor_lorem_6975cd4909389e3d, []int{0}
}
func (m *ProtobufLorem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtobufLorem.Unmarshal(m, b)
}
func (m *ProtobufLorem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtobufLorem.Marshal(b, m, deterministic)
}
func (dst *ProtobufLorem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtobufLorem.Merge(dst, src)
}
func (m *ProtobufLorem) XXX_Size() int {
	return xxx_messageInfo_ProtobufLorem.Size(m)
}
func (m *ProtobufLorem) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtobufLorem.DiscardUnknown(m)
}

var xxx_messageInfo_ProtobufLorem proto.InternalMessageInfo

func (m *ProtobufLorem) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ProtobufLorem) GetData() map[int32]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ProtobufLorem) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtobufLorem)(nil), "schemas.ProtobufLorem")
	proto.RegisterMapType((map[int32]string)(nil), "schemas.ProtobufLorem.DataEntry")
}

func init() { proto.RegisterFile("lorem.proto", fileDescriptor_lorem_6975cd4909389e3d) }

var fileDescriptor_lorem_6975cd4909389e3d = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x2f, 0x4a,
	0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c,
	0x56, 0x5a, 0xc3, 0xc8, 0xc5, 0x1b, 0x00, 0x12, 0x4a, 0x2a, 0x4d, 0xf3, 0x01, 0x29, 0x10, 0xe2,
	0xe3, 0x62, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x62, 0xf2, 0x4c, 0x11, 0x32,
	0xe1, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd0, 0x83,
	0xea, 0xd4, 0x43, 0xd1, 0xa5, 0x07, 0x52, 0xe2, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0x56, 0x2d,
	0x24, 0xc3, 0xc5, 0x19, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x20, 0xc1, 0xac, 0xc0,
	0xa8, 0xc1, 0x1c, 0x84, 0x10, 0x90, 0x32, 0xe7, 0xe2, 0x84, 0x6b, 0x10, 0x12, 0xe0, 0x62, 0xce,
	0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb,
	0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x2b, 0x26, 0x0b,
	0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x85, 0x5a, 0xd3, 0xc5, 0x00, 0x00, 0x00,
}
