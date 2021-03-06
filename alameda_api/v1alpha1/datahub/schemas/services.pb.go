// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/schemas/services.proto

package schemas

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type CreateSchemasRequest struct {
	Schemas              []*Schema `protobuf:"bytes,1,rep,name=schemas,proto3" json:"schemas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateSchemasRequest) Reset()         { *m = CreateSchemasRequest{} }
func (m *CreateSchemasRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSchemasRequest) ProtoMessage()    {}
func (*CreateSchemasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c53eb3dee73a92, []int{0}
}

func (m *CreateSchemasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSchemasRequest.Unmarshal(m, b)
}
func (m *CreateSchemasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSchemasRequest.Marshal(b, m, deterministic)
}
func (m *CreateSchemasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSchemasRequest.Merge(m, src)
}
func (m *CreateSchemasRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSchemasRequest.Size(m)
}
func (m *CreateSchemasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSchemasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSchemasRequest proto.InternalMessageInfo

func (m *CreateSchemasRequest) GetSchemas() []*Schema {
	if m != nil {
		return m.Schemas
	}
	return nil
}

type ListSchemasRequest struct {
	SchemaMetas          []*SchemaMeta `protobuf:"bytes,1,rep,name=schema_metas,json=schemaMetas,proto3" json:"schema_metas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListSchemasRequest) Reset()         { *m = ListSchemasRequest{} }
func (m *ListSchemasRequest) String() string { return proto.CompactTextString(m) }
func (*ListSchemasRequest) ProtoMessage()    {}
func (*ListSchemasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c53eb3dee73a92, []int{1}
}

func (m *ListSchemasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSchemasRequest.Unmarshal(m, b)
}
func (m *ListSchemasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSchemasRequest.Marshal(b, m, deterministic)
}
func (m *ListSchemasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSchemasRequest.Merge(m, src)
}
func (m *ListSchemasRequest) XXX_Size() int {
	return xxx_messageInfo_ListSchemasRequest.Size(m)
}
func (m *ListSchemasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSchemasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSchemasRequest proto.InternalMessageInfo

func (m *ListSchemasRequest) GetSchemaMetas() []*SchemaMeta {
	if m != nil {
		return m.SchemaMetas
	}
	return nil
}

type ListSchemasResponse struct {
	Status               *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Schemas              []*Schema      `protobuf:"bytes,2,rep,name=schemas,proto3" json:"schemas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListSchemasResponse) Reset()         { *m = ListSchemasResponse{} }
func (m *ListSchemasResponse) String() string { return proto.CompactTextString(m) }
func (*ListSchemasResponse) ProtoMessage()    {}
func (*ListSchemasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c53eb3dee73a92, []int{2}
}

func (m *ListSchemasResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSchemasResponse.Unmarshal(m, b)
}
func (m *ListSchemasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSchemasResponse.Marshal(b, m, deterministic)
}
func (m *ListSchemasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSchemasResponse.Merge(m, src)
}
func (m *ListSchemasResponse) XXX_Size() int {
	return xxx_messageInfo_ListSchemasResponse.Size(m)
}
func (m *ListSchemasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSchemasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSchemasResponse proto.InternalMessageInfo

func (m *ListSchemasResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListSchemasResponse) GetSchemas() []*Schema {
	if m != nil {
		return m.Schemas
	}
	return nil
}

type DeleteSchemasRequest struct {
	SchemaMetas          []*SchemaMeta `protobuf:"bytes,1,rep,name=schema_metas,json=schemaMetas,proto3" json:"schema_metas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeleteSchemasRequest) Reset()         { *m = DeleteSchemasRequest{} }
func (m *DeleteSchemasRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSchemasRequest) ProtoMessage()    {}
func (*DeleteSchemasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c53eb3dee73a92, []int{3}
}

func (m *DeleteSchemasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSchemasRequest.Unmarshal(m, b)
}
func (m *DeleteSchemasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSchemasRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSchemasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSchemasRequest.Merge(m, src)
}
func (m *DeleteSchemasRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSchemasRequest.Size(m)
}
func (m *DeleteSchemasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSchemasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSchemasRequest proto.InternalMessageInfo

func (m *DeleteSchemasRequest) GetSchemaMetas() []*SchemaMeta {
	if m != nil {
		return m.SchemaMetas
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSchemasRequest)(nil), "containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest")
	proto.RegisterType((*ListSchemasRequest)(nil), "containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest")
	proto.RegisterType((*ListSchemasResponse)(nil), "containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse")
	proto.RegisterType((*DeleteSchemasRequest)(nil), "containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/schemas/services.proto", fileDescriptor_96c53eb3dee73a92)
}

var fileDescriptor_96c53eb3dee73a92 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xd9, 0xef, 0x17, 0x2a, 0xa4, 0x9e, 0xd6, 0x82, 0xa5, 0xa7, 0xd2, 0x53, 0x11, 0x9a,
	0xd8, 0x16, 0x0f, 0x1e, 0xb5, 0x1e, 0x15, 0x61, 0x7b, 0x13, 0xa1, 0x4c, 0xd3, 0x61, 0x1b, 0xd8,
	0x6d, 0x62, 0x66, 0xb6, 0xe0, 0xbf, 0xe1, 0x5f, 0x2c, 0x6d, 0xb2, 0xf5, 0xd7, 0x65, 0x15, 0x3d,
	0x25, 0x30, 0xf3, 0xde, 0x4b, 0x3e, 0x3c, 0x31, 0x85, 0x02, 0x4a, 0x5c, 0xc1, 0x02, 0x9c, 0x51,
	0xdb, 0x31, 0x14, 0x6e, 0x0d, 0x63, 0xb5, 0x02, 0x86, 0x75, 0xb5, 0x54, 0xa4, 0xd7, 0x58, 0x02,
	0x29, 0x42, 0xbf, 0x35, 0x1a, 0x49, 0x3a, 0x6f, 0xd9, 0xa6, 0x23, 0x6d, 0x37, 0x0c, 0x66, 0x83,
	0x9e, 0xc0, 0xc8, 0xe8, 0x20, 0x6b, 0xb5, 0x8c, 0x6a, 0x19, 0xd5, 0xbd, 0x49, 0xb3, 0x8c, 0x70,
	0x86, 0x88, 0xde, 0x79, 0x23, 0x0d, 0x3f, 0xbb, 0xfa, 0x51, 0xbd, 0xd3, 0xdc, 0xda, 0xbc, 0x40,
	0xe5, 0x9d, 0x56, 0xc4, 0xc0, 0x55, 0x1c, 0x0c, 0x72, 0xd1, 0x99, 0x79, 0x04, 0xc6, 0x79, 0x50,
	0x65, 0xf8, 0x54, 0x21, 0x71, 0x7a, 0x2f, 0x8e, 0xa2, 0x4f, 0x37, 0xe9, 0xff, 0x1f, 0xb6, 0x27,
	0x17, 0xf2, 0x5b, 0xff, 0x92, 0xc1, 0x2f, 0xab, 0x5d, 0x06, 0x5e, 0xa4, 0xb7, 0x86, 0xf8, 0x53,
	0xcc, 0xa3, 0x38, 0x0e, 0x0b, 0x8b, 0x12, 0xf9, 0x90, 0x75, 0xf9, 0xa3, 0xac, 0x3b, 0x64, 0xc8,
	0xda, 0x74, 0xb8, 0xd3, 0xe0, 0x25, 0x11, 0x27, 0x1f, 0x42, 0xc9, 0xd9, 0x0d, 0x61, 0x7a, 0x26,
	0x5a, 0x01, 0x42, 0x37, 0xe9, 0x27, 0xc3, 0xf6, 0x24, 0x95, 0x01, 0x8f, 0xf4, 0x4e, 0xcb, 0xf9,
	0x7e, 0x92, 0xc5, 0x8d, 0xf7, 0x20, 0xfe, 0xfd, 0x0a, 0x08, 0x16, 0x9d, 0x1b, 0x2c, 0xf0, 0x0b,
	0xf1, 0x3f, 0x45, 0x71, 0x3d, 0x7b, 0xb8, 0xca, 0x0d, 0xef, 0x36, 0xb5, 0x2d, 0xd5, 0x9b, 0xe7,
	0x08, 0x8c, 0xda, 0xb5, 0xa8, 0x49, 0xa3, 0x96, 0xad, 0x7d, 0x67, 0xa6, 0xaf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0xc5, 0x6f, 0xd2, 0x18, 0x03, 0x00, 0x00,
}
