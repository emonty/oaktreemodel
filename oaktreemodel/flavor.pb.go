// Code generated by protoc-gen-go.
// source: flavor.proto
// DO NOT EDIT!

/*
Package oaktree is a generated protocol buffer package.

It is generated from these files:
	flavor.proto

It has these top-level messages:
	Flavor
	FlavorList
*/
package oaktree

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import oaktree1 "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Flavor struct {
	Location   *oaktree1.Location `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Id         string             `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name       string             `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	IsPublic   bool               `protobuf:"varint,4,opt,name=is_public,json=isPublic" json:"is_public,omitempty"`
	Disabled   bool               `protobuf:"varint,5,opt,name=disabled" json:"disabled,omitempty"`
	Ram        uint32             `protobuf:"varint,6,opt,name=ram" json:"ram,omitempty"`
	Vcpus      uint32             `protobuf:"varint,7,opt,name=vcpus" json:"vcpus,omitempty"`
	Disk       uint64             `protobuf:"varint,8,opt,name=disk" json:"disk,omitempty"`
	Ephemeral  uint32             `protobuf:"varint,9,opt,name=ephemeral" json:"ephemeral,omitempty"`
	Swap       uint64             `protobuf:"varint,10,opt,name=swap" json:"swap,omitempty"`
	RxtxFactor float32            `protobuf:"fixed32,11,opt,name=rxtx_factor,json=rxtxFactor" json:"rxtx_factor,omitempty"`
	Properties map[string]string  `protobuf:"bytes,99,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Flavor) Reset()                    { *m = Flavor{} }
func (m *Flavor) String() string            { return proto.CompactTextString(m) }
func (*Flavor) ProtoMessage()               {}
func (*Flavor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Flavor) GetLocation() *oaktree1.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Flavor) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type FlavorList struct {
	Flavors []*Flavor `protobuf:"bytes,1,rep,name=flavors" json:"flavors,omitempty"`
}

func (m *FlavorList) Reset()                    { *m = FlavorList{} }
func (m *FlavorList) String() string            { return proto.CompactTextString(m) }
func (*FlavorList) ProtoMessage()               {}
func (*FlavorList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FlavorList) GetFlavors() []*Flavor {
	if m != nil {
		return m.Flavors
	}
	return nil
}

func init() {
	proto.RegisterType((*Flavor)(nil), "oaktree.Flavor")
	proto.RegisterType((*FlavorList)(nil), "oaktree.FlavorList")
}

func init() { proto.RegisterFile("flavor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x51, 0x4d, 0x4f, 0xc2, 0x30,
	0x18, 0x0e, 0x1b, 0x1f, 0xdb, 0x0b, 0x8a, 0xbe, 0xf1, 0xd0, 0xa0, 0x09, 0x84, 0x93, 0x1e, 0xdc,
	0x41, 0x0f, 0x1a, 0x13, 0xe3, 0x49, 0x4e, 0x1c, 0x48, 0xff, 0x00, 0x29, 0xa3, 0xc4, 0x86, 0x6d,
	0x5d, 0xda, 0x82, 0xf0, 0x7b, 0xfc, 0xa3, 0xf6, 0x03, 0xd0, 0x70, 0x7b, 0x3e, 0xd3, 0xf6, 0x29,
	0xf4, 0x56, 0x05, 0xdb, 0x4a, 0x95, 0xd5, 0x4a, 0x1a, 0x89, 0x1d, 0xc9, 0xd6, 0x46, 0x71, 0x3e,
	0xe8, 0xe5, 0xb2, 0x2c, 0x65, 0x15, 0xe4, 0xf1, 0x4f, 0x0c, 0xed, 0x89, 0xcf, 0xe1, 0x23, 0x24,
	0x85, 0xcc, 0x99, 0x11, 0xb2, 0x22, 0x8d, 0x51, 0xe3, 0xbe, 0xfb, 0x74, 0x9d, 0x1d, 0x4a, 0xd9,
	0xf4, 0x60, 0xd0, 0x53, 0x04, 0x2f, 0x21, 0x12, 0x4b, 0x12, 0xd9, 0x60, 0x4a, 0x2d, 0x42, 0x84,
	0x66, 0xc5, 0x4a, 0x4e, 0x62, 0xaf, 0x78, 0x8c, 0xb7, 0x90, 0x0a, 0x3d, 0xaf, 0x37, 0x8b, 0x42,
	0xe4, 0xa4, 0x69, 0x8d, 0x84, 0x26, 0x42, 0xcf, 0x3c, 0xc7, 0x01, 0x24, 0x4b, 0xa1, 0xd9, 0xa2,
	0xe0, 0x4b, 0xd2, 0x0a, 0xde, 0x91, 0xe3, 0x15, 0xc4, 0x8a, 0x95, 0xa4, 0x6d, 0xe5, 0x0b, 0xea,
	0x20, 0xde, 0x40, 0x6b, 0x9b, 0xd7, 0x1b, 0x4d, 0x3a, 0x5e, 0x0b, 0xc4, 0x1d, 0x6a, 0x3b, 0x6b,
	0x92, 0x58, 0xb1, 0x49, 0x3d, 0xc6, 0x3b, 0x48, 0x79, 0xfd, 0xc5, 0x4b, 0xae, 0x58, 0x41, 0x52,
	0x9f, 0xfe, 0x13, 0x5c, 0x43, 0x7f, 0xb3, 0x9a, 0x40, 0x68, 0x38, 0x8c, 0x43, 0xe8, 0xaa, 0x9d,
	0xd9, 0xcd, 0x57, 0x2c, 0x37, 0x52, 0x91, 0xae, 0xb5, 0x22, 0x0a, 0x4e, 0x9a, 0x78, 0x05, 0x3f,
	0x00, 0xec, 0x5c, 0x35, 0x57, 0x46, 0x70, 0x4d, 0xf2, 0x51, 0x6c, 0xc7, 0x19, 0x9e, 0xc6, 0x09,
	0xfb, 0x65, 0xb3, 0x53, 0xe2, 0xb3, 0x32, 0x6a, 0x4f, 0xff, 0x55, 0x06, 0xef, 0xd0, 0x3f, 0xb3,
	0xdd, 0x13, 0xd7, 0x7c, 0xef, 0x97, 0x4e, 0xa9, 0x83, 0xfe, 0x89, 0xac, 0xd8, 0xf0, 0xc3, 0xa8,
	0x81, 0xbc, 0x45, 0xaf, 0x8d, 0xf1, 0x0b, 0x40, 0x38, 0x64, 0x2a, 0xb4, 0xc1, 0x07, 0xe8, 0x84,
	0xaf, 0xd5, 0xb6, 0xed, 0xae, 0xd2, 0x3f, 0xbb, 0x0a, 0x3d, 0xfa, 0x8b, 0xb6, 0xff, 0xe5, 0xe7,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xa6, 0x39, 0x8f, 0x0c, 0x02, 0x00, 0x00,
}
