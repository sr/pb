// Code generated by protoc-gen-gogo.
// source: google/api/monitored_resource.proto
// DO NOT EDIT!

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A descriptor that describes the schema of [MonitoredResource][google.api.MonitoredResource].
type MonitoredResourceDescriptor struct {
	// The monitored resource type. For example, the type `"cloudsql_database"`
	// represents databases in Google Cloud SQL.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// A concise name for the monitored resource type that can be displayed in
	// user interfaces. For example, `"Google Cloud SQL Database"`.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,proto3" json:"display_name,omitempty"`
	// A detailed description of the monitored resource type that can be used in
	// documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// A set of labels that can be used to describe instances of this monitored
	// resource type. For example, Google Cloud SQL databases can be labeled with
	// their `"database_id"` and their `"zone"`.
	Labels []*LabelDescriptor `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty"`
}

func (m *MonitoredResourceDescriptor) Reset()         { *m = MonitoredResourceDescriptor{} }
func (m *MonitoredResourceDescriptor) String() string { return proto.CompactTextString(m) }
func (*MonitoredResourceDescriptor) ProtoMessage()    {}

func (m *MonitoredResourceDescriptor) GetLabels() []*LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

// A monitored resource describes a resource that can be used for monitoring
// purpose. It can also be used for logging, billing, and other purposes. Each
// resource has a `type` and a set of `labels`. The labels contain information
// that identifies the resource and describes attributes of it. For example,
// you can use monitored resource to describe a normal file, where the resource
// has `type` as `"file"`, the label `path` identifies the file, and the label
// `size` describes the file size. The monitoring system can use a set of
// monitored resources of files to generate file size distribution.
type MonitoredResource struct {
	// The monitored resource type. This field must match the corresponding
	// [MonitoredResourceDescriptor.type][google.api.MonitoredResourceDescriptor.type] to this resource..  For example,
	// `"cloudsql_database"` represents Cloud SQL databases.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Values for some or all of the labels listed in the associated monitored
	// resource descriptor. For example, you specify a specific Cloud SQL database
	// by supplying values for both the `"database_id"` and `"zone"` labels.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MonitoredResource) Reset()         { *m = MonitoredResource{} }
func (m *MonitoredResource) String() string { return proto.CompactTextString(m) }
func (*MonitoredResource) ProtoMessage()    {}

func (m *MonitoredResource) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*MonitoredResourceDescriptor)(nil), "google.api.MonitoredResourceDescriptor")
	proto.RegisterType((*MonitoredResource)(nil), "google.api.MonitoredResource")
}
func (m *MonitoredResourceDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MonitoredResourceDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMonitoredResource(data, i, uint64(len(m.Type)))
		i += copy(data[i:], m.Type)
	}
	if len(m.DisplayName) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintMonitoredResource(data, i, uint64(len(m.DisplayName)))
		i += copy(data[i:], m.DisplayName)
	}
	if len(m.Description) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintMonitoredResource(data, i, uint64(len(m.Description)))
		i += copy(data[i:], m.Description)
	}
	if len(m.Labels) > 0 {
		for _, msg := range m.Labels {
			data[i] = 0x22
			i++
			i = encodeVarintMonitoredResource(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *MonitoredResource) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MonitoredResource) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMonitoredResource(data, i, uint64(len(m.Type)))
		i += copy(data[i:], m.Type)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			data[i] = 0x12
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovMonitoredResource(uint64(len(k))) + 1 + len(v) + sovMonitoredResource(uint64(len(v)))
			i = encodeVarintMonitoredResource(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintMonitoredResource(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintMonitoredResource(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	return i, nil
}

func encodeFixed64MonitoredResource(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32MonitoredResource(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMonitoredResource(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *MonitoredResourceDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovMonitoredResource(uint64(l))
		}
	}
	return n
}

func (m *MonitoredResource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMonitoredResource(uint64(len(k))) + 1 + len(v) + sovMonitoredResource(uint64(len(v)))
			n += mapEntrySize + 1 + sovMonitoredResource(uint64(mapEntrySize))
		}
	}
	return n
}

func sovMonitoredResource(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMonitoredResource(x uint64) (n int) {
	return sovMonitoredResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MonitoredResourceDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoredResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MonitoredResourceDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoredResourceDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, &LabelDescriptor{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoredResource(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MonitoredResource) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoredResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MonitoredResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoredResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoredResource(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoredResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMonitoredResource(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitoredResource
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMonitoredResource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMonitoredResource
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMonitoredResource
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMonitoredResource(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMonitoredResource = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitoredResource   = fmt.Errorf("proto: integer overflow")
)
