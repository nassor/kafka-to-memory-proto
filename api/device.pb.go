// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Device struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationId       string               `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Tz                   string               `protobuf:"bytes,3,opt,name=tz,proto3" json:"tz,omitempty"`
	InstalledAt          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=installed_at,json=installedAt,proto3" json:"installed_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_20117bc49b71fb14, []int{0}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *Device) GetTz() string {
	if m != nil {
		return m.Tz
	}
	return ""
}

func (m *Device) GetInstalledAt() *timestamp.Timestamp {
	if m != nil {
		return m.InstalledAt
	}
	return nil
}

func (m *Device) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "api.Device")
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_device_20117bc49b71fb14) }

var fileDescriptor_device_20117bc49b71fb14 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x4f,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6,
	0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x54, 0x29, 0x9d, 0x66, 0xe4, 0x62, 0x73, 0x01, 0x6b, 0x13,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11,
	0x52, 0xe7, 0xe2, 0xcf, 0x2f, 0x4a, 0x4f, 0xcc, 0xcb, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b,
	0xcf, 0x4c, 0x91, 0x60, 0x02, 0x4b, 0xf2, 0x21, 0x0b, 0x7b, 0xa6, 0x80, 0x34, 0x96, 0x54, 0x49,
	0x30, 0x43, 0x34, 0x96, 0x54, 0x09, 0xd9, 0x72, 0xf1, 0x64, 0xe6, 0x15, 0x97, 0x24, 0xe6, 0xe4,
	0xa4, 0xa6, 0xc4, 0x27, 0x96, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xe9, 0x41, 0xdc,
	0xa2, 0x07, 0x73, 0x8b, 0x5e, 0x08, 0xcc, 0x2d, 0x41, 0xdc, 0x70, 0xf5, 0x8e, 0x25, 0x42, 0x96,
	0x5c, 0x5c, 0xa5, 0x05, 0x29, 0x89, 0x25, 0x10, 0xcd, 0xac, 0x04, 0x35, 0x73, 0x42, 0x55, 0x3b,
	0x96, 0x38, 0xb1, 0x7b, 0x30, 0x46, 0x81, 0xfc, 0x9d, 0xc4, 0x06, 0x56, 0x67, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x5f, 0x75, 0x8b, 0x40, 0x13, 0x01, 0x00, 0x00,
}