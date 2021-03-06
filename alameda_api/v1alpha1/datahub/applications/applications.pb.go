// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/applications/applications.proto

package applications

import (
	fmt "fmt"
	schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	common "github.com/containers-ai/api/common"
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

type ReadApplication struct {
	SchemaMeta           *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	Groups               []*common.Group     `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ReadApplication) Reset()         { *m = ReadApplication{} }
func (m *ReadApplication) String() string { return proto.CompactTextString(m) }
func (*ReadApplication) ProtoMessage()    {}
func (*ReadApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_d64a8ef2ec8cbd56, []int{0}
}

func (m *ReadApplication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadApplication.Unmarshal(m, b)
}
func (m *ReadApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadApplication.Marshal(b, m, deterministic)
}
func (m *ReadApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadApplication.Merge(m, src)
}
func (m *ReadApplication) XXX_Size() int {
	return xxx_messageInfo_ReadApplication.Size(m)
}
func (m *ReadApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadApplication.DiscardUnknown(m)
}

var xxx_messageInfo_ReadApplication proto.InternalMessageInfo

func (m *ReadApplication) GetSchemaMeta() *schemas.SchemaMeta {
	if m != nil {
		return m.SchemaMeta
	}
	return nil
}

func (m *ReadApplication) GetGroups() []*common.Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type WriteApplication struct {
	SchemaMeta           *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	Columns              []string            `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows                 []*common.Row       `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WriteApplication) Reset()         { *m = WriteApplication{} }
func (m *WriteApplication) String() string { return proto.CompactTextString(m) }
func (*WriteApplication) ProtoMessage()    {}
func (*WriteApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_d64a8ef2ec8cbd56, []int{1}
}

func (m *WriteApplication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteApplication.Unmarshal(m, b)
}
func (m *WriteApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteApplication.Marshal(b, m, deterministic)
}
func (m *WriteApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteApplication.Merge(m, src)
}
func (m *WriteApplication) XXX_Size() int {
	return xxx_messageInfo_WriteApplication.Size(m)
}
func (m *WriteApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteApplication.DiscardUnknown(m)
}

var xxx_messageInfo_WriteApplication proto.InternalMessageInfo

func (m *WriteApplication) GetSchemaMeta() *schemas.SchemaMeta {
	if m != nil {
		return m.SchemaMeta
	}
	return nil
}

func (m *WriteApplication) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *WriteApplication) GetRows() []*common.Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadApplication)(nil), "containersai.alameda.v1alpha1.datahub.applications.ReadApplication")
	proto.RegisterType((*WriteApplication)(nil), "containersai.alameda.v1alpha1.datahub.applications.WriteApplication")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/applications/applications.proto", fileDescriptor_d64a8ef2ec8cbd56)
}

var fileDescriptor_d64a8ef2ec8cbd56 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0x95, 0xaf, 0x9f, 0x8a, 0x70, 0x06, 0x50, 0x58, 0xa2, 0x4c, 0x55, 0xa7, 0x0e, 0x60,
	0xd3, 0x30, 0x21, 0xb1, 0xc0, 0x00, 0x2c, 0x2c, 0x66, 0x40, 0xea, 0x52, 0x5d, 0x1d, 0xab, 0xb1,
	0x14, 0xfb, 0x2c, 0xdb, 0xa1, 0xe2, 0x71, 0x78, 0x09, 0x9e, 0x0f, 0x25, 0x71, 0x29, 0x95, 0x10,
	0xea, 0xc4, 0x74, 0x3e, 0xe9, 0xef, 0xdf, 0xfd, 0xac, 0x33, 0xb9, 0x81, 0x06, 0xb4, 0xac, 0x60,
	0x09, 0x56, 0xb1, 0xd7, 0x39, 0x34, 0xb6, 0x86, 0x39, 0xab, 0x20, 0x40, 0xdd, 0xae, 0x18, 0x58,
	0xdb, 0x28, 0x01, 0x41, 0xa1, 0xf1, 0x7b, 0x0d, 0xb5, 0x0e, 0x03, 0x66, 0xa5, 0x40, 0x13, 0x40,
	0x19, 0xe9, 0x3c, 0x28, 0x1a, 0x51, 0x74, 0x8b, 0xa1, 0x11, 0x43, 0xbf, 0xdf, 0x2c, 0xce, 0x04,
	0x6a, 0x8d, 0x86, 0x0d, 0x65, 0x00, 0x15, 0x97, 0xbf, 0x6a, 0x78, 0x51, 0x4b, 0x0d, 0x9e, 0x85,
	0x37, 0x2b, 0xe3, 0xe8, 0xe9, 0x7b, 0x42, 0x4e, 0xb8, 0x84, 0xea, 0x76, 0xc7, 0xce, 0x16, 0x24,
	0x1d, 0xa2, 0x4b, 0x2d, 0x03, 0xe4, 0xc9, 0x24, 0x99, 0xa5, 0xe5, 0x35, 0x3d, 0x4c, 0x32, 0x0e,
	0xa1, 0xcf, 0x7d, 0x7d, 0x92, 0x01, 0x38, 0xf1, 0x5f, 0xe7, 0xac, 0x24, 0xe3, 0xb5, 0xc3, 0xd6,
	0xfa, 0xfc, 0xdf, 0x64, 0x34, 0x4b, 0xcb, 0x62, 0x1f, 0x1b, 0x5f, 0xf3, 0xd0, 0x45, 0x78, 0x4c,
	0x4e, 0x3f, 0x12, 0x72, 0xfa, 0xe2, 0x54, 0x90, 0x7f, 0x25, 0x99, 0x93, 0x23, 0x81, 0x4d, 0xab,
	0xcd, 0x60, 0x79, 0xcc, 0xb7, 0x6d, 0x76, 0x4e, 0xfe, 0x3b, 0xdc, 0xf8, 0x7c, 0xd4, 0xcb, 0xe7,
	0x3f, 0xca, 0x73, 0xdc, 0xf0, 0x3e, 0x75, 0xf7, 0xb8, 0xb8, 0x5f, 0xab, 0xd0, 0x0d, 0x14, 0xa8,
	0xd9, 0x2e, 0x7b, 0x01, 0x8a, 0x75, 0x1b, 0x3a, 0xf8, 0xd3, 0xac, 0xc6, 0xfd, 0xb6, 0xae, 0x3e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xa3, 0xfa, 0xf6, 0x68, 0x02, 0x00, 0x00,
}
