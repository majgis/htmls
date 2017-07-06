// Code generated by protoc-gen-go. DO NOT EDIT.
// source: htmls.proto

/*
Package htmls is a generated protocol buffer package.

It is generated from these files:
	htmls.proto

It has these top-level messages:
	QueryValues
	HTMLRequest
	Resource
	HTMLResponse
*/
package htmls

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

type ResourceType int32

const (
	ResourceType_JS  ResourceType = 0
	ResourceType_CSS ResourceType = 1
)

var ResourceType_name = map[int32]string{
	0: "JS",
	1: "CSS",
}
var ResourceType_value = map[string]int32{
	"JS":  0,
	"CSS": 1,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}
func (ResourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type QueryValues struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *QueryValues) Reset()                    { *m = QueryValues{} }
func (m *QueryValues) String() string            { return proto.CompactTextString(m) }
func (*QueryValues) ProtoMessage()               {}
func (*QueryValues) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryValues) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type HTMLRequest struct {
	Headers map[string]string       `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Query   map[string]*QueryValues `protobuf:"bytes,2,rep,name=query" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HTMLRequest) Reset()                    { *m = HTMLRequest{} }
func (m *HTMLRequest) String() string            { return proto.CompactTextString(m) }
func (*HTMLRequest) ProtoMessage()               {}
func (*HTMLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HTMLRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTMLRequest) GetQuery() map[string]*QueryValues {
	if m != nil {
		return m.Query
	}
	return nil
}

type Resource struct {
	Hash         string       `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	HashType     string       `protobuf:"bytes,2,opt,name=hashType" json:"hashType,omitempty"`
	ResourceType ResourceType `protobuf:"varint,3,opt,name=resourceType,enum=htmls.ResourceType" json:"resourceType,omitempty"`
	ResourcePath string       `protobuf:"bytes,4,opt,name=resourcePath" json:"resourcePath,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resource) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Resource) GetHashType() string {
	if m != nil {
		return m.HashType
	}
	return ""
}

func (m *Resource) GetResourceType() ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return ResourceType_JS
}

func (m *Resource) GetResourcePath() string {
	if m != nil {
		return m.ResourcePath
	}
	return ""
}

