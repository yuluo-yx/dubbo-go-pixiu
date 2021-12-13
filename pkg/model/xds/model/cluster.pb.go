// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// ApiConfigSource
type ApiConfigSource struct {
	ApitypeStr           string   `protobuf:"bytes,1,opt,name=apitype_str,json=apitypeStr,proto3" json:"apitype_str,omitempty"`
	ClusterName          []string `protobuf:"bytes,2,rep,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiConfigSource) Reset()         { *m = ApiConfigSource{} }
func (m *ApiConfigSource) String() string { return proto.CompactTextString(m) }
func (*ApiConfigSource) ProtoMessage()    {}
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{0}
}

func (m *ApiConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiConfigSource.Unmarshal(m, b)
}
func (m *ApiConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiConfigSource.Marshal(b, m, deterministic)
}
func (m *ApiConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiConfigSource.Merge(m, src)
}
func (m *ApiConfigSource) XXX_Size() int {
	return xxx_messageInfo_ApiConfigSource.Size(m)
}
func (m *ApiConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ApiConfigSource proto.InternalMessageInfo

func (m *ApiConfigSource) GetApitypeStr() string {
	if m != nil {
		return m.ApitypeStr
	}
	return ""
}

func (m *ApiConfigSource) GetClusterName() []string {
	if m != nil {
		return m.ClusterName
	}
	return nil
}

// ConfigSource
type ConfigSource struct {
	Path                 string           `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	ApiConfigSource      *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3" json:"api_config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ConfigSource) Reset()         { *m = ConfigSource{} }
func (m *ConfigSource) String() string { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()    {}
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{1}
}

func (m *ConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSource.Unmarshal(m, b)
}
func (m *ConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSource.Marshal(b, m, deterministic)
}
func (m *ConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSource.Merge(m, src)
}
func (m *ConfigSource) XXX_Size() int {
	return xxx_messageInfo_ConfigSource.Size(m)
}
func (m *ConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSource proto.InternalMessageInfo

func (m *ConfigSource) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if m != nil {
		return m.ApiConfigSource
	}
	return nil
}

// EdsClusterConfig
type EdsClusterConfig struct {
	EdsConfig            *ConfigSource `protobuf:"bytes,1,opt,name=eds_config,json=edsConfig,proto3" json:"eds_config,omitempty"`
	ServiceName          string        `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EdsClusterConfig) Reset()         { *m = EdsClusterConfig{} }
func (m *EdsClusterConfig) String() string { return proto.CompactTextString(m) }
func (*EdsClusterConfig) ProtoMessage()    {}
func (*EdsClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{2}
}

func (m *EdsClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdsClusterConfig.Unmarshal(m, b)
}
func (m *EdsClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdsClusterConfig.Marshal(b, m, deterministic)
}
func (m *EdsClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdsClusterConfig.Merge(m, src)
}
func (m *EdsClusterConfig) XXX_Size() int {
	return xxx_messageInfo_EdsClusterConfig.Size(m)
}
func (m *EdsClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EdsClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EdsClusterConfig proto.InternalMessageInfo

func (m *EdsClusterConfig) GetEdsConfig() *ConfigSource {
	if m != nil {
		return m.EdsConfig
	}
	return nil
}

func (m *EdsClusterConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

// Cluster a single upstream cluster
type Cluster struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TypeStr              string            `protobuf:"bytes,2,opt,name=type_str,json=typeStr,proto3" json:"type_str,omitempty"`
	Type                 int32             `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	EdsClusterConfig     *EdsClusterConfig `protobuf:"bytes,4,opt,name=eds_cluster_config,json=edsClusterConfig,proto3" json:"eds_cluster_config,omitempty"`
	LbStr                string            `protobuf:"bytes,5,opt,name=lb_str,json=lbStr,proto3" json:"lb_str,omitempty"`
	Lb                   int32             `protobuf:"varint,6,opt,name=lb,proto3" json:"lb,omitempty"`
	HealthChecks         []*HealthCheck    `protobuf:"bytes,7,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	Endpoints            *Endpoint         `protobuf:"bytes,8,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{3}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetTypeStr() string {
	if m != nil {
		return m.TypeStr
	}
	return ""
}

func (m *Cluster) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Cluster) GetEdsClusterConfig() *EdsClusterConfig {
	if m != nil {
		return m.EdsClusterConfig
	}
	return nil
}

func (m *Cluster) GetLbStr() string {
	if m != nil {
		return m.LbStr
	}
	return ""
}

func (m *Cluster) GetLb() int32 {
	if m != nil {
		return m.Lb
	}
	return 0
}

func (m *Cluster) GetHealthChecks() []*HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *Cluster) GetEndpoints() *Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Endpoint
type Endpoint struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              *SocketAddress    `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{4}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetAddress() *SocketAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type HttpHealthCheck struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	UseHttp2             bool     `protobuf:"varint,3,opt,name=useHttp2,proto3" json:"useHttp2,omitempty"`
	ExpectedStatuses     int64    `protobuf:"varint,4,opt,name=expectedStatuses,proto3" json:"expectedStatuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpHealthCheck) Reset()         { *m = HttpHealthCheck{} }
func (m *HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HttpHealthCheck) ProtoMessage()    {}
func (*HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{5}
}

func (m *HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpHealthCheck.Merge(m, src)
}
func (m *HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HttpHealthCheck.Size(m)
}
func (m *HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HttpHealthCheck proto.InternalMessageInfo

func (m *HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HttpHealthCheck) GetUseHttp2() bool {
	if m != nil {
		return m.UseHttp2
	}
	return false
}

func (m *HttpHealthCheck) GetExpectedStatuses() int64 {
	if m != nil {
		return m.ExpectedStatuses
	}
	return 0
}

type GrpcHealthCheck struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcHealthCheck) Reset()         { *m = GrpcHealthCheck{} }
func (m *GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*GrpcHealthCheck) ProtoMessage()    {}
func (*GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{6}
}

