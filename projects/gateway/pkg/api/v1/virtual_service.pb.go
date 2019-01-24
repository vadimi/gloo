// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//@solo-kit:resource.short_name=vs
//@solo-kit:resource.plural_name=virtual_services
//@solo-kit:resource.resource_groups=api.gateway.solo.io
//A virtual service describes the set of routes to match for a set of domains.
//Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
type VirtualService struct {
	VirtualHost *v1.VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost,proto3" json:"virtual_host,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic for this set of routes
	SslConfig *v1.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VirtualService) Reset()         { *m = VirtualService{} }
func (m *VirtualService) String() string { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()    {}
func (*VirtualService) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{0}
}
func (m *VirtualService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualService.Unmarshal(m, b)
}
func (m *VirtualService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualService.Marshal(b, m, deterministic)
}
func (m *VirtualService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualService.Merge(m, src)
}
func (m *VirtualService) XXX_Size() int {
	return xxx_messageInfo_VirtualService.Size(m)
}
func (m *VirtualService) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualService.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualService proto.InternalMessageInfo

func (m *VirtualService) GetVirtualHost() *v1.VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *v1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *VirtualService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto", fileDescriptor_93fa9472926a2049)
}

var fileDescriptor_93fa9472926a2049 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0x3a, 0x31,
	0x10, 0xc6, 0xff, 0x90, 0x7f, 0x50, 0x8b, 0xd1, 0xb8, 0x21, 0xba, 0x70, 0x10, 0xc3, 0xc9, 0x8b,
	0x6d, 0x90, 0xc4, 0xa0, 0xf1, 0x84, 0x31, 0x7a, 0xf1, 0xb2, 0x24, 0x1e, 0xbc, 0x90, 0xb2, 0x94,
	0x52, 0x59, 0x98, 0x4d, 0x3b, 0xbb, 0xca, 0x23, 0xf8, 0x26, 0x3e, 0x8a, 0x4f, 0xc1, 0xc1, 0x47,
	0xf0, 0x09, 0xcc, 0x76, 0xbb, 0x12, 0x12, 0x13, 0x39, 0x75, 0x26, 0xdf, 0xfc, 0xbe, 0xc9, 0x37,
	0x29, 0xb9, 0x95, 0x0a, 0x27, 0xc9, 0x90, 0x86, 0x30, 0x63, 0x06, 0x22, 0x38, 0x53, 0xc0, 0x64,
	0x04, 0xc0, 0x62, 0x0d, 0xcf, 0x22, 0x44, 0xc3, 0x24, 0x47, 0xf1, 0xc2, 0x17, 0x8c, 0xc7, 0x8a,
	0xa5, 0x6d, 0x96, 0x2a, 0x8d, 0x09, 0x8f, 0x06, 0x46, 0xe8, 0x54, 0x85, 0x82, 0xc6, 0x1a, 0x10,
	0xbc, 0x7d, 0x37, 0x45, 0x33, 0x0f, 0xaa, 0xa0, 0x51, 0x93, 0x20, 0xc1, 0x6a, 0x2c, 0xab, 0xf2,
	0xb1, 0x46, 0xfb, 0x97, 0x6d, 0xf6, 0x9d, 0x2a, 0x2c, 0x16, 0xcc, 0x04, 0xf2, 0x11, 0x47, 0xee,
	0x10, 0xb6, 0x01, 0x62, 0x90, 0x63, 0x62, 0x1c, 0xd0, 0xfd, 0x3b, 0x51, 0xd6, 0x39, 0x34, 0xd6,
	0xf0, 0xba, 0xc8, 0xc9, 0xd6, 0x5b, 0x99, 0xec, 0x3d, 0xe6, 0xf1, 0xfa, 0x79, 0x3a, 0xef, 0x9a,
	0xec, 0x16, 0x81, 0x27, 0x60, 0xd0, 0x2f, 0x9d, 0x94, 0x4e, 0xab, 0xe7, 0x75, 0x9a, 0x59, 0x14,
	0x59, 0xa9, 0x63, 0xee, 0xc1, 0x60, 0x50, 0x4d, 0x57, 0x8d, 0x77, 0x41, 0x88, 0x31, 0xd1, 0x20,
	0x84, 0xf9, 0x58, 0x49, 0xbf, 0x6c, 0xd9, 0xa3, 0x75, 0xb6, 0x6f, 0xa2, 0x1b, 0x2b, 0x07, 0x3b,
	0xa6, 0x28, 0xbd, 0x3b, 0x52, 0xc9, 0x23, 0xf9, 0x15, 0xcb, 0xd4, 0x68, 0x08, 0x5a, 0xac, 0x18,
	0xab, 0xf5, 0xea, 0x1f, 0xcb, 0xe6, 0xbf, 0xaf, 0x65, 0xf3, 0x00, 0x85, 0xc1, 0x91, 0x1a, 0x8f,
	0xaf, 0x5a, 0x4a, 0xce, 0x41, 0x8b, 0x56, 0xe0, 0x70, 0xaf, 0x4b, 0xb6, 0x8b, 0x73, 0xfa, 0x5b,
	0xd6, 0xea, 0x70, 0xdd, 0xea, 0xc1, 0xa9, 0xbd, 0xff, 0x99, 0x59, 0xf0, 0x33, 0xdd, 0xbb, 0x7c,
	0xff, 0x3c, 0x2e, 0x3d, 0x75, 0x36, 0xfe, 0x1d, 0xf1, 0x54, 0xba, 0x93, 0x0e, 0x2b, 0xf6, 0x9a,
	0x9d, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x3e, 0x4d, 0x63, 0x5b, 0x02, 0x00, 0x00,
}

func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