type HTMLResponse struct {
	Resources []*Resource `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
	HTMLChunk string      `protobuf:"bytes,2,opt,name=HTMLChunk" json:"HTMLChunk,omitempty"`
}

func (m *HTMLResponse) Reset()                    { *m = HTMLResponse{} }
func (m *HTMLResponse) String() string            { return proto.CompactTextString(m) }
func (*HTMLResponse) ProtoMessage()               {}
func (*HTMLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HTMLResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *HTMLResponse) GetHTMLChunk() string {
	if m != nil {
		return m.HTMLChunk
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryValues)(nil), "htmls.QueryValues")
	proto.RegisterType((*HTMLRequest)(nil), "htmls.HTMLRequest")
	proto.RegisterType((*Resource)(nil), "htmls.Resource")
	proto.RegisterType((*HTMLResponse)(nil), "htmls.HTMLResponse")
	proto.RegisterEnum("htmls.ResourceType", ResourceType_name, ResourceType_value)
}

func init() { proto.RegisterFile("htmls.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x7d, 0x93, 0xf4, 0x2b, 0x37, 0xe1, 0xbd, 0x72, 0xfb, 0x16, 0x21, 0x28, 0x2d, 0x71, 0x13,
	0x04, 0x8b, 0xa4, 0x82, 0xda, 0x6d, 0x11, 0x8b, 0x54, 0xd0, 0x49, 0x71, 0xe3, 0x2a, 0xd6, 0x81,
	0x48, 0x6b, 0xd3, 0x66, 0x92, 0x42, 0x7e, 0x81, 0x3f, 0xc1, 0xbf, 0x2b, 0x99, 0x49, 0xcc, 0x54,
	0xbb, 0x9a, 0x7b, 0xcf, 0x3d, 0xe7, 0xdc, 0xc3, 0xcc, 0x80, 0x19, 0xa5, 0xef, 0x2b, 0x3e, 0xdc,
	0x24, 0x71, 0x1a, 0x63, 0x53, 0x34, 0xee, 0x09, 0x98, 0x8f, 0x19, 0x4b, 0xf2, 0xa7, 0x70, 0x95,
	0x31, 0x8e, 0xff, 0xa1, 0xb9, 0x2b, 0x2a, 0x9b, 0x0c, 0x74, 0xcf, 0xa0, 0xb2, 0x71, 0x3f, 0x34,
	0x30, 0xa7, 0xf3, 0xfb, 0x19, 0x65, 0xdb, 0x8c, 0xf1, 0x14, 0xaf, 0xa1, 0x1d, 0xb1, 0xf0, 0x95,
	0x25, 0x5c, 0xf0, 0x4c, 0xbf, 0x3f, 0x94, 0xd6, 0x0a, 0x69, 0x38, 0x95, 0x8c, 0x9b, 0x75, 0x9a,
	0xe4, 0xb4, 0xe2, 0xe3, 0x08, 0x9a, 0xdb, 0x62, 0x9f, 0xad, 0x09, 0xe1, 0xf1, 0x01, 0xa1, 0xc8,
	0x23, 0x65, 0x92, 0xeb, 0x8c, 0xc1, 0x52, 0xdd, 0xb0, 0x0b, 0xfa, 0x92, 0xe5, 0x36, 0x19, 0x10,
	0xcf, 0xa0, 0x45, 0x59, 0xe7, 0xd6, 0x04, 0x26, 0x9b, 0xb1, 0x76, 0x45, 0x9c, 0x19, 0x40, 0x6d,
	0x78, 0x40, 0xe9, 0xa9, 0x4a, 0xd3, 0xc7, 0x32, 0x90, 0x72, 0x29, 0x8a, 0x9b, 0xfb, 0x49, 0xa0,
	0x43, 0x19, 0x8f, 0xb3, 0x64, 0xc1, 0x10, 0xa1, 0x11, 0x85, 0x3c, 0x2a, 0xdd, 0x44, 0x8d, 0x0e,
	0x74, 0x8a, 0x73, 0x9e, 0x6f, 0xaa, 0x2c, 0xdf, 0x3d, 0x5e, 0x82, 0x95, 0x94, 0x5a, 0x31, 0xd7,
	0x07, 0xc4, 0xfb, 0xeb, 0xf7, 0xca, 0x8d, 0x54, 0x19, 0xd1, 0x3d, 0x22, 0xba, 0xb5, 0xf0, 0x21,
	0x4c, 0x23, 0xbb, 0x21, 0x8c, 0xf7, 0x30, 0xf7, 0x19, 0x2c, 0x79, 0x89, 0x7c, 0x13, 0xaf, 0x39,
	0xc3, 0x33, 0x30, 0xaa, 0x79, 0xf5, 0x4a, 0xff, 0x7e, 0x6c, 0xa2, 0x35, 0x03, 0x8f, 0xc0, 0x28,
	0xe4, 0x93, 0x28, 0x5b, 0x2f, 0xcb, 0xe0, 0x35, 0x70, 0xda, 0x07, 0x4b, 0x8d, 0x87, 0x2d, 0xd0,
	0xee, 0x82, 0xee, 0x1f, 0x6c, 0x83, 0x3e, 0x09, 0x82, 0x2e, 0xf1, 0x27, 0xf2, 0x83, 0x04, 0x2c,
	0xd9, 0xbd, 0x2d, 0x18, 0x5e, 0x40, 0xfb, 0x96, 0xa5, 0x05, 0x82, 0xf8, 0xfb, 0x85, 0x9d, 0xde,
	0x1e, 0x26, 0x03, 0x9f, 0x93, 0x97, 0x96, 0xf8, 0x99, 0xa3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0xda, 0x7d, 0x20, 0xa8, 0x02, 0x00, 0x00,
}
