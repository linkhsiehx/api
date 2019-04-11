// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/types.proto

package containers_ai_alameda_v1alpha1_datahub

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

// The valid statuses of pods
type PodPhase int32

const (
	PodPhase_Pending          PodPhase = 0
	PodPhase_Running          PodPhase = 1
	PodPhase_Succeeded        PodPhase = 2
	PodPhase_Failed           PodPhase = 3
	PodPhase_Unknown          PodPhase = 4
	PodPhase_Completed        PodPhase = 5
	PodPhase_CrashLoopBackOff PodPhase = 6
)

var PodPhase_name = map[int32]string{
	0: "Pending",
	1: "Running",
	2: "Succeeded",
	3: "Failed",
	4: "Unknown",
	5: "Completed",
	6: "CrashLoopBackOff",
}
var PodPhase_value = map[string]int32{
	"Pending":          0,
	"Running":          1,
	"Succeeded":        2,
	"Failed":           3,
	"Unknown":          4,
	"Completed":        5,
	"CrashLoopBackOff": 6,
}

func (x PodPhase) String() string {
	return proto.EnumName(PodPhase_name, int32(x))
}
func (PodPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_027c8215935d06a7, []int{0}
}

// ContainerStateWaiting is a waiting state of a container
type ContainerStateWaiting struct {
	Reason               string   `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerStateWaiting) Reset()         { *m = ContainerStateWaiting{} }
func (m *ContainerStateWaiting) String() string { return proto.CompactTextString(m) }
func (*ContainerStateWaiting) ProtoMessage()    {}
func (*ContainerStateWaiting) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_027c8215935d06a7, []int{0}
}
func (m *ContainerStateWaiting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerStateWaiting.Unmarshal(m, b)
}
func (m *ContainerStateWaiting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerStateWaiting.Marshal(b, m, deterministic)
}
func (dst *ContainerStateWaiting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerStateWaiting.Merge(dst, src)
}
func (m *ContainerStateWaiting) XXX_Size() int {
	return xxx_messageInfo_ContainerStateWaiting.Size(m)
}
func (m *ContainerStateWaiting) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerStateWaiting.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerStateWaiting proto.InternalMessageInfo

func (m *ContainerStateWaiting) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *ContainerStateWaiting) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// ContainerStateRunning is a running state of a container
type ContainerStateRunning struct {
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ContainerStateRunning) Reset()         { *m = ContainerStateRunning{} }
func (m *ContainerStateRunning) String() string { return proto.CompactTextString(m) }
func (*ContainerStateRunning) ProtoMessage()    {}
func (*ContainerStateRunning) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_027c8215935d06a7, []int{1}
}
func (m *ContainerStateRunning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerStateRunning.Unmarshal(m, b)
}
func (m *ContainerStateRunning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerStateRunning.Marshal(b, m, deterministic)
}
func (dst *ContainerStateRunning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerStateRunning.Merge(dst, src)
}
func (m *ContainerStateRunning) XXX_Size() int {
	return xxx_messageInfo_ContainerStateRunning.Size(m)
}
func (m *ContainerStateRunning) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerStateRunning.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerStateRunning proto.InternalMessageInfo

func (m *ContainerStateRunning) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

// ContainerStateTerminated is a terminated state of a container
type ContainerStateTerminated struct {
	ExitCode             int32                `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Reason               string               `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ContainerStateTerminated) Reset()         { *m = ContainerStateTerminated{} }
func (m *ContainerStateTerminated) String() string { return proto.CompactTextString(m) }
func (*ContainerStateTerminated) ProtoMessage()    {}
func (*ContainerStateTerminated) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_027c8215935d06a7, []int{2}
}
func (m *ContainerStateTerminated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerStateTerminated.Unmarshal(m, b)
}
func (m *ContainerStateTerminated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerStateTerminated.Marshal(b, m, deterministic)
}
func (dst *ContainerStateTerminated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerStateTerminated.Merge(dst, src)
}
func (m *ContainerStateTerminated) XXX_Size() int {
	return xxx_messageInfo_ContainerStateTerminated.Size(m)
}
func (m *ContainerStateTerminated) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerStateTerminated.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerStateTerminated proto.InternalMessageInfo

func (m *ContainerStateTerminated) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *ContainerStateTerminated) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *ContainerStateTerminated) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ContainerStateTerminated) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *ContainerStateTerminated) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

// ContainerState holds a possible state of container
type ContainerState struct {
	Waiting              *ContainerStateWaiting    `protobuf:"bytes,1,opt,name=waiting,proto3" json:"waiting,omitempty"`
	Running              *ContainerStateRunning    `protobuf:"bytes,2,opt,name=running,proto3" json:"running,omitempty"`
	Terminated           *ContainerStateTerminated `protobuf:"bytes,3,opt,name=terminated,proto3" json:"terminated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ContainerState) Reset()         { *m = ContainerState{} }
func (m *ContainerState) String() string { return proto.CompactTextString(m) }
func (*ContainerState) ProtoMessage()    {}
func (*ContainerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_027c8215935d06a7, []int{3}
}
func (m *ContainerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerState.Unmarshal(m, b)
}
func (m *ContainerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerState.Marshal(b, m, deterministic)
}
func (dst *ContainerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerState.Merge(dst, src)
}
func (m *ContainerState) XXX_Size() int {
	return xxx_messageInfo_ContainerState.Size(m)
}
func (m *ContainerState) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerState.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerState proto.InternalMessageInfo

func (m *ContainerState) GetWaiting() *ContainerStateWaiting {
	if m != nil {
		return m.Waiting
	}
	return nil
}

func (m *ContainerState) GetRunning() *ContainerStateRunning {
	if m != nil {
		return m.Running
	}
	return nil
}

func (m *ContainerState) GetTerminated() *ContainerStateTerminated {
	if m != nil {
		return m.Terminated
	}
	return nil
}

// ContainerStatus contains details for the current status of this container
type ContainerStatus struct {
	State                *ContainerState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	LastTerminationState *ContainerState `protobuf:"bytes,2,opt,name=last_termination_state,json=lastTerminationState,proto3" json:"last_termination_state,omitempty"`
	RestartCount         int32           `protobuf:"varint,3,opt,name=restart_count,json=restartCount,proto3" json:"restart_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ContainerStatus) Reset()         { *m = ContainerStatus{} }
func (m *ContainerStatus) String() string { return proto.CompactTextString(m) }
func (*ContainerStatus) ProtoMessage()    {}
func (*ContainerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_027c8215935d06a7, []int{4}
}
func (m *ContainerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerStatus.Unmarshal(m, b)
}
func (m *ContainerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerStatus.Marshal(b, m, deterministic)
}
func (dst *ContainerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerStatus.Merge(dst, src)
}
func (m *ContainerStatus) XXX_Size() int {
	return xxx_messageInfo_ContainerStatus.Size(m)
}
func (m *ContainerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerStatus proto.InternalMessageInfo

func (m *ContainerStatus) GetState() *ContainerState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ContainerStatus) GetLastTerminationState() *ContainerState {
	if m != nil {
		return m.LastTerminationState
	}
	return nil
}

func (m *ContainerStatus) GetRestartCount() int32 {
	if m != nil {
		return m.RestartCount
	}
	return 0
}

// PodStatus represents information about the status of a pod
type PodStatus struct {
	Phase                PodPhase `protobuf:"varint,1,opt,name=phase,proto3,enum=containers_ai.alameda.v1alpha1.datahub.PodPhase" json:"phase,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Reason               string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PodStatus) Reset()         { *m = PodStatus{} }
func (m *PodStatus) String() string { return proto.CompactTextString(m) }
func (*PodStatus) ProtoMessage()    {}
func (*PodStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_027c8215935d06a7, []int{5}
}
func (m *PodStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodStatus.Unmarshal(m, b)
}
func (m *PodStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodStatus.Marshal(b, m, deterministic)
}
func (dst *PodStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodStatus.Merge(dst, src)
}
func (m *PodStatus) XXX_Size() int {
	return xxx_messageInfo_PodStatus.Size(m)
}
func (m *PodStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PodStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PodStatus proto.InternalMessageInfo

func (m *PodStatus) GetPhase() PodPhase {
	if m != nil {
		return m.Phase
	}
	return PodPhase_Pending
}

func (m *PodStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PodStatus) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*ContainerStateWaiting)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerStateWaiting")
	proto.RegisterType((*ContainerStateRunning)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerStateRunning")
	proto.RegisterType((*ContainerStateTerminated)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerStateTerminated")
	proto.RegisterType((*ContainerState)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerState")
	proto.RegisterType((*ContainerStatus)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerStatus")
	proto.RegisterType((*PodStatus)(nil), "containers_ai.alameda.v1alpha1.datahub.PodStatus")
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.PodPhase", PodPhase_name, PodPhase_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/types.proto", fileDescriptor_types_027c8215935d06a7)
}

var fileDescriptor_types_027c8215935d06a7 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0xdd, 0xa4, 0x9b, 0x34, 0x27, 0xb6, 0x2e, 0x43, 0x2d, 0xa1, 0x5e, 0x28, 0x2b, 0x48,
	0xf1, 0x62, 0x63, 0x23, 0x08, 0x22, 0x82, 0x35, 0x50, 0x10, 0x0a, 0x86, 0x6d, 0xa4, 0x97, 0xeb,
	0xc9, 0xce, 0x49, 0x32, 0x74, 0x77, 0x66, 0xdd, 0x99, 0xb5, 0xfa, 0x00, 0xbe, 0x89, 0x2f, 0xe6,
	0x2b, 0xf8, 0x04, 0xb2, 0x7f, 0xc6, 0x26, 0xc5, 0x62, 0x8d, 0x97, 0x33, 0x7b, 0xbe, 0xdf, 0x7e,
	0xe7, 0x7c, 0x73, 0xe0, 0x10, 0x13, 0x4c, 0x89, 0x63, 0x84, 0x99, 0x18, 0x7e, 0x3e, 0xc2, 0x24,
	0x5b, 0xe2, 0xd1, 0x90, 0xa3, 0xc1, 0x65, 0x31, 0x1b, 0x9a, 0xaf, 0x19, 0xe9, 0x20, 0xcb, 0x95,
	0x51, 0xec, 0x49, 0xac, 0xa4, 0x41, 0x21, 0x29, 0xd7, 0x11, 0x8a, 0xa0, 0xd1, 0x05, 0x56, 0x13,
	0x34, 0x9a, 0x83, 0x87, 0x0b, 0xa5, 0x16, 0x09, 0x0d, 0x2b, 0xd5, 0xac, 0x98, 0x0f, 0x8d, 0x48,
	0x49, 0x1b, 0x4c, 0xb3, 0x1a, 0xe4, 0xbf, 0x83, 0xfb, 0x63, 0x8b, 0x3a, 0x33, 0x68, 0xe8, 0x1c,
	0x85, 0x11, 0x72, 0xc1, 0xf6, 0xa1, 0x93, 0x13, 0x6a, 0x25, 0x07, 0xce, 0x23, 0xe7, 0xb0, 0x17,
	0x36, 0x27, 0x36, 0x80, 0x6e, 0x4a, 0x5a, 0xe3, 0x82, 0x06, 0xad, 0xea, 0x83, 0x3d, 0xfa, 0xe1,
	0x75, 0x54, 0x58, 0x48, 0x59, 0xa2, 0x5e, 0x02, 0x68, 0x83, 0xb9, 0x21, 0x1e, 0xa1, 0xa9, 0x70,
	0xfd, 0xd1, 0x41, 0x50, 0x3b, 0x0b, 0xac, 0xb3, 0x60, 0x6a, 0x9d, 0x85, 0xbd, 0xa6, 0xfa, 0xd8,
	0xf8, 0x3f, 0x1c, 0x18, 0xac, 0x43, 0xa7, 0x94, 0xa7, 0x42, 0xa2, 0x21, 0xce, 0x1e, 0x40, 0x8f,
	0xbe, 0x08, 0x13, 0xc5, 0x8a, 0x53, 0x85, 0x75, 0xc3, 0xed, 0xf2, 0x62, 0xac, 0x38, 0xad, 0xf8,
	0x6f, 0xdd, 0xe4, 0xbf, 0xbd, 0xe6, 0xff, 0x9a, 0xcd, 0xad, 0x7f, 0xb0, 0xc9, 0x5e, 0x41, 0x7f,
	0x2e, 0xa4, 0xd0, 0xcb, 0x5a, 0xeb, 0xfe, 0x55, 0x0b, 0xb6, 0xfc, 0xd8, 0xf8, 0xdf, 0x5b, 0xb0,
	0xbb, 0xde, 0x23, 0x3b, 0x87, 0xee, 0x65, 0x9d, 0x43, 0x33, 0xae, 0xd7, 0xc1, 0xed, 0x02, 0x0f,
	0xfe, 0x18, 0x66, 0x68, 0x69, 0x25, 0x38, 0xaf, 0x53, 0xa9, 0xc6, 0xb2, 0x31, 0xb8, 0x89, 0x36,
	0xb4, 0x34, 0xf6, 0x11, 0xc0, 0xfc, 0x4e, 0xa6, 0x9a, 0x6c, 0x7f, 0xf4, 0x66, 0x33, 0xf6, 0x55,
	0xc2, 0xe1, 0x0a, 0xd3, 0xff, 0xe9, 0xc0, 0xbd, 0xb5, 0xc2, 0x42, 0xb3, 0x53, 0x70, 0x75, 0x29,
	0x69, 0xa6, 0xf4, 0x62, 0xc3, 0x66, 0x6a, 0x08, 0x4b, 0x60, 0x3f, 0x41, 0x6d, 0x22, 0xfb, 0x53,
	0xa1, 0x64, 0x54, 0xe3, 0x5b, 0xff, 0x85, 0xdf, 0x2b, 0xa9, 0xd3, 0x2b, 0x68, 0x9d, 0xf1, 0x63,
	0xd8, 0xc9, 0xa9, 0x7a, 0x42, 0x51, 0xac, 0x0a, 0x69, 0xaa, 0xa1, 0xb9, 0xe1, 0xdd, 0xe6, 0x72,
	0x5c, 0xde, 0xf9, 0xdf, 0x1c, 0xe8, 0x4d, 0x14, 0x6f, 0xda, 0x3d, 0x01, 0x37, 0x5b, 0xa2, 0xae,
	0xdb, 0xdd, 0x1d, 0x3d, 0xbb, 0xad, 0x9f, 0x89, 0xe2, 0x93, 0x52, 0x17, 0xd6, 0xf2, 0x9b, 0x77,
	0x78, 0x65, 0x6b, 0xda, 0xab, 0x5b, 0xf3, 0xf4, 0x13, 0x6c, 0x5b, 0x08, 0xeb, 0x43, 0x77, 0x42,
	0x92, 0x0b, 0xb9, 0xf0, 0xee, 0x94, 0x87, 0xe6, 0x2d, 0x78, 0x0e, 0xdb, 0x81, 0xde, 0x59, 0x11,
	0xc7, 0x44, 0x9c, 0xb8, 0xd7, 0x62, 0x00, 0x9d, 0x13, 0x14, 0x09, 0x71, 0xaf, 0x5d, 0xd6, 0x7d,
	0x90, 0x17, 0x52, 0x5d, 0x4a, 0x6f, 0xab, 0xac, 0x1b, 0xab, 0x34, 0x4b, 0xc8, 0x10, 0xf7, 0x5c,
	0xb6, 0x07, 0xde, 0x38, 0x47, 0xbd, 0x3c, 0x55, 0x2a, 0x7b, 0x8b, 0xf1, 0xc5, 0xfb, 0xf9, 0xdc,
	0xeb, 0xcc, 0x3a, 0xd5, 0xda, 0x3c, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x26, 0x56, 0xfe, 0xc2,
	0x15, 0x05, 0x00, 0x00,
}
