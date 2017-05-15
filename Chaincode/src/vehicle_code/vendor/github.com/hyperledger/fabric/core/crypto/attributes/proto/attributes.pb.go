// Code generated by protoc-gen-go.
// source: attributes.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	attributes.proto

It has these top-level messages:
	AttributesMetadataEntry
	AttributesMetadata
*/
package protos

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

// AttributesMetadataEntry is an entry within the metadata that store an attribute name with its respective key.
type AttributesMetadataEntry struct {
	AttributeName string `protobuf:"bytes,1,opt,name=AttributeName,json=attributeName" json:"AttributeName,omitempty"`
	AttributeKey  []byte `protobuf:"bytes,2,opt,name=AttributeKey,json=attributeKey,proto3" json:"AttributeKey,omitempty"`
}

func (m *AttributesMetadataEntry) Reset()                    { *m = AttributesMetadataEntry{} }
func (m *AttributesMetadataEntry) String() string            { return proto.CompactTextString(m) }
func (*AttributesMetadataEntry) ProtoMessage()               {}
func (*AttributesMetadataEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// AttributesMetadata holds both the original metadata bytes and the metadata required to access attributes.
type AttributesMetadata struct {
	// Original metadata bytes
	Metadata []byte `protobuf:"bytes,1,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	// Entries for each attributes considered.
	Entries []*AttributesMetadataEntry `protobuf:"bytes,2,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *AttributesMetadata) Reset()                    { *m = AttributesMetadata{} }
func (m *AttributesMetadata) String() string            { return proto.CompactTextString(m) }
func (*AttributesMetadata) ProtoMessage()               {}
func (*AttributesMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AttributesMetadata) GetEntries() []*AttributesMetadataEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*AttributesMetadataEntry)(nil), "protos.AttributesMetadataEntry")
	proto.RegisterType((*AttributesMetadata)(nil), "protos.AttributesMetadata")
}

func init() { proto.RegisterFile("attributes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0x29, 0x29,
	0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53,
	0xc5, 0x4a, 0xc9, 0x5c, 0xe2, 0x8e, 0x70, 0x39, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0x44,
	0xd7, 0xbc, 0x92, 0xa2, 0x4a, 0x21, 0x15, 0x2e, 0x5e, 0xb8, 0x94, 0x5f, 0x62, 0x6e, 0xaa, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x6f, 0x22, 0xb2, 0xa0, 0x90, 0x12, 0x17, 0x0f, 0x5c, 0x95,
	0x77, 0x6a, 0xa5, 0x04, 0x13, 0x50, 0x11, 0x4f, 0x10, 0x4f, 0x22, 0x92, 0x98, 0x52, 0x36, 0x97,
	0x10, 0xa6, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0x30, 0x36, 0xd8, 0x68, 0x9e, 0x20, 0x8e, 0x5c, 0x98,
	0x9c, 0x25, 0x17, 0x3b, 0xc8, 0x11, 0x99, 0xa9, 0xc5, 0x40, 0x03, 0x99, 0x35, 0xb8, 0x8d, 0xe4,
	0x21, 0xee, 0x2e, 0xd6, 0xc3, 0xe1, 0xda, 0x20, 0xf6, 0x54, 0x88, 0xfa, 0x24, 0x88, 0xcf, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x6e, 0x61, 0x35, 0xf4, 0x00, 0x00, 0x00,
}
