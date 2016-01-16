// Code generated by protoc-gen-go.
// source: google/datastore/v1beta3/entity.proto
// DO NOT EDIT!

package google_datastore_v1beta3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"
import google_type "go.pedge.io/pb/go/google/type"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A partition ID identifies a grouping of entities. The grouping is always
// by project and namespace, however the namespace ID may be empty.
//
// A partition ID contains several dimensions:
// project ID and namespace ID.
// Partition dimensions:
// - A dimension may be `""`.
// - A dimension must be valid UTF-8 bytes.
// - A dimension's value must match regex `[A-Za-z\d\.\-_]{1,100}`
// If the value of any dimension matches regex `__.*__`, the partition is
// reserved/read-only.
// A reserved/read-only partition ID is forbidden in certain documented
// contexts.
//
// Foreign partition IDs (in which the project ID does
// not match the context project ID ) are discouraged.
// Reads and writes of foreign partition IDs may fail if the project is not in an active state.
type PartitionId struct {
	// Project ID.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id" json:"project_id,omitempty"`
	// Namespace ID.
	NamespaceId string `protobuf:"bytes,4,opt,name=namespace_id" json:"namespace_id,omitempty"`
}

func (m *PartitionId) Reset()                    { *m = PartitionId{} }
func (m *PartitionId) String() string            { return proto.CompactTextString(m) }
func (*PartitionId) ProtoMessage()               {}
func (*PartitionId) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// A unique identifier for an entity.
// If a key's partition id or any of its path kinds or names are
// reserved/read-only, the key is reserved/read-only.
// A reserved/read-only key is forbidden in certain documented contexts.
type Key struct {
	// Entities are partitioned into subsets, currently identified by a dataset
	// (usually implicitly specified by the project) and namespace ID.
	// Queries are scoped to a single partition.
	PartitionId *PartitionId `protobuf:"bytes,1,opt,name=partition_id" json:"partition_id,omitempty"`
	// The entity path.
	// An entity path consists of one or more elements composed of a kind and a
	// string or numerical identifier, which identify entities. The first
	// element identifies a _root entity_, the second element identifies
	// a _child_ of the root entity, the third element a child of the
	// second entity, and so forth. The entities identified by all prefixes of
	// the path are called the element's _ancestors_.
	// An entity path is always fully complete: *all* of the entity's ancestors
	// are required to be in the path along with the entity identifier itself.
	// The only exception is that in some documented cases, the identifier in the
	// last path element (for the entity) itself may be omitted. A path can never
	// be empty. The path can have at most 100 elements.
	Path []*Key_PathElement `protobuf:"bytes,2,rep,name=path" json:"path,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Key) GetPartitionId() *PartitionId {
	if m != nil {
		return m.PartitionId
	}
	return nil
}

func (m *Key) GetPath() []*Key_PathElement {
	if m != nil {
		return m.Path
	}
	return nil
}

// A (kind, ID/name) pair used to construct a key path.
//
// If either name nor ID is set, the element is complete.
// If neither is set, the element is incomplete.
type Key_PathElement struct {
	// The kind of the entity.
	// A kind matching regex `__.*__` is reserved/read-only.
	// A kind must not contain more than 1500 bytes when UTF-8 encoded.
	// Cannot be `""`.
	Kind string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// The type of id.
	//
	// Types that are valid to be assigned to IdType:
	//	*Key_PathElement_Id
	//	*Key_PathElement_Name
	IdType isKey_PathElement_IdType `protobuf_oneof:"id_type"`
}

func (m *Key_PathElement) Reset()                    { *m = Key_PathElement{} }
func (m *Key_PathElement) String() string            { return proto.CompactTextString(m) }
func (*Key_PathElement) ProtoMessage()               {}
func (*Key_PathElement) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type isKey_PathElement_IdType interface {
	isKey_PathElement_IdType()
}

type Key_PathElement_Id struct {
	Id int64 `protobuf:"varint,2,opt,name=id,oneof"`
}
type Key_PathElement_Name struct {
	Name string `protobuf:"bytes,3,opt,name=name,oneof"`
}

func (*Key_PathElement_Id) isKey_PathElement_IdType()   {}
func (*Key_PathElement_Name) isKey_PathElement_IdType() {}

func (m *Key_PathElement) GetIdType() isKey_PathElement_IdType {
	if m != nil {
		return m.IdType
	}
	return nil
}

func (m *Key_PathElement) GetId() int64 {
	if x, ok := m.GetIdType().(*Key_PathElement_Id); ok {
		return x.Id
	}
	return 0
}

func (m *Key_PathElement) GetName() string {
	if x, ok := m.GetIdType().(*Key_PathElement_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Key_PathElement) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Key_PathElement_OneofMarshaler, _Key_PathElement_OneofUnmarshaler, _Key_PathElement_OneofSizer, []interface{}{
		(*Key_PathElement_Id)(nil),
		(*Key_PathElement_Name)(nil),
	}
}

func _Key_PathElement_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Key_PathElement)
	// id_type
	switch x := m.IdType.(type) {
	case *Key_PathElement_Id:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Id))
	case *Key_PathElement_Name:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case nil:
	default:
		return fmt.Errorf("Key_PathElement.IdType has unexpected type %T", x)
	}
	return nil
}

func _Key_PathElement_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Key_PathElement)
	switch tag {
	case 2: // id_type.id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.IdType = &Key_PathElement_Id{int64(x)}
		return true, err
	case 3: // id_type.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.IdType = &Key_PathElement_Name{x}
		return true, err
	default:
		return false, nil
	}
}

func _Key_PathElement_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Key_PathElement)
	// id_type
	switch x := m.IdType.(type) {
	case *Key_PathElement_Id:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Id))
	case *Key_PathElement_Name:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An array value.
type ArrayValue struct {
	// Values in the array.
	// The order of this array may not be preserved if it contains a mix of
	// indexed and unindexed values.
	Values []*Value `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ArrayValue) Reset()                    { *m = ArrayValue{} }
func (m *ArrayValue) String() string            { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()               {}
func (*ArrayValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ArrayValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// A message that can hold any of the supported value types and associated
// metadata.
type Value struct {
	// Must have a value set.
	//
	// Types that are valid to be assigned to ValueType:
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_KeyValue
	//	*Value_StringValue
	//	*Value_BlobValue
	//	*Value_GeoPointValue
	//	*Value_EntityValue
	//	*Value_ArrayValue
	ValueType isValue_ValueType `protobuf_oneof:"value_type"`
	// The `meaning` field should only be populated for backwards compatibility.
	Meaning int32 `protobuf:"varint,14,opt,name=meaning" json:"meaning,omitempty"`
	// If the value should be excluded from all indexes including those defined
	// explicitly.
	ExcludeFromIndexes bool `protobuf:"varint,19,opt,name=exclude_from_indexes" json:"exclude_from_indexes,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	NullValue google_protobuf1.NullValue `protobuf:"varint,11,opt,name=null_value,enum=google.protobuf.NullValue,oneof"`
}
type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,oneof"`
}
type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,oneof"`
}
type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,oneof"`
}
type Value_TimestampValue struct {
	TimestampValue *google_protobuf2.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,oneof"`
}
type Value_KeyValue struct {
	KeyValue *Key `protobuf:"bytes,5,opt,name=key_value,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,17,opt,name=string_value,oneof"`
}
type Value_BlobValue struct {
	BlobValue []byte `protobuf:"bytes,18,opt,name=blob_value,proto3,oneof"`
}
type Value_GeoPointValue struct {
	GeoPointValue *google_type.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,oneof"`
}
type Value_EntityValue struct {
	EntityValue *Entity `protobuf:"bytes,6,opt,name=entity_value,oneof"`
}
type Value_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,oneof"`
}

