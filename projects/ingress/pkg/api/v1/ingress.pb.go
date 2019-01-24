// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//@solo-kit:resource.short_name=ig
//@solo-kit:resource.plural_name=ingresses
//
//A simple wrapper for a Kubernetes Ingress Object.
type Ingress struct {
	// a raw byte representation of the kubernetes ingress this resource wraps
	KubeIngressSpec *types.Any `protobuf:"bytes,1,opt,name=kube_ingress_spec,json=kubeIngressSpec,proto3" json:"kube_ingress_spec,omitempty"`
	// a raw byte representation of the ingress status of the kubernetes ingress object
	KubeIngressStatus *types.Any `protobuf:"bytes,2,opt,name=kube_ingress_status,json=kubeIngressStatus,proto3" json:"kube_ingress_status,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ingress) Reset()         { *m = Ingress{} }
func (m *Ingress) String() string { return proto.CompactTextString(m) }
func (*Ingress) ProtoMessage()    {}
func (*Ingress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e3857ca3a6e6b32, []int{0}
}
func (m *Ingress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingress.Unmarshal(m, b)
}
func (m *Ingress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingress.Marshal(b, m, deterministic)
}
func (m *Ingress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingress.Merge(m, src)
}
func (m *Ingress) XXX_Size() int {
	return xxx_messageInfo_Ingress.Size(m)
}
func (m *Ingress) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingress.DiscardUnknown(m)
}

var xxx_messageInfo_Ingress proto.InternalMessageInfo

func (m *Ingress) GetKubeIngressSpec() *types.Any {
	if m != nil {
		return m.KubeIngressSpec
	}
	return nil
}

func (m *Ingress) GetKubeIngressStatus() *types.Any {
	if m != nil {
		return m.KubeIngressStatus
	}
	return nil
}

func (m *Ingress) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Ingress)(nil), "ingress.solo.io.Ingress")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto", fileDescriptor_7e3857ca3a6e6b32)
}

var fileDescriptor_7e3857ca3a6e6b32 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0xcf, 0xcc, 0x4b,
	0x2f, 0x4a, 0x2d, 0x2e, 0xd6, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0x84, 0x71, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x61, 0x5c, 0x90, 0x5e, 0xbd, 0xcc, 0x7c, 0x29, 0x91, 0xf4, 0xfc,
	0xf4, 0x7c, 0xb0, 0x9c, 0x3e, 0x88, 0x05, 0x51, 0x26, 0x25, 0x99, 0x9e, 0x9f, 0x9f, 0x9e, 0x93,
	0xaa, 0x0f, 0xe6, 0x25, 0x95, 0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42, 0xa5, 0x0c, 0xb1, 0x38, 0x00,
	0x4c, 0x67, 0x67, 0x96, 0xc0, 0xec, 0xcc, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x84, 0x6a,
	0xd1, 0x27, 0x42, 0x4b, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31, 0x09, 0x76, 0xc0, 0xf8, 0x10, 0x2d,
	0x4a, 0x17, 0x19, 0xb9, 0xd8, 0x3d, 0x21, 0x7e, 0x13, 0x72, 0xe0, 0x12, 0xcc, 0x2e, 0x4d, 0x4a,
	0x8d, 0x87, 0xfa, 0x35, 0xbe, 0xb8, 0x20, 0x35, 0x59, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48,
	0x44, 0x0f, 0xe2, 0x33, 0x3d, 0x98, 0xcf, 0xf4, 0x1c, 0xf3, 0x2a, 0x83, 0xf8, 0x41, 0xca, 0xa1,
	0xba, 0x83, 0x0b, 0x52, 0x93, 0x85, 0xbc, 0xb8, 0x84, 0x51, 0x4d, 0x00, 0xbb, 0x4e, 0x82, 0x09,
	0xb7, 0x19, 0x4e, 0x2c, 0x0d, 0x1f, 0x59, 0x18, 0x83, 0x04, 0x91, 0x4d, 0x02, 0x6b, 0x12, 0xb2,
	0xe0, 0xe2, 0x80, 0x85, 0x87, 0x04, 0x3b, 0xd8, 0x00, 0x31, 0xbd, 0xe4, 0xfc, 0xa2, 0x54, 0x58,
	0x14, 0xe8, 0xf9, 0x42, 0x65, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0xab, 0x76, 0xb2,
	0x5c, 0xf1, 0x48, 0x8e, 0x31, 0xca, 0x98, 0xe8, 0x18, 0x2f, 0xc8, 0x4e, 0x87, 0x86, 0x4e, 0x12,
	0x1b, 0xd8, 0x6d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x5e, 0x11, 0x49, 0x2f, 0x02,
	0x00, 0x00,
}

func (this *Ingress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ingress)
	if !ok {
		that2, ok := that.(Ingress)
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
	if !this.KubeIngressSpec.Equal(that1.KubeIngressSpec) {
		return false
	}
	if !this.KubeIngressStatus.Equal(that1.KubeIngressStatus) {
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
