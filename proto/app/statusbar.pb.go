// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/proto/app/statusbar.proto

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	gomatcha.io/matcha/proto/app/statusbar.proto

It has these top-level messages:
	ActivityIndicator
	StatusBar
*/
package app

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

type StatusBarStyle int32

const (
	StatusBarStyle_STATUS_BAR_STYLE_DEFAULT StatusBarStyle = 0
	StatusBarStyle_STATUS_BAR_STYLE_LIGHT   StatusBarStyle = 1
	StatusBarStyle_STATUS_BAR_STYLE_DARK    StatusBarStyle = 2
)

var StatusBarStyle_name = map[int32]string{
	0: "STATUS_BAR_STYLE_DEFAULT",
	1: "STATUS_BAR_STYLE_LIGHT",
	2: "STATUS_BAR_STYLE_DARK",
}
var StatusBarStyle_value = map[string]int32{
	"STATUS_BAR_STYLE_DEFAULT": 0,
	"STATUS_BAR_STYLE_LIGHT":   1,
	"STATUS_BAR_STYLE_DARK":    2,
}

func (x StatusBarStyle) String() string {
	return proto.EnumName(StatusBarStyle_name, int32(x))
}
func (StatusBarStyle) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ActivityIndicator struct {
	Visible bool `protobuf:"varint,1,opt,name=visible" json:"visible,omitempty"`
}

func (m *ActivityIndicator) Reset()                    { *m = ActivityIndicator{} }
func (m *ActivityIndicator) String() string            { return proto.CompactTextString(m) }
func (*ActivityIndicator) ProtoMessage()               {}
func (*ActivityIndicator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActivityIndicator) GetVisible() bool {
	if m != nil {
		return m.Visible
	}
	return false
}

type StatusBar struct {
	Hidden bool           `protobuf:"varint,1,opt,name=hidden" json:"hidden,omitempty"`
	Style  StatusBarStyle `protobuf:"varint,2,opt,name=style,enum=app.StatusBarStyle" json:"style,omitempty"`
}

func (m *StatusBar) Reset()                    { *m = StatusBar{} }
func (m *StatusBar) String() string            { return proto.CompactTextString(m) }
func (*StatusBar) ProtoMessage()               {}
func (*StatusBar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusBar) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *StatusBar) GetStyle() StatusBarStyle {
	if m != nil {
		return m.Style
	}
	return StatusBarStyle_STATUS_BAR_STYLE_DEFAULT
}

func init() {
	proto.RegisterType((*ActivityIndicator)(nil), "app.ActivityIndicator")
	proto.RegisterType((*StatusBar)(nil), "app.StatusBar")
	proto.RegisterEnum("app.StatusBarStyle", StatusBarStyle_name, StatusBarStyle_value)
}

func init() { proto.RegisterFile("gomatcha.io/matcha/proto/app/statusbar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x0b, 0x8a, 0xf2, 0x4b, 0xf2,
	0xf5, 0x13, 0x0b, 0x0a, 0xf4, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x93, 0x12, 0x8b, 0xf4, 0xc0,
	0x62, 0x42, 0xcc, 0x89, 0x05, 0x05, 0x4a, 0xba, 0x5c, 0x82, 0x8e, 0xc9, 0x25, 0x99, 0x65, 0x99,
	0x25, 0x95, 0x9e, 0x79, 0x29, 0x99, 0xc9, 0x89, 0x25, 0xf9, 0x45, 0x42, 0x12, 0x5c, 0xec, 0x65,
	0x99, 0xc5, 0x99, 0x49, 0x39, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x30, 0xae, 0x92,
	0x1f, 0x17, 0x67, 0x30, 0xd8, 0x18, 0xa7, 0xc4, 0x22, 0x21, 0x31, 0x2e, 0xb6, 0x8c, 0xcc, 0x94,
	0x94, 0xd4, 0x3c, 0xa8, 0x2a, 0x28, 0x4f, 0x48, 0x93, 0x8b, 0xb5, 0xb8, 0xa4, 0x32, 0x27, 0x55,
	0x82, 0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x58, 0x2f, 0xb1, 0xa0, 0x40, 0x0f, 0xae, 0x2d, 0x18,
	0x24, 0x15, 0x04, 0x51, 0xa1, 0x95, 0xca, 0xc5, 0x87, 0x2a, 0x21, 0x24, 0xc3, 0x25, 0x11, 0x1c,
	0xe2, 0x18, 0x12, 0x1a, 0x1c, 0xef, 0xe4, 0x18, 0x14, 0x1f, 0x1c, 0x12, 0xe9, 0xe3, 0x1a, 0xef,
	0xe2, 0xea, 0xe6, 0x18, 0xea, 0x13, 0x22, 0xc0, 0x20, 0x24, 0xc5, 0x25, 0x86, 0x21, 0xeb, 0xe3,
	0xe9, 0xee, 0x11, 0x22, 0xc0, 0x28, 0x24, 0xc9, 0x25, 0x8a, 0xa9, 0xd3, 0x31, 0xc8, 0x5b, 0x80,
	0xc9, 0xc9, 0x82, 0x4b, 0x26, 0x33, 0x5f, 0x0f, 0x1e, 0x3a, 0x50, 0x0a, 0x1c, 0x0c, 0x20, 0xc7,
	0x39, 0xb1, 0x06, 0x24, 0x39, 0x16, 0x14, 0x44, 0x81, 0x42, 0x64, 0x11, 0x13, 0xb7, 0x2f, 0x58,
	0xda, 0xb1, 0xa0, 0x20, 0xc0, 0x29, 0x89, 0x0d, 0xac, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x37, 0x6c, 0x45, 0xcd, 0x5b, 0x01, 0x00, 0x00,
}