// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/resource/v1alpha2/resource.proto

package v1alpha2 // import "github.com/containers-ai/api/datahub/resource/v1alpha2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha2 "github.com/containers-ai/api/datahub/metric/v1alpha2"
import v1alpha22 "github.com/containers-ai/api/datahub/prediction/v1alpha2"
import v1alpha21 "github.com/containers-ai/api/datahub/resource/metadata/v1alpha2"
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

// *
// Represents a container and its containing limit and requeset configurations
//
type Container struct {
	Name                 string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitResource        map[int32]*v1alpha2.MetricData `protobuf:"bytes,2,rep,name=limit_resource,json=limitResource,proto3" json:"limit_resource,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequestResource      map[int32]*v1alpha2.MetricData `protobuf:"bytes,3,rep,name=request_resource,json=requestResource,proto3" json:"request_resource,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_03b9abd168d75ba3, []int{0}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetLimitResource() map[int32]*v1alpha2.MetricData {
	if m != nil {
		return m.LimitResource
	}
	return nil
}

func (m *Container) GetRequestResource() map[int32]*v1alpha2.MetricData {
	if m != nil {
		return m.RequestResource
	}
	return nil
}

// *
// Represents a Kubernetes pod
//
type Pod struct {
	NamespacedName       *v1alpha21.NamespacedName      `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ResourceLink         string                         `protobuf:"bytes,2,opt,name=resource_link,json=resourceLink,proto3" json:"resource_link,omitempty"`
	Containers           []*Container                   `protobuf:"bytes,3,rep,name=containers,proto3" json:"containers,omitempty"`
	IsPredicted          bool                           `protobuf:"varint,4,opt,name=is_predicted,json=isPredicted,proto3" json:"is_predicted,omitempty"`
	Scaler               *v1alpha21.NamespacedName      `protobuf:"bytes,5,opt,name=scaler,proto3" json:"scaler,omitempty"`
	NodeName             string                         `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	StartTime            *timestamp.Timestamp           `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Policy               v1alpha22.RecommendationPolicy `protobuf:"varint,8,opt,name=policy,proto3,enum=containersai.datahub.prediction.v1alpha2.RecommendationPolicy" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_03b9abd168d75ba3, []int{1}
}
func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (dst *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(dst, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetNamespacedName() *v1alpha21.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *Pod) GetResourceLink() string {
	if m != nil {
		return m.ResourceLink
	}
	return ""
}

func (m *Pod) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Pod) GetIsPredicted() bool {
	if m != nil {
		return m.IsPredicted
	}
	return false
}

func (m *Pod) GetScaler() *v1alpha21.NamespacedName {
	if m != nil {
		return m.Scaler
	}
	return nil
}

func (m *Pod) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Pod) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Pod) GetPolicy() v1alpha22.RecommendationPolicy {
	if m != nil {
		return m.Policy
	}
	return v1alpha22.RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED
}

// *
// Represents a Kubernetes node
//
type Node struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_03b9abd168d75ba3, []int{2}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Container)(nil), "containersai.datahub.resource.v1alpha2.Container")
	proto.RegisterMapType((map[int32]*v1alpha2.MetricData)(nil), "containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry")
	proto.RegisterMapType((map[int32]*v1alpha2.MetricData)(nil), "containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry")
	proto.RegisterType((*Pod)(nil), "containersai.datahub.resource.v1alpha2.Pod")
	proto.RegisterType((*Node)(nil), "containersai.datahub.resource.v1alpha2.Node")
}

func init() {
	proto.RegisterFile("datahub/resource/v1alpha2/resource.proto", fileDescriptor_resource_03b9abd168d75ba3)
}

