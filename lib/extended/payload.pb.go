// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kchristidis/nested/lib/extended/payload.proto

package extended

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common1 "github.com/kchristidis/nested/lib/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Payload defines the values that an extended Get request can carry.
type Payload struct {
	Common *common1.Payload `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
	Extra  string           `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Payload) GetCommon() *common1.Payload {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *Payload) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func init() {
	proto.RegisterType((*Payload)(nil), "extended.Payload")
}

func init() {
	proto.RegisterFile("github.com/kchristidis/nested/lib/extended/payload.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4e, 0xce, 0x28, 0xca, 0x2c, 0x2e, 0xc9, 0x4c,
	0xc9, 0x2c, 0xd6, 0xcf, 0x4b, 0x2d, 0x2e, 0x49, 0x4d, 0xd1, 0xcf, 0xc9, 0x4c, 0xd2, 0x4f, 0xad,
	0x28, 0x49, 0xcd, 0x4b, 0x49, 0x4d, 0xd1, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x4b, 0x99, 0x11, 0x36, 0x23, 0x39, 0x3f, 0x37,
	0x37, 0x3f, 0x0f, 0xd5, 0x04, 0x25, 0x0f, 0x2e, 0xf6, 0x00, 0x88, 0x80, 0x90, 0x3a, 0x17, 0x1b,
	0x44, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xbf, 0x1e, 0x84, 0xab, 0x07, 0x55, 0x10,
	0x04, 0x95, 0x16, 0x12, 0xe1, 0x62, 0x4d, 0xad, 0x28, 0x29, 0x4a, 0x94, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x70, 0x92, 0xd8, 0xc0, 0x06, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0x4b, 0x07, 0x4f, 0xce, 0x00, 0x00, 0x00,
}
