// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id          int32                      `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email       string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones      []*Person_PhoneNumber      `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0x6d, 0xd7, 0x95, 0xed, 0x2a, 0xa3, 0x06, 0x91, 0x52, 0x04, 0xcb, 0x9e, 0x0a, 0x42,
	0x06, 0x53, 0xf0, 0xc9, 0x07, 0x07, 0x03, 0x45, 0xe7, 0x46, 0xd8, 0xf0, 0x51, 0x52, 0x1a, 0x67,
	0x59, 0xdb, 0x0b, 0x6d, 0xfa, 0xb0, 0xcf, 0xe6, 0x97, 0x93, 0x26, 0xad, 0x88, 0xf8, 0x76, 0xb9,
	0xfb, 0x71, 0xf9, 0xdd, 0x1f, 0xce, 0x78, 0x9a, 0x56, 0xa2, 0xae, 0x13, 0xc4, 0x03, 0x95, 0x15,
	0x2a, 0x24, 0x23, 0xd5, 0x28, 0xac, 0x32, 0x9e, 0x87, 0x57, 0x7b, 0xc4, 0x7d, 0x2e, 0x66, 0xba,
	0x9f, 0x34, 0x1f, 0x33, 0x95, 0x15, 0xa2, 0x56, 0xbc, 0x90, 0x06, 0x9d, 0x7e, 0xd9, 0xe0, 0x6e,
	0x44, 0x55, 0x63, 0x49, 0x08, 0x38, 0x25, 0x2f, 0x44, 0x60, 0x45, 0x56, 0x3c, 0x66, 0xba, 0x26,
	0x13, 0xb0, 0xb3, 0x34, 0xb0, 0x23, 0x2b, 0x1e, 0x32, 0x3b, 0x4b, 0xc9, 0x39, 0x0c, 0x45, 0xc1,
	0xb3, 0x3c, 0x18, 0x68, 0xc8, 0x3c, 0xc8, 0x2d, 0xb8, 0xf2, 0x13, 0x4b, 0x51, 0x07, 0x4e, 0x34,
	0x88, 0xbd, 0xf9, 0x25, 0xed, 0x05, 0xa8, 0xd9, 0x4d, 0x37, 0xed, 0xf8, 0xb5, 0x29, 0x12, 0x51,
	0xb1, 0x8e, 0x25, 0xf7, 0x70, 0x9a, 0xf3, 0x5a, 0xbd, 0x37, 0x32, 0xe5, 0x4a, 0xa4, 0xc1, 0x30,
	0xb2, 0x62, 0x6f, 0x1e, 0x52, 0xa3, 0x4c, 0x7b, 0x65, 0xba, 0xed, 0x95, 0x99, 0xd7, 0xf2, 0x3b,
	0x83, 0x87, 0x3b, 0xf0, 0x7e, 0x6d, 0x25, 0x17, 0xe0, 0x96, 0xba, 0xea, 0xfc, 0xbb, 0x17, 0xa1,
	0xe0, 0xa8, 0xa3, 0x14, 0xfa, 0x86, 0xc9, 0x3c, 0xfc, 0xdf, 0x6c, 0x7b, 0x94, 0x82, 0x69, 0x6e,
	0x7a, 0x0d, 0xe3, 0x9f, 0x16, 0x01, 0x70, 0x57, 0xeb, 0xc5, 0xd3, 0xcb, 0xd2, 0x3f, 0x21, 0x23,
	0x70, 0x1e, 0xd7, 0xab, 0xa5, 0x6f, 0xb5, 0xd5, 0xdb, 0x9a, 0x3d, 0xfb, 0xf6, 0xf4, 0x0e, 0xbc,
	0x07, 0x93, 0xfe, 0x02, 0xf1, 0x40, 0x62, 0x70, 0xa5, 0x40, 0x99, 0xb7, 0x19, 0xb6, 0x39, 0xf8,
	0x7f, 0x7f, 0x63, 0xdd, 0x3c, 0x71, 0xf5, 0x75, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x82, 0x41, 0x71, 0xbd, 0x01, 0x00, 0x00,
}