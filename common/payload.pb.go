// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/payload.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Payload struct {
	StartFrom int32 `protobuf:"varint,1,opt,name=start_from,json=startFrom" json:"start_from,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Payload) GetStartFrom() int32 {
	if m != nil {
		return m.StartFrom
	}
	return 0
}

func init() {
	proto.RegisterType((*Payload)(nil), "common.Payload")
}

func init() { proto.RegisterFile("common/payload.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x83, 0x88, 0x2a, 0x69, 0x70, 0xb1, 0x07, 0x40, 0x24, 0x84, 0x64, 0xb9, 0xb8, 0x8a,
	0x4b, 0x12, 0x8b, 0x4a, 0xe2, 0xd3, 0x8a, 0xf2, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0x38, 0xc1, 0x22, 0x6e, 0x45, 0xf9, 0xb9, 0x4e, 0x6a, 0x51, 0x2a, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xd9, 0xc9, 0x19, 0x45, 0x99, 0xc5, 0x25, 0x99, 0x29, 0x99,
	0xc5, 0xfa, 0x79, 0xa9, 0xc5, 0x25, 0xa9, 0x29, 0xfa, 0x10, 0x13, 0x93, 0xd8, 0xc0, 0x16, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x56, 0x4e, 0x27, 0x78, 0x00, 0x00, 0x00,
}