func (m *GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcHealthCheck.Merge(m, src)
}
func (m *GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_GrpcHealthCheck.Size(m)
}
func (m *GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcHealthCheck proto.InternalMessageInfo

func (m *GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type CustomHealthCheck struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config               *structpb.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomHealthCheck) Reset()         { *m = CustomHealthCheck{} }
func (m *CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*CustomHealthCheck) ProtoMessage()    {}
func (*CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{7}
}

func (m *CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHealthCheck.Unmarshal(m, b)
}
func (m *CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHealthCheck.Merge(m, src)
}
func (m *CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_CustomHealthCheck.Size(m)
}
func (m *CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHealthCheck proto.InternalMessageInfo

func (m *CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomHealthCheck) GetConfig() *structpb.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// HealthCheck
type HealthCheck struct {
	// Types that are valid to be assigned to Checker:
	//	*HealthCheck_HttpChecker
	//	*HealthCheck_GrpcChecker
	//	*HealthCheck_CustomChecker
	Checker              isHealthCheck_Checker `protobuf_oneof:"checker"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{8}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

type isHealthCheck_Checker interface {
	isHealthCheck_Checker()
}

type HealthCheck_HttpChecker struct {
	HttpChecker *HttpHealthCheck `protobuf:"bytes,1,opt,name=httpChecker,proto3,oneof"`
}

type HealthCheck_GrpcChecker struct {
	GrpcChecker *GrpcHealthCheck `protobuf:"bytes,2,opt,name=grpcChecker,proto3,oneof"`
}

type HealthCheck_CustomChecker struct {
	CustomChecker *CustomHealthCheck `protobuf:"bytes,3,opt,name=customChecker,proto3,oneof"`
}

func (*HealthCheck_HttpChecker) isHealthCheck_Checker() {}

func (*HealthCheck_GrpcChecker) isHealthCheck_Checker() {}

func (*HealthCheck_CustomChecker) isHealthCheck_Checker() {}

func (m *HealthCheck) GetChecker() isHealthCheck_Checker {
	if m != nil {
		return m.Checker
	}
	return nil
}

func (m *HealthCheck) GetHttpChecker() *HttpHealthCheck {
	if x, ok := m.GetChecker().(*HealthCheck_HttpChecker); ok {
		return x.HttpChecker
	}
	return nil
}

func (m *HealthCheck) GetGrpcChecker() *GrpcHealthCheck {
	if x, ok := m.GetChecker().(*HealthCheck_GrpcChecker); ok {
		return x.GrpcChecker
	}
	return nil
}

func (m *HealthCheck) GetCustomChecker() *CustomHealthCheck {
	if x, ok := m.GetChecker().(*HealthCheck_CustomChecker); ok {
		return x.CustomChecker
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpChecker)(nil),
		(*HealthCheck_GrpcChecker)(nil),
		(*HealthCheck_CustomChecker)(nil),
	}
}

func init() {
	proto.RegisterType((*ApiConfigSource)(nil), "pixiu.api.v1.ApiConfigSource")
	proto.RegisterType((*ConfigSource)(nil), "pixiu.api.v1.ConfigSource")
	proto.RegisterType((*EdsClusterConfig)(nil), "pixiu.api.v1.EdsClusterConfig")
	proto.RegisterType((*Cluster)(nil), "pixiu.api.v1.Cluster")
	proto.RegisterType((*Endpoint)(nil), "pixiu.api.v1.Endpoint")
	proto.RegisterMapType((map[string]string)(nil), "pixiu.api.v1.Endpoint.MetadataEntry")
	proto.RegisterType((*HttpHealthCheck)(nil), "pixiu.api.v1.HttpHealthCheck")
	proto.RegisterType((*GrpcHealthCheck)(nil), "pixiu.api.v1.GrpcHealthCheck")
	proto.RegisterType((*CustomHealthCheck)(nil), "pixiu.api.v1.CustomHealthCheck")
	proto.RegisterType((*HealthCheck)(nil), "pixiu.api.v1.HealthCheck")
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor_3cfb3b8ec240c376) }

var fileDescriptor_3cfb3b8ec240c376 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xae, 0x9d, 0xf3, 0x38, 0xf9, 0x93, 0xae, 0x7e, 0xc0, 0x0d, 0x85, 0xa6, 0x16, 0x17, 0x11,
	0x17, 0xae, 0x08, 0x20, 0x71, 0x90, 0x10, 0x25, 0xaa, 0x1a, 0x24, 0x40, 0x62, 0x23, 0x24, 0xee,
	0xa2, 0x8d, 0xbd, 0x8d, 0xad, 0x3a, 0xf1, 0xca, 0xbb, 0xae, 0x9a, 0x4b, 0x9e, 0x8a, 0x17, 0x42,
	0xe2, 0x35, 0xd0, 0xae, 0xd7, 0x89, 0xed, 0xf6, 0x6e, 0x76, 0x76, 0x0e, 0xdf, 0xf7, 0x69, 0x66,
	0xa0, 0xe7, 0x45, 0x29, 0x17, 0x34, 0x71, 0x59, 0x12, 0x8b, 0x18, 0x75, 0x59, 0x78, 0x1b, 0xa6,
	0x2e, 0x61, 0xa1, 0x7b, 0xf3, 0x62, 0xd8, 0x23, 0xbe, 0x9f, 0x50, 0xce, 0xb3, 0xcf, 0xe1, 0xf1,
	0x2a, 0x8e, 0x57, 0x11, 0x3d, 0x53, 0xaf, 0x65, 0x7a, 0x75, 0xc6, 0x45, 0x92, 0x7a, 0x22, 0xfb,
	0x75, 0x7e, 0x40, 0xff, 0x9c, 0x85, 0xd3, 0x78, 0x73, 0x15, 0xae, 0xe6, 0x71, 0x9a, 0x78, 0x14,
	0x9d, 0x80, 0x45, 0x58, 0x28, 0xb6, 0x8c, 0x2e, 0xb8, 0x48, 0x6c, 0x63, 0x64, 0x8c, 0x3b, 0x18,
	0xb4, 0x6b, 0x2e, 0x12, 0x74, 0x0a, 0x5d, 0xdd, 0x7f, 0xb1, 0x21, 0x6b, 0x6a, 0x9b, 0xa3, 0xda,
	0xb8, 0x83, 0x2d, 0xed, 0xfb, 0x46, 0xd6, 0xd4, 0x59, 0x43, 0xb7, 0x54, 0x13, 0x41, 0x9d, 0x11,
	0x11, 0xe8, 0x62, 0xca, 0x46, 0x9f, 0xe1, 0x90, 0xb0, 0x70, 0xe1, 0xa9, 0xb8, 0x05, 0x57, 0x81,
	0xb6, 0x39, 0x32, 0xc6, 0xd6, 0xe4, 0x89, 0x5b, 0x64, 0xe4, 0x56, 0x10, 0xe2, 0x3e, 0x29, 0x3b,
	0x1c, 0x06, 0x83, 0x0b, 0x9f, 0x4f, 0x33, 0x00, 0xd9, 0x0f, 0x7a, 0x0b, 0x40, 0x7d, 0xae, 0xcb,
	0xab, 0xc6, 0xd6, 0x64, 0x58, 0xae, 0x5b, 0x2a, 0xda, 0xa1, 0x3e, 0xd7, 0xa9, 0xa7, 0xd0, 0xe5,
	0x34, 0xb9, 0x09, 0x3d, 0x9a, 0x13, 0x94, 0xa8, 0x2d, 0xed, 0x53, 0x04, 0x7f, 0x9b, 0xd0, 0xd2,
	0xfd, 0x24, 0x39, 0x15, 0xa6, 0xc9, 0x49, 0x1b, 0x1d, 0x41, 0x7b, 0xa7, 0x60, 0x96, 0xde, 0xca,
	0xe5, 0x43, 0x50, 0x97, 0xa6, 0x5d, 0x1b, 0x19, 0xe3, 0x06, 0x56, 0x36, 0xfa, 0x02, 0x48, 0x81,
	0xd5, 0xb2, 0x6a, 0xd0, 0x75, 0x05, 0xfa, 0x69, 0x19, 0x74, 0x95, 0x28, 0x1e, 0xd0, 0x2a, 0xf5,
	0x07, 0xd0, 0x8c, 0x96, 0xaa, 0x75, 0x43, 0xb5, 0x6e, 0x44, 0x4b, 0xd9, 0xf8, 0x3f, 0x30, 0xa3,
	0xa5, 0xdd, 0x54, 0x6d, 0xcd, 0x68, 0x89, 0x3e, 0x40, 0x2f, 0xa0, 0x24, 0x12, 0xc1, 0xc2, 0x0b,
	0xa8, 0x77, 0xcd, 0xed, 0xd6, 0xa8, 0x36, 0xb6, 0x26, 0x47, 0xe5, 0x7e, 0x33, 0x15, 0x32, 0x95,
	0x11, 0xb8, 0x1b, 0xec, 0x1f, 0x1c, 0xbd, 0x82, 0x0e, 0xdd, 0xf8, 0x2c, 0x0e, 0x37, 0x82, 0xdb,
	0x6d, 0x85, 0xf5, 0x61, 0x05, 0xab, 0xfe, 0xc6, 0xfb, 0x40, 0xe7, 0x8f, 0x01, 0xed, 0xdc, 0x2f,
	0x21, 0x85, 0xbe, 0x16, 0xce, 0x0c, 0xfd, 0x9d, 0x94, 0x66, 0x41, 0xca, 0xd7, 0xd0, 0xd2, 0x13,
	0xad, 0x24, 0xb3, 0x26, 0x8f, 0xcb, 0x4d, 0xe6, 0xb1, 0x77, 0x4d, 0xc5, 0x79, 0x16, 0x82, 0xf3,
	0x58, 0xf4, 0x11, 0xda, 0x6b, 0x2a, 0x88, 0x4f, 0x04, 0xb1, 0xeb, 0x8a, 0xd8, 0xb3, 0xfb, 0xc1,
	0xb9, 0x5f, 0x75, 0xd8, 0xc5, 0x46, 0x24, 0x5b, 0xbc, 0xcb, 0x1a, 0xbe, 0x87, 0x5e, 0xe9, 0x0b,
	0x0d, 0xa0, 0x76, 0x4d, 0xb7, 0x1a, 0xae, 0x34, 0xd1, 0xff, 0xd0, 0xb8, 0x21, 0x51, 0x9a, 0x03,
	0xce, 0x1e, 0xef, 0xcc, 0x37, 0x86, 0xf3, 0xcb, 0x80, 0xfe, 0x4c, 0x08, 0x56, 0x90, 0x4f, 0xb2,
	0x0b, 0x62, 0x2e, 0xf2, 0x41, 0x91, 0xf6, 0x6e, 0x33, 0xcc, 0xc2, 0x66, 0x0c, 0xa1, 0x9d, 0x72,
	0x2a, 0xb3, 0x27, 0x8a, 0x72, 0x1b, 0xef, 0xde, 0xe8, 0x39, 0x0c, 0xe8, 0x2d, 0xa3, 0x9e, 0xa0,
	0xfe, 0x5c, 0x10, 0x91, 0x72, 0xca, 0xd5, 0x9c, 0xd4, 0xf0, 0x1d, 0xbf, 0xf3, 0x1d, 0xfa, 0x97,
	0x09, 0xf3, 0x8a, 0x10, 0x46, 0x50, 0x1c, 0x63, 0x8d, 0xa4, 0xe8, 0x42, 0xc7, 0xd0, 0x21, 0xa9,
	0x08, 0xe2, 0x24, 0x14, 0x5b, 0x8d, 0x6a, 0xef, 0x70, 0x7e, 0xc2, 0xe1, 0x34, 0xe5, 0x22, 0x5e,
	0x57, 0x78, 0xdd, 0x59, 0x80, 0x33, 0x68, 0xea, 0x29, 0xce, 0x56, 0xfa, 0x91, 0x9b, 0xdd, 0x21,
	0x37, 0xbf, 0x43, 0xee, 0x5c, 0xdd, 0x21, 0xac, 0xc3, 0x9c, 0xbf, 0x06, 0x58, 0xc5, 0xa2, 0xe7,
	0x60, 0x05, 0x42, 0x30, 0xf5, 0xa0, 0x89, 0x5e, 0xe0, 0xca, 0x61, 0xa8, 0x08, 0x3c, 0x3b, 0xc0,
	0xc5, 0x1c, 0x59, 0x62, 0x95, 0x30, 0x2f, 0x2f, 0x71, 0xef, 0x6d, 0xa9, 0x08, 0x24, 0x4b, 0x14,
	0x72, 0xd0, 0x25, 0xf4, 0x3c, 0xc5, 0x37, 0x2f, 0x92, 0x8d, 0xe0, 0x49, 0xe5, 0x90, 0x54, 0x25,
	0x99, 0x1d, 0xe0, 0x72, 0xde, 0xa7, 0x0e, 0xb4, 0xbc, 0xcc, 0x5c, 0x36, 0x95, 0x04, 0x2f, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x32, 0x99, 0x2c, 0xc6, 0x05, 0x00, 0x00,
}