var fileDescriptor_resource_03b9abd168d75ba3 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x0f, 0x92, 0x49, 0x9b, 0x56, 0x2b, 0x0e, 0x96, 0x39, 0x10, 0x02, 0x42, 0x3e,
	0x80, 0x4d, 0x82, 0x84, 0x0a, 0x07, 0x90, 0xa0, 0xf4, 0x54, 0xaa, 0xb0, 0x42, 0x20, 0x71, 0x89,
	0x36, 0xeb, 0x25, 0x59, 0xc5, 0xde, 0x75, 0xd7, 0x9b, 0x4a, 0xf9, 0x1d, 0xfc, 0x4a, 0xfe, 0x05,
	0xda, 0xb5, 0xd7, 0x49, 0xd4, 0x20, 0xa8, 0x44, 0x6f, 0x3b, 0x4f, 0x33, 0xef, 0xcd, 0x9b, 0xf1,
	0x18, 0xc2, 0x84, 0x68, 0xb2, 0x58, 0xcd, 0x62, 0xc5, 0x0a, 0xb9, 0x52, 0x94, 0xc5, 0xd7, 0x23,
	0x92, 0xe6, 0x0b, 0x32, 0xae, 0x91, 0x28, 0x57, 0x52, 0x4b, 0xf4, 0x94, 0x4a, 0xa1, 0x09, 0x17,
	0x4c, 0x15, 0x84, 0x47, 0x55, 0x59, 0x54, 0x27, 0xb9, 0xb2, 0x60, 0x74, 0x83, 0x31, 0x63, 0x9a,
	0x18, 0x70, 0x43, 0xed, 0x90, 0x92, 0x3a, 0x78, 0xe6, 0x4a, 0x72, 0xc5, 0x12, 0x4e, 0x35, 0x97,
	0x62, 0x93, 0xbb, 0xc1, 0xaa, 0xec, 0x87, 0x73, 0x29, 0xe7, 0x29, 0x8b, 0x6d, 0x34, 0x5b, 0xfd,
	0x88, 0x35, 0xcf, 0x58, 0xa1, 0x49, 0x96, 0x57, 0x09, 0x4f, 0x1c, 0x5d, 0xc6, 0xb4, 0xe2, 0x74,
	0x47, 0x56, 0x71, 0x5a, 0x66, 0x0d, 0x7f, 0x35, 0xa0, 0xfb, 0xc1, 0x59, 0x42, 0x08, 0x9a, 0x82,
	0x64, 0xcc, 0xf7, 0x06, 0x5e, 0xd8, 0xc5, 0xf6, 0x8d, 0x96, 0xd0, 0x4f, 0x79, 0xc6, 0xf5, 0xd4,
	0x39, 0xf1, 0x0f, 0x06, 0x8d, 0xb0, 0x37, 0x3e, 0x8b, 0xfe, 0x6d, 0x14, 0x51, 0x4d, 0x1f, 0x5d,
	0x18, 0x1e, 0x5c, 0x25, 0x7c, 0x14, 0x5a, 0xad, 0xf1, 0x51, 0xba, 0x8d, 0xa1, 0x2b, 0x38, 0x51,
	0xec, 0x6a, 0xc5, 0x8a, 0x2d, 0xb9, 0x86, 0x95, 0x3b, 0xbf, 0xbd, 0x1c, 0x2e, 0x99, 0x76, 0x05,
	0x8f, 0xd5, 0x2e, 0x1a, 0x28, 0x40, 0x37, 0xfb, 0x42, 0x27, 0xd0, 0x58, 0xb2, 0xb5, 0x1d, 0x44,
	0x0b, 0x9b, 0x27, 0x3a, 0x87, 0xd6, 0x35, 0x49, 0x57, 0xc6, 0xbe, 0x17, 0xf6, 0xc6, 0x2f, 0xf6,
	0xf7, 0x53, 0x0d, 0xb7, 0xee, 0xe6, 0x93, 0x8d, 0xcf, 0x88, 0x26, 0xb8, 0x2c, 0x7f, 0x73, 0x70,
	0xea, 0x05, 0x1a, 0xee, 0xef, 0x6b, 0xee, 0x6e, 0x55, 0x87, 0x3f, 0x9b, 0xd0, 0x98, 0xc8, 0x04,
	0x2d, 0xe0, 0xd8, 0x6c, 0xb6, 0xc8, 0x09, 0x65, 0xc9, 0xb4, 0x5e, 0x78, 0x6f, 0xfc, 0xee, 0x2f,
	0x33, 0xae, 0x3f, 0xd8, 0x5a, 0xe8, 0xb2, 0xe6, 0x31, 0x2f, 0xdc, 0x17, 0x3b, 0x31, 0x7a, 0x0c,
	0x47, 0xae, 0x78, 0x9a, 0x72, 0xb1, 0xb4, 0x2e, 0xba, 0xf8, 0xd0, 0x81, 0x17, 0x5c, 0x2c, 0xd1,
	0x67, 0x80, 0x8d, 0x6c, 0xb5, 0xed, 0xd1, 0xad, 0xb7, 0x8d, 0xb7, 0x48, 0xd0, 0x23, 0x38, 0xe4,
	0xc5, 0xb4, 0xba, 0x19, 0x96, 0xf8, 0xcd, 0x81, 0x17, 0x76, 0x70, 0x8f, 0x17, 0x13, 0x07, 0xa1,
	0x6f, 0xd0, 0x2e, 0x28, 0x49, 0x99, 0xf2, 0x5b, 0xff, 0xc7, 0x7b, 0x45, 0x87, 0x1e, 0x40, 0x57,
	0xc8, 0x84, 0x95, 0x73, 0x6d, 0x5b, 0xbf, 0x1d, 0x03, 0xd8, 0x81, 0xbc, 0x06, 0x28, 0x34, 0x51,
	0x7a, 0x6a, 0xae, 0xd5, 0xbf, 0x67, 0x95, 0x83, 0xa8, 0x3c, 0xe5, 0xc8, 0x9d, 0x72, 0xf4, 0xc5,
	0x9d, 0x32, 0xee, 0xda, 0x6c, 0x13, 0xa3, 0xaf, 0xd0, 0xce, 0x65, 0xca, 0xe9, 0xda, 0xef, 0x0c,
	0xbc, 0xb0, 0x3f, 0x7e, 0xbb, 0xbf, 0xe1, 0xad, 0x1f, 0x45, 0xdd, 0x29, 0x66, 0x54, 0x66, 0x19,
	0x13, 0x09, 0x31, 0xf8, 0xc4, 0xb2, 0xe0, 0x8a, 0x6d, 0x18, 0x40, 0xf3, 0x52, 0x26, 0x6c, 0xdf,
	0xed, 0xbf, 0x3f, 0xfd, 0xfe, 0x6a, 0xce, 0xb5, 0xa1, 0xa5, 0x32, 0x8b, 0x37, 0x7a, 0xcf, 0x09,
	0x8f, 0x49, 0xce, 0xe3, 0x3f, 0xfe, 0x36, 0x67, 0x6d, 0x6b, 0xe6, 0xe5, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0xdb, 0x1b, 0x31, 0x5a, 0x05, 0x00, 0x00,
}
