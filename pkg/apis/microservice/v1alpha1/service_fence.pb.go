// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/microservice/v1alpha1/service_fence.proto

package v1alpha1

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

type Destinations_Status int32

const (
	Destinations_ACTIVE Destinations_Status = 0
	Destinations_EXPIRE Destinations_Status = 1
)

var Destinations_Status_name = map[int32]string{
	0: "ACTIVE",
	1: "EXPIRE",
}

var Destinations_Status_value = map[string]int32{
	"ACTIVE": 0,
	"EXPIRE": 1,
}

func (x Destinations_Status) String() string {
	return proto.EnumName(Destinations_Status_name, int32(x))
}

func (Destinations_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{3, 0}
}

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{0}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type ServiceFenceSpec struct {
	// 访问的服务以及其回收策略
	Host                 map[string]*RecyclingStrategy `protobuf:"bytes,1,rep,name=host,proto3" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Enable               bool                          `protobuf:"varint,2,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ServiceFenceSpec) Reset()         { *m = ServiceFenceSpec{} }
func (m *ServiceFenceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceFenceSpec) ProtoMessage()    {}
func (*ServiceFenceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{1}
}

func (m *ServiceFenceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFenceSpec.Unmarshal(m, b)
}
func (m *ServiceFenceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFenceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceFenceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFenceSpec.Merge(m, src)
}
func (m *ServiceFenceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceFenceSpec.Size(m)
}
func (m *ServiceFenceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFenceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFenceSpec proto.InternalMessageInfo

func (m *ServiceFenceSpec) GetHost() map[string]*RecyclingStrategy {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ServiceFenceSpec) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type RecyclingStrategy struct {
	// 永久存在的配置
	Stable *RecyclingStrategy_Stable `protobuf:"bytes,1,opt,name=stable,proto3" json:"stable,omitempty"`
	// 到期失效
	Deadline *RecyclingStrategy_Deadline `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// 当最近一次访问时间大于duration时，将其状态改变为失效
	Auto                 *RecyclingStrategy_Auto `protobuf:"bytes,3,opt,name=auto,proto3" json:"auto,omitempty"`
	RecentlyCalled       *Timestamp              `protobuf:"bytes,4,opt,name=RecentlyCalled,proto3" json:"RecentlyCalled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RecyclingStrategy) Reset()         { *m = RecyclingStrategy{} }
func (m *RecyclingStrategy) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy) ProtoMessage()    {}
func (*RecyclingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{2}
}

func (m *RecyclingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy.Unmarshal(m, b)
}
func (m *RecyclingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy.Merge(m, src)
}
func (m *RecyclingStrategy) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy.Size(m)
}
func (m *RecyclingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy proto.InternalMessageInfo

func (m *RecyclingStrategy) GetStable() *RecyclingStrategy_Stable {
	if m != nil {
		return m.Stable
	}
	return nil
}

func (m *RecyclingStrategy) GetDeadline() *RecyclingStrategy_Deadline {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *RecyclingStrategy) GetAuto() *RecyclingStrategy_Auto {
	if m != nil {
		return m.Auto
	}
	return nil
}

func (m *RecyclingStrategy) GetRecentlyCalled() *Timestamp {
	if m != nil {
		return m.RecentlyCalled
	}
	return nil
}

type RecyclingStrategy_Stable struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecyclingStrategy_Stable) Reset()         { *m = RecyclingStrategy_Stable{} }
func (m *RecyclingStrategy_Stable) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Stable) ProtoMessage()    {}
func (*RecyclingStrategy_Stable) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{2, 0}
}

func (m *RecyclingStrategy_Stable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Stable.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Stable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Stable.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Stable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Stable.Merge(m, src)
}
func (m *RecyclingStrategy_Stable) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Stable.Size(m)
}
func (m *RecyclingStrategy_Stable) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Stable.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Stable proto.InternalMessageInfo

type RecyclingStrategy_Deadline struct {
	Expire               *Timestamp `protobuf:"bytes,1,opt,name=expire,proto3" json:"expire,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RecyclingStrategy_Deadline) Reset()         { *m = RecyclingStrategy_Deadline{} }
func (m *RecyclingStrategy_Deadline) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Deadline) ProtoMessage()    {}
func (*RecyclingStrategy_Deadline) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{2, 1}
}

func (m *RecyclingStrategy_Deadline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Deadline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Deadline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Deadline.Merge(m, src)
}
func (m *RecyclingStrategy_Deadline) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Size(m)
}
func (m *RecyclingStrategy_Deadline) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Deadline.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Deadline proto.InternalMessageInfo

func (m *RecyclingStrategy_Deadline) GetExpire() *Timestamp {
	if m != nil {
		return m.Expire
	}
	return nil
}

type RecyclingStrategy_Auto struct {
	Duration             *Timestamp `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RecyclingStrategy_Auto) Reset()         { *m = RecyclingStrategy_Auto{} }
