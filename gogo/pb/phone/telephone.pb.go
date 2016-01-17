// Code generated by protoc-gen-gogo.
// source: pb/phone/telephone.proto
// DO NOT EDIT!

package pbphone

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TelephoneNumber struct {
	CountryCode             uint32 `protobuf:"varint,1,opt,name=country_code,proto3" json:"country_code,omitempty"`
	NationalDestinationCode uint64 `protobuf:"varint,2,opt,name=national_destination_code,proto3" json:"national_destination_code,omitempty"`
	SubscriberNumber        uint64 `protobuf:"varint,3,opt,name=subscriber_number,proto3" json:"subscriber_number,omitempty"`
	Extension               uint64 `protobuf:"varint,4,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *TelephoneNumber) Reset()         { *m = TelephoneNumber{} }
func (m *TelephoneNumber) String() string { return proto.CompactTextString(m) }
func (*TelephoneNumber) ProtoMessage()    {}

func init() {
	proto.RegisterType((*TelephoneNumber)(nil), "pb.phone.TelephoneNumber")
}