func (*Value_NullValue) isValue_ValueType()      {}
func (*Value_BooleanValue) isValue_ValueType()   {}
func (*Value_IntegerValue) isValue_ValueType()   {}
func (*Value_DoubleValue) isValue_ValueType()    {}
func (*Value_TimestampValue) isValue_ValueType() {}
func (*Value_KeyValue) isValue_ValueType()       {}
func (*Value_StringValue) isValue_ValueType()    {}
func (*Value_BlobValue) isValue_ValueType()      {}
func (*Value_GeoPointValue) isValue_ValueType()  {}
func (*Value_EntityValue) isValue_ValueType()    {}
func (*Value_ArrayValue) isValue_ValueType()     {}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (m *Value) GetNullValue() google_protobuf1.NullValue {
	if x, ok := m.GetValueType().(*Value_NullValue); ok {
		return x.NullValue
	}
	return google_protobuf1.NullValue_NULL_VALUE
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValueType().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetIntegerValue() int64 {
	if x, ok := m.GetValueType().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValueType().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetTimestampValue() *google_protobuf2.Timestamp {
	if x, ok := m.GetValueType().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Value) GetKeyValue() *Key {
	if x, ok := m.GetValueType().(*Value_KeyValue); ok {
		return x.KeyValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValueType().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBlobValue() []byte {
	if x, ok := m.GetValueType().(*Value_BlobValue); ok {
		return x.BlobValue
	}
	return nil
}

func (m *Value) GetGeoPointValue() *google_type.LatLng {
	if x, ok := m.GetValueType().(*Value_GeoPointValue); ok {
		return x.GeoPointValue
	}
	return nil
}

func (m *Value) GetEntityValue() *Entity {
	if x, ok := m.GetValueType().(*Value_EntityValue); ok {
		return x.EntityValue
	}
	return nil
}

func (m *Value) GetArrayValue() *ArrayValue {
	if x, ok := m.GetValueType().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_KeyValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BlobValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_EntityValue)(nil),
		(*Value_ArrayValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_NullValue:
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.NullValue))
	case *Value_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_IntegerValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_TimestampValue:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampValue); err != nil {
			return err
		}
	case *Value_KeyValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeyValue); err != nil {
			return err
		}
	case *Value_StringValue:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BlobValue:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BlobValue)
	case *Value_GeoPointValue:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GeoPointValue); err != nil {
			return err
		}
	case *Value_EntityValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EntityValue); err != nil {
			return err
		}
	case *Value_ArrayValue:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ArrayValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.ValueType has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 11: // value_type.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_NullValue{google_protobuf1.NullValue(x)}
		return true, err
	case 1: // value_type.boolean_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_BooleanValue{x != 0}
		return true, err
	case 2: // value_type.integer_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_IntegerValue{int64(x)}
		return true, err
	case 3: // value_type.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ValueType = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 10: // value_type.timestamp_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Timestamp)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_TimestampValue{msg}
		return true, err
	case 5: // value_type.key_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Key)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_KeyValue{msg}
		return true, err
	case 17: // value_type.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ValueType = &Value_StringValue{x}
		return true, err
	case 18: // value_type.blob_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ValueType = &Value_BlobValue{x}
		return true, err
	case 8: // value_type.geo_point_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_type.LatLng)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_GeoPointValue{msg}
		return true, err
	case 6: // value_type.entity_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Entity)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_EntityValue{msg}
		return true, err
	case 9: // value_type.array_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ArrayValue)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_ArrayValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_NullValue:
		n += proto.SizeVarint(11<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Value_BooleanValue:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *Value_IntegerValue:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case *Value_TimestampValue:
		s := proto.Size(x.TimestampValue)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_KeyValue:
		s := proto.Size(x.KeyValue)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_StringValue:
		n += proto.SizeVarint(17<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BlobValue:
		n += proto.SizeVarint(18<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BlobValue)))
		n += len(x.BlobValue)
	case *Value_GeoPointValue:
		s := proto.Size(x.GeoPointValue)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_EntityValue:
		s := proto.Size(x.EntityValue)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ArrayValue:
		s := proto.Size(x.ArrayValue)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An entity.
//
// An entity is limited to 1 megabyte when stored. That _roughly_
// corresponds to a limit of 1 megabyte for the serialized form of this
// message.
type Entity struct {
	// The entity's key.
	//
	// An entity must have a key, unless otherwise documented (for example,
	// an entity in `Value.entity_value` may have no key).
	// An entity's kind is its key's path's last element's kind,
	// or null if it has no key.
	Key *Key `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The entity's properties.
	// The map's keys are property names.
	// A property name matching regex `__.*__` is reserved.
	// A reserved property name is forbidden in certain documented contexts.
	// The name must not contain more than 500 characters.
	// The name cannot be `""`.
	Properties map[string]*Value `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Entity) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Entity) GetProperties() map[string]*Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*PartitionId)(nil), "google.datastore.v1beta3.PartitionId")
	proto.RegisterType((*Key)(nil), "google.datastore.v1beta3.Key")
	proto.RegisterType((*Key_PathElement)(nil), "google.datastore.v1beta3.Key.PathElement")
	proto.RegisterType((*ArrayValue)(nil), "google.datastore.v1beta3.ArrayValue")
	proto.RegisterType((*Value)(nil), "google.datastore.v1beta3.Value")
	proto.RegisterType((*Entity)(nil), "google.datastore.v1beta3.Entity")
}

var fileDescriptor1 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x4b, 0xdc, 0x40,
	0x14, 0x35, 0xee, 0x87, 0xee, 0x4d, 0xea, 0xd2, 0x51, 0xda, 0x65, 0xb1, 0x28, 0x4b, 0x85, 0xb6,
	0xd0, 0xc4, 0x2a, 0x45, 0xa9, 0xf4, 0xa1, 0x8b, 0x82, 0xa5, 0x52, 0x96, 0x52, 0xfa, 0x1a, 0x26,
	0x9b, 0x31, 0xa6, 0x26, 0x33, 0x21, 0x3b, 0x11, 0xf3, 0xd6, 0xbf, 0xd6, 0x3f, 0xd0, 0xdf, 0xd4,
	0x3b, 0x1f, 0x59, 0x8b, 0x25, 0xfa, 0x96, 0xdc, 0x73, 0xce, 0xdc, 0x33, 0xf7, 0x9e, 0x81, 0xbd,
	0x44, 0x88, 0x24, 0x63, 0x41, 0x4c, 0x25, 0x5d, 0x48, 0x51, 0xb2, 0xe0, 0xe6, 0x5d, 0xc4, 0x24,
	0x3d, 0x0c, 0x18, 0x97, 0xa9, 0xac, 0xfd, 0xa2, 0x14, 0x52, 0x90, 0x91, 0xa1, 0xf9, 0x4b, 0x9a,
	0x6f, 0x69, 0xe3, 0x6d, 0x7b, 0x00, 0x2d, 0xd2, 0x80, 0x72, 0x2e, 0x24, 0x95, 0xa9, 0xe0, 0x0b,
	0xa3, 0x5b, 0xa2, 0xfa, 0x2f, 0xaa, 0x2e, 0x83, 0x85, 0x2c, 0xab, 0xb9, 0xb4, 0xe8, 0xce, 0x7d,
	0x54, 0xa6, 0x39, 0x5b, 0x48, 0x9a, 0x17, 0x96, 0x60, 0xdb, 0x06, 0xb2, 0x2e, 0x58, 0x90, 0x51,
	0x99, 0xf1, 0xc4, 0x20, 0x93, 0x23, 0x70, 0x67, 0xb4, 0x44, 0x87, 0xd8, 0xec, 0x73, 0x4c, 0x08,
	0x00, 0xd6, 0x7f, 0xb2, 0xb9, 0x0c, 0xd3, 0x78, 0xb4, 0xba, 0xeb, 0xbc, 0x1a, 0x90, 0x2d, 0xf0,
	0x38, 0xc5, 0xf3, 0x0a, 0x3a, 0x67, 0xaa, 0xda, 0x55, 0xd5, 0xc9, 0x6f, 0x07, 0x3a, 0x5f, 0x58,
	0x4d, 0x4e, 0xc0, 0x2b, 0x9a, 0x03, 0x14, 0xea, 0x20, 0xea, 0x1e, 0xec, 0xf9, 0x6d, 0x17, 0xf5,
	0xff, 0x6d, 0x77, 0x04, 0xdd, 0x82, 0xca, 0x2b, 0x6c, 0xd4, 0x41, 0xd1, 0xeb, 0x76, 0x11, 0x76,
	0x42, 0xa1, 0xbc, 0x3a, 0xcb, 0x58, 0x8e, 0x03, 0x1d, 0x9f, 0x2a, 0xdb, 0xcb, 0x5f, 0xe2, 0x41,
	0xf7, 0x3a, 0xe5, 0xa6, 0xf9, 0x00, 0xff, 0x56, 0xad, 0xf9, 0xce, 0xf9, 0x0a, 0xd9, 0x80, 0xae,
	0xb2, 0x3f, 0xea, 0x28, 0xec, 0x7c, 0x65, 0x3a, 0x80, 0xb5, 0x34, 0x0e, 0xd5, 0x24, 0x26, 0x1f,
	0x01, 0x3e, 0x95, 0x25, 0xad, 0x7f, 0xd0, 0xac, 0x62, 0x24, 0x80, 0xfe, 0x8d, 0xfa, 0x58, 0xe0,
	0x31, 0xca, 0xce, 0x4e, 0xbb, 0x1d, 0x2d, 0x98, 0xfc, 0xea, 0x42, 0xcf, 0x48, 0xf7, 0x01, 0x78,
	0x95, 0x65, 0xa1, 0xd6, 0x8f, 0x5c, 0xec, 0xb4, 0x71, 0x30, 0x6e, 0xe4, 0xcd, 0x56, 0xfc, 0xaf,
	0x48, 0xd1, 0x7c, 0x74, 0xf5, 0x1c, 0x9e, 0x44, 0x42, 0x64, 0x8c, 0x72, 0x2b, 0x52, 0xd6, 0xd7,
	0x0d, 0x90, 0x72, 0xc9, 0x12, 0x56, 0x5a, 0xa0, 0xb9, 0xc7, 0x33, 0xf0, 0x62, 0x51, 0x45, 0x19,
	0xb3, 0x75, 0x75, 0x1f, 0x07, 0xeb, 0xef, 0x61, 0xb8, 0x5c, 0xb7, 0x85, 0x40, 0xef, 0xe0, 0x7f,
	0x03, 0xdf, 0x1b, 0x1e, 0xca, 0x0e, 0x60, 0x70, 0xcd, 0x6a, 0x2b, 0xe8, 0x69, 0xc1, 0x8b, 0x07,
	0xe7, 0x6f, 0x2c, 0x60, 0xee, 0x52, 0x9e, 0x58, 0xd9, 0x53, 0x33, 0x52, 0x4c, 0x08, 0x44, 0x99,
	0x88, 0x6c, 0x95, 0x60, 0xd5, 0xc3, 0xaa, 0x0f, 0xc3, 0x84, 0x89, 0xb0, 0x10, 0x78, 0x1f, 0x0b,
	0xad, 0xeb, 0x3e, 0x9b, 0x4d, 0x1f, 0xb5, 0x04, 0xff, 0x82, 0xca, 0x0b, 0x9e, 0x20, 0xff, 0x18,
	0x3c, 0xf3, 0x56, 0x2c, 0xb9, 0xaf, 0xc9, 0xbb, 0xed, 0xa6, 0xce, 0x34, 0x1b, 0x95, 0x27, 0xe0,
	0x52, 0xb5, 0x47, 0x2b, 0x1c, 0x68, 0xe1, 0xcb, 0x76, 0xe1, 0xdd, 0xd2, 0x51, 0x3c, 0x84, 0xb5,
	0x1c, 0xd7, 0x80, 0xb7, 0x1a, 0x6d, 0xa0, 0xb0, 0x47, 0xb6, 0x61, 0x8b, 0xdd, 0xce, 0xb3, 0x2a,
	0x66, 0xe1, 0x65, 0x29, 0xf2, 0x10, 0x83, 0xc5, 0x6e, 0x31, 0x15, 0x9b, 0x6a, 0x43, 0x53, 0x0f,
	0x40, 0x77, 0x31, 0x09, 0xfa, 0xe3, 0x40, 0xdf, 0xd8, 0x20, 0x6f, 0xa0, 0x83, 0x03, 0xb5, 0xf9,
	0x7f, 0x78, 0x94, 0xe4, 0x54, 0x3f, 0xb3, 0x82, 0xe1, 0x4b, 0xc0, 0x83, 0x3b, 0x3a, 0x6e, 0xfb,
	0x8f, 0x5d, 0xd4, 0x9f, 0x2d, 0x25, 0x58, 0x28, 0xeb, 0xf1, 0x37, 0x18, 0xde, 0x2b, 0x11, 0xf7,
	0xce, 0xc4, 0x00, 0x17, 0xd0, 0xbb, 0x8b, 0xd0, 0xe3, 0x79, 0xfe, 0xb0, 0x7a, 0xec, 0x4c, 0xdf,
	0xc2, 0xf6, 0x5c, 0xe4, 0xad, 0xcc, 0xa9, 0x6b, 0xbc, 0xcc, 0x54, 0xa4, 0x66, 0x4e, 0xd4, 0xd7,
	0xd9, 0x3a, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xe1, 0x95, 0xba, 0xff, 0x04, 0x00, 0x00,
}