func (m *RecyclingStrategy_Auto) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Auto) ProtoMessage()    {}
func (*RecyclingStrategy_Auto) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{2, 2}
}

func (m *RecyclingStrategy_Auto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Auto.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Auto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Auto.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Auto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Auto.Merge(m, src)
}
func (m *RecyclingStrategy_Auto) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Auto.Size(m)
}
func (m *RecyclingStrategy_Auto) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Auto.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Auto proto.InternalMessageInfo

func (m *RecyclingStrategy_Auto) GetDuration() *Timestamp {
	if m != nil {
		return m.Duration
	}
	return nil
}

type Destinations struct {
	// 记录最近一次调用，当失效策略为auto时，需定期刷新该值
	RecentlyCalled *Timestamp `protobuf:"bytes,1,opt,name=RecentlyCalled,proto3" json:"RecentlyCalled,omitempty"`
	// domain对应的上游服务
	Hosts                []string            `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Status               Destinations_Status `protobuf:"varint,3,opt,name=status,proto3,enum=netease.microservice.v1alpha1.Destinations_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Destinations) Reset()         { *m = Destinations{} }
func (m *Destinations) String() string { return proto.CompactTextString(m) }
func (*Destinations) ProtoMessage()    {}
func (*Destinations) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{3}
}

func (m *Destinations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destinations.Unmarshal(m, b)
}
func (m *Destinations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destinations.Marshal(b, m, deterministic)
}
func (m *Destinations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destinations.Merge(m, src)
}
func (m *Destinations) XXX_Size() int {
	return xxx_messageInfo_Destinations.Size(m)
}
func (m *Destinations) XXX_DiscardUnknown() {
	xxx_messageInfo_Destinations.DiscardUnknown(m)
}

var xxx_messageInfo_Destinations proto.InternalMessageInfo

func (m *Destinations) GetRecentlyCalled() *Timestamp {
	if m != nil {
		return m.RecentlyCalled
	}
	return nil
}

func (m *Destinations) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Destinations) GetStatus() Destinations_Status {
	if m != nil {
		return m.Status
	}
	return Destinations_ACTIVE
}

type ServiceFenceStatus struct {
	Domains              map[string]*Destinations `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MetricStatus         map[string]string        `protobuf:"bytes,3,rep,name=metricStatus,proto3" json:"metricStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Visitor              map[string]bool          `protobuf:"bytes,2,rep,name=visitor,proto3" json:"visitor,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ServiceFenceStatus) Reset()         { *m = ServiceFenceStatus{} }
func (m *ServiceFenceStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceFenceStatus) ProtoMessage()    {}
func (*ServiceFenceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c73264bd217447, []int{4}
}

