// Code generated by protoc-gen-go.
// source: micro/go-platform/event/proto/event.proto
// DO NOT EDIT!

/*
Package event is a generated protocol buffer package.

It is generated from these files:
	micro/go-platform/event/proto/event.proto

It has these top-level messages:
	Record
*/
package event

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Record struct {
	Id        string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type      string            `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Origin    string            `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	Timestamp int64             `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	RootId    string            `protobuf:"bytes,5,opt,name=root_id" json:"root_id,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Data      string            `protobuf:"bytes,7,opt,name=data" json:"data,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "Record")
}

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x4f, 0xc1, 0x4e, 0x85, 0x30,
	0x10, 0x4c, 0xe9, 0x7b, 0x45, 0x16, 0xd1, 0x48, 0x62, 0xd2, 0x78, 0x22, 0x5e, 0xc4, 0x83, 0x34,
	0xd1, 0x8b, 0xf1, 0xee, 0xd1, 0x8b, 0x3f, 0x60, 0x2a, 0x54, 0xd2, 0x48, 0xd9, 0xa6, 0xae, 0x24,
	0xfc, 0x9c, 0xdf, 0x26, 0x14, 0x2e, 0xef, 0x36, 0x33, 0x3b, 0xbb, 0x33, 0x0b, 0xf7, 0xce, 0xb6,
	0x01, 0x55, 0x8f, 0x0f, 0x7e, 0xd0, 0xf4, 0x85, 0xc1, 0x29, 0x33, 0x99, 0x91, 0x94, 0x0f, 0x48,
	0xb8, 0xe1, 0x26, 0xe2, 0xdb, 0x3f, 0x06, 0xe2, 0xdd, 0xb4, 0x18, 0xba, 0x12, 0x20, 0xb1, 0x9d,
	0x64, 0x15, 0xab, 0xb3, 0xf2, 0x1c, 0x0e, 0x34, 0x7b, 0x23, 0x93, 0xc8, 0x2e, 0x40, 0x60, 0xb0,
	0xbd, 0x1d, 0x25, 0x8f, 0xfc, 0x0a, 0x32, 0xb2, 0xce, 0xfc, 0x90, 0x76, 0x5e, 0x1e, 0x16, 0x89,
	0x97, 0x97, 0x90, 0x06, 0x44, 0xfa, 0x58, 0x2e, 0x1c, 0xa3, 0xe7, 0x0e, 0xce, 0x9c, 0x21, 0xdd,
	0x69, 0xd2, 0x52, 0x54, 0xbc, 0xce, 0x1f, 0xaf, 0x9b, 0x2d, 0xa8, 0x79, 0xdb, 0xf5, 0xd7, 0x91,
	0xc2, 0xbc, 0x46, 0x45, 0x53, 0xba, 0xae, 0xdd, 0x28, 0x28, 0x4e, 0xc7, 0x39, 0xf0, 0x6f, 0x33,
	0xef, 0xb5, 0x0a, 0x38, 0x4e, 0x7a, 0xf8, 0xdd, 0x7b, 0xbd, 0x24, 0xcf, 0xec, 0x53, 0xc4, 0x3f,
	0x9e, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x10, 0xd5, 0xf0, 0xf4, 0x00, 0x00, 0x00,
}