func (m *ServiceFenceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFenceStatus.Unmarshal(m, b)
}
func (m *ServiceFenceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFenceStatus.Marshal(b, m, deterministic)
}
func (m *ServiceFenceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFenceStatus.Merge(m, src)
}
func (m *ServiceFenceStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceFenceStatus.Size(m)
}
func (m *ServiceFenceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFenceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFenceStatus proto.InternalMessageInfo

func (m *ServiceFenceStatus) GetDomains() map[string]*Destinations {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *ServiceFenceStatus) GetMetricStatus() map[string]string {
	if m != nil {
		return m.MetricStatus
	}
	return nil
}

func (m *ServiceFenceStatus) GetVisitor() map[string]bool {
	if m != nil {
		return m.Visitor
	}
	return nil
}

func init() {
	proto.RegisterEnum("netease.microservice.v1alpha1.Destinations_Status", Destinations_Status_name, Destinations_Status_value)
	proto.RegisterType((*Timestamp)(nil), "netease.microservice.v1alpha1.Timestamp")
	proto.RegisterType((*ServiceFenceSpec)(nil), "netease.microservice.v1alpha1.ServiceFenceSpec")
	proto.RegisterMapType((map[string]*RecyclingStrategy)(nil), "netease.microservice.v1alpha1.ServiceFenceSpec.HostEntry")
	proto.RegisterType((*RecyclingStrategy)(nil), "netease.microservice.v1alpha1.RecyclingStrategy")
	proto.RegisterType((*RecyclingStrategy_Stable)(nil), "netease.microservice.v1alpha1.RecyclingStrategy.Stable")
	proto.RegisterType((*RecyclingStrategy_Deadline)(nil), "netease.microservice.v1alpha1.RecyclingStrategy.Deadline")
	proto.RegisterType((*RecyclingStrategy_Auto)(nil), "netease.microservice.v1alpha1.RecyclingStrategy.Auto")
	proto.RegisterType((*Destinations)(nil), "netease.microservice.v1alpha1.Destinations")
	proto.RegisterType((*ServiceFenceStatus)(nil), "netease.microservice.v1alpha1.ServiceFenceStatus")
	proto.RegisterMapType((map[string]*Destinations)(nil), "netease.microservice.v1alpha1.ServiceFenceStatus.DomainsEntry")
	proto.RegisterMapType((map[string]string)(nil), "netease.microservice.v1alpha1.ServiceFenceStatus.MetricStatusEntry")
	proto.RegisterMapType((map[string]bool)(nil), "netease.microservice.v1alpha1.ServiceFenceStatus.VisitorEntry")
}

func init() {
	proto.RegisterFile("pkg/apis/microservice/v1alpha1/service_fence.proto", fileDescriptor_39c73264bd217447)
}

var fileDescriptor_39c73264bd217447 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc7, 0xe7, 0x3a, 0x4d, 0x93, 0xd3, 0x52, 0x52, 0x51, 0x46, 0x08, 0x0c, 0x4a, 0xae, 0x02,
	0x03, 0x7b, 0xcd, 0x3e, 0xdb, 0xc1, 0xb6, 0x2c, 0x49, 0x59, 0xc6, 0xca, 0x8a, 0xd2, 0x95, 0xb2,
	0x9b, 0xa1, 0x3a, 0x5a, 0x2a, 0x6a, 0xcb, 0xc6, 0x52, 0xc2, 0xfc, 0x24, 0x7b, 0xc1, 0x3d, 0xc7,
	0x98, 0x24, 0xdb, 0xc1, 0x69, 0x43, 0x33, 0x6f, 0x77, 0x3a, 0xb2, 0xfe, 0xbf, 0xf3, 0xa1, 0x73,
	0x64, 0xe8, 0x46, 0x37, 0x53, 0x97, 0x44, 0x4c, 0xb8, 0x01, 0xf3, 0xe2, 0x50, 0xd0, 0x78, 0xce,
	0x3c, 0xea, 0xce, 0x0f, 0x89, 0x1f, 0x5d, 0x93, 0x43, 0x37, 0xdb, 0xf8, 0xf6, 0x9d, 0x72, 0x8f,
	0x3a, 0x51, 0x1c, 0xca, 0x10, 0x3d, 0xe2, 0x54, 0x52, 0x22, 0xa8, 0x53, 0x94, 0x38, 0xb9, 0xa4,
	0xfd, 0x1a, 0xea, 0xe7, 0x2c, 0xa0, 0x42, 0x92, 0x20, 0x42, 0x4d, 0xd8, 0x12, 0xd4, 0x0b, 0xf9,
	0x44, 0x34, 0xad, 0x03, 0xab, 0x63, 0xe3, 0xdc, 0x44, 0xfb, 0xb0, 0xc9, 0x09, 0x0f, 0x45, 0x73,
	0x43, 0xed, 0x6f, 0xe2, 0xd4, 0x68, 0xff, 0xb2, 0xa0, 0x31, 0x4e, 0x89, 0x27, 0xda, 0xe5, 0x38,
	0xa2, 0x1e, 0x3a, 0x85, 0xca, 0x75, 0x28, 0xa4, 0x22, 0xd8, 0x9d, 0xed, 0xee, 0x91, 0x73, 0xaf,
	0x7f, 0xe7, 0xb6, 0xdc, 0xf9, 0xa0, 0xb4, 0x43, 0x2e, 0xe3, 0x04, 0x1b, 0x0c, 0x7a, 0x08, 0x55,
	0xca, 0xc9, 0x95, 0x4f, 0x8d, 0xeb, 0x1a, 0xce, 0xac, 0x16, 0x83, 0xfa, 0xe2, 0x28, 0x6a, 0x80,
	0x7d, 0x43, 0x13, 0x13, 0x74, 0x1d, 0xeb, 0x25, 0x3a, 0x81, 0xcd, 0x39, 0xf1, 0x67, 0xa9, 0x6a,
	0xbb, 0xfb, 0x64, 0x4d, 0x18, 0x98, 0x7a, 0x89, 0xe7, 0x33, 0x3e, 0x1d, 0xcb, 0x98, 0x48, 0x3a,
	0x4d, 0x70, 0x2a, 0x3f, 0xde, 0x78, 0x65, 0xb5, 0x7f, 0xdb, 0xb0, 0x77, 0xe7, 0x00, 0xfa, 0x0c,
	0x55, 0x55, 0x35, 0x1d, 0x98, 0x65, 0x5c, 0xbc, 0x2c, 0xeb, 0xc2, 0x19, 0x1b, 0x39, 0xce, 0x30,
	0xe8, 0x0b, 0xd4, 0x26, 0x94, 0x4c, 0xd4, 0x91, 0x3c, 0xea, 0xa3, 0xd2, 0xc8, 0x41, 0x06, 0xc0,
	0x0b, 0x14, 0x1a, 0x41, 0x85, 0xcc, 0x64, 0xd8, 0xb4, 0x0d, 0xf2, 0x79, 0x69, 0x64, 0x4f, 0x89,
	0xb1, 0x41, 0xa0, 0x33, 0xd8, 0x55, 0xdf, 0x29, 0x97, 0x7e, 0xd2, 0x27, 0xbe, 0x4f, 0x27, 0xcd,
	0x8a, 0x81, 0x76, 0xd6, 0x40, 0x17, 0x1d, 0x86, 0x6f, 0xe9, 0x5b, 0x35, 0xa8, 0xa6, 0x55, 0x68,
	0x7d, 0x82, 0x5a, 0x1e, 0x3c, 0x7a, 0xa7, 0xee, 0xfc, 0x47, 0xc4, 0xe2, 0xbc, 0xb4, 0x7f, 0xcf,
	0xcf, 0x74, 0x8a, 0x56, 0xd1, 0x71, 0xa3, 0x81, 0xaa, 0xe9, 0x4c, 0x25, 0xc2, 0x42, 0x5e, 0x9a,
	0xb5, 0x50, 0xea, 0x3e, 0xdf, 0x19, 0xa8, 0x5d, 0xc6, 0x8d, 0x2d, 0x56, 0x14, 0xc2, 0xfa, 0xbf,
	0x42, 0xe8, 0x01, 0xd3, 0xed, 0xae, 0x07, 0xcc, 0x56, 0x3d, 0x9c, 0x1a, 0xe8, 0xa3, 0xe9, 0x31,
	0x39, 0x13, 0xe6, 0xf6, 0x76, 0xbb, 0xdd, 0x35, 0xfc, 0x62, 0x90, 0xba, 0xbd, 0x94, 0x12, 0x67,
	0x84, 0xf6, 0x81, 0x29, 0xb5, 0x5a, 0x21, 0x80, 0x6a, 0xaf, 0x7f, 0x3e, 0xba, 0x18, 0x36, 0x1e,
	0xe8, 0xf5, 0xf0, 0xf2, 0x6c, 0x84, 0x87, 0x0d, 0xab, 0xfd, 0xb3, 0x02, 0x68, 0x69, 0x1e, 0xd3,
	0xe3, 0x97, 0xb0, 0x35, 0x09, 0x03, 0xc2, 0xb8, 0xc8, 0x66, 0xfa, 0x4d, 0x99, 0x99, 0x36, 0x0c,
	0x67, 0x90, 0x02, 0xd2, 0xc1, 0xce, 0x71, 0x68, 0x0a, 0x3b, 0x01, 0x95, 0x31, 0xf3, 0xc6, 0x79,
	0x92, 0x1a, 0xdf, 0x2f, 0x8f, 0x3f, 0x2d, 0x50, 0x52, 0x1f, 0x4b, 0x60, 0x9d, 0xc2, 0x9c, 0x09,
	0x26, 0xc3, 0xd8, 0xd4, 0xf7, 0x9f, 0x52, 0xb8, 0x48, 0x01, 0x59, 0x0a, 0x19, 0xae, 0xa5, 0x52,
	0x28, 0xe6, 0xb6, 0xe2, 0x25, 0xea, 0x2d, 0xbf, 0x44, 0x8f, 0x4b, 0x5c, 0x61, 0xe1, 0x11, 0x6a,
	0xbd, 0x85, 0xbd, 0x3b, 0x59, 0xae, 0xf0, 0xb6, 0x5f, 0xf4, 0x56, 0x2f, 0x02, 0x8e, 0x61, 0xa7,
	0x98, 0xc2, 0x3a, 0x6d, 0xad, 0xa0, 0x7d, 0xff, 0xe2, 0xeb, 0xb3, 0x64, 0xc6, 0x17, 0x91, 0x7b,
	0x61, 0xe0, 0x0a, 0x5f, 0xf5, 0xb2, 0x7b, 0xff, 0x0f, 0xe9, 0xaa, 0x6a, 0xfe, 0x41, 0x4f, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xa4, 0x1a, 0xd0, 0xb9, 0x06, 0x00, 0x00,
}
