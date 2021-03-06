// Code generated by protoc-gen-go.
// source: google/devtools/cloudtrace/v1/trace.proto
// DO NOT EDIT!

package google_devtools_cloudtrace_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import _ "go.pedge.io/pb/go/google/protobuf"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Type of span. Can be used to specify additional relationships between spans
// in addition to a parent/child relationship.
type TraceSpan_SpanKind int32

const (
	// Unspecified.
	TraceSpan_SPAN_KIND_UNSPECIFIED TraceSpan_SpanKind = 0
	// Indicates that the span covers server-side handling of an RPC or other
	// remote network request.
	TraceSpan_RPC_SERVER TraceSpan_SpanKind = 1
	// Indicates that the span covers the client-side wrapper around an RPC or
	// other remote request.
	TraceSpan_RPC_CLIENT TraceSpan_SpanKind = 2
)

var TraceSpan_SpanKind_name = map[int32]string{
	0: "SPAN_KIND_UNSPECIFIED",
	1: "RPC_SERVER",
	2: "RPC_CLIENT",
}
var TraceSpan_SpanKind_value = map[string]int32{
	"SPAN_KIND_UNSPECIFIED": 0,
	"RPC_SERVER":            1,
	"RPC_CLIENT":            2,
}

func (x TraceSpan_SpanKind) String() string {
	return proto.EnumName(TraceSpan_SpanKind_name, int32(x))
}
func (TraceSpan_SpanKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// Type of data returned for traces in the list.
type ListTracesRequest_ViewType int32

const (
	// Default is `MINIMAL` if unspecified.
	ListTracesRequest_VIEW_TYPE_UNSPECIFIED ListTracesRequest_ViewType = 0
	// Minimal view of the trace record that contains only the project
	// and trace IDs.
	ListTracesRequest_MINIMAL ListTracesRequest_ViewType = 1
	// Root span view of the trace record that returns the root spans along
	// with the minimal trace data.
	ListTracesRequest_ROOTSPAN ListTracesRequest_ViewType = 2
	// Complete view of the trace record that contains the actual trace data.
	// This is equivalent to calling the REST `get` or RPC `GetTrace` method
	// using the ID of each listed trace.
	ListTracesRequest_COMPLETE ListTracesRequest_ViewType = 3
)

var ListTracesRequest_ViewType_name = map[int32]string{
	0: "VIEW_TYPE_UNSPECIFIED",
	1: "MINIMAL",
	2: "ROOTSPAN",
	3: "COMPLETE",
}
var ListTracesRequest_ViewType_value = map[string]int32{
	"VIEW_TYPE_UNSPECIFIED": 0,
	"MINIMAL":               1,
	"ROOTSPAN":              2,
	"COMPLETE":              3,
}

func (x ListTracesRequest_ViewType) String() string {
	return proto.EnumName(ListTracesRequest_ViewType_name, int32(x))
}
func (ListTracesRequest_ViewType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// A trace describes how long it takes for an application to perform an
// operation. It consists of a set of spans, each of which represent a single
// timed event within the operation.
type Trace struct {
	// Project ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id" json:"project_id,omitempty"`
	// Globally unique identifier for the trace. This identifier is a 128-bit
	// numeric value formatted as a 32-byte hex string.
	TraceId string `protobuf:"bytes,2,opt,name=trace_id" json:"trace_id,omitempty"`
	// Collection of spans in the trace.
	Spans []*TraceSpan `protobuf:"bytes,3,rep,name=spans" json:"spans,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Trace) GetSpans() []*TraceSpan {
	if m != nil {
		return m.Spans
	}
	return nil
}

// List of new or updated traces.
type Traces struct {
	// List of traces.
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
}

func (m *Traces) Reset()                    { *m = Traces{} }
func (m *Traces) String() string            { return proto.CompactTextString(m) }
func (*Traces) ProtoMessage()               {}
func (*Traces) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Traces) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// A span represents a single timed event within a trace. Spans can be nested
// and form a trace tree. Often, a trace contains a root span that describes the
// end-to-end latency of an operation and, optionally, one or more subspans for
// its suboperations. Spans do not need to be contiguous. There may be gaps
// between spans in a trace.
type TraceSpan struct {
	// Identifier for the span. This identifier must be unique within a trace.
	SpanId uint64 `protobuf:"fixed64,1,opt,name=span_id" json:"span_id,omitempty"`
	// Distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `RPC_CLIENT`
	// and `RPC_SERVER` to identify queueing latency associated with the span.
	Kind TraceSpan_SpanKind `protobuf:"varint,2,opt,name=kind,enum=google.devtools.cloudtrace.v1.TraceSpan_SpanKind" json:"kind,omitempty"`
	// Name of the trace. The trace name is sanitized and displayed in the
	// Cloud Trace tool in the Google Developers Console. The name may be a method
	// name or some other per-call site name. For the same executable and the same
	// call point, a best practice is to use a consistent name, which makes it
	// easier to correlate cross-trace spans.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Start time of the span in nanoseconds from the UNIX epoch.
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=start_time" json:"start_time,omitempty"`
	// End time of the span in nanoseconds from the UNIX epoch.
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=end_time" json:"end_time,omitempty"`
	// ID of the parent span, if any. Optional.
	ParentSpanId uint64 `protobuf:"fixed64,6,opt,name=parent_span_id" json:"parent_span_id,omitempty"`
	// Collection of labels associated with the span.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TraceSpan) Reset()                    { *m = TraceSpan{} }
func (m *TraceSpan) String() string            { return proto.CompactTextString(m) }
func (*TraceSpan) ProtoMessage()               {}
func (*TraceSpan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TraceSpan) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TraceSpan) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TraceSpan) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// The request message for the `ListTraces` method. All fields are required
// unless specified.
type ListTracesRequest struct {
	// ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id" json:"project_id,omitempty"`
	// Type of data returned for traces in the list. Optional. Default is
	// `MINIMAL`.
	View ListTracesRequest_ViewType `protobuf:"varint,2,opt,name=view,enum=google.devtools.cloudtrace.v1.ListTracesRequest_ViewType" json:"view,omitempty"`
	// Maximum number of traces to return. If not specified or <= 0, the
	// implementation selects a reasonable value.  The implementation may
	// return fewer traces than the requested page size. Optional.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size" json:"page_size,omitempty"`
	// Token identifying the page of results to return. If provided, use the
	// value of the `next_page_token` field from a previous request. Optional.
	PageToken string `protobuf:"bytes,4,opt,name=page_token" json:"page_token,omitempty"`
	// End of the time interval (inclusive) during which the trace data was
	// collected from the application.
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=start_time" json:"start_time,omitempty"`
	// Start of the time interval (inclusive) during which the trace data was
	// collected from the application.
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=end_time" json:"end_time,omitempty"`
	// An optional filter for the request.
	Filter string `protobuf:"bytes,7,opt,name=filter" json:"filter,omitempty"`
	// Field used to sort the returned traces. Optional.
	// Can be one of the following:
	//
	// *   `trace_id`
	// *   `name` (`name` field of root span in the trace)
	// *   `duration` (difference between `end_time` and `start_time` fields of
	//      the root span)
	// *   `start` (`start_time` field of the root span)
	//
	// Descending order can be specified by appending `desc` to the sort field
	// (for example, `name desc`).
	//
	// Only one sort field is permitted.
	OrderBy string `protobuf:"bytes,8,opt,name=order_by" json:"order_by,omitempty"`
}

func (m *ListTracesRequest) Reset()                    { *m = ListTracesRequest{} }
func (m *ListTracesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTracesRequest) ProtoMessage()               {}
func (*ListTracesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListTracesRequest) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ListTracesRequest) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// The response message for the `ListTraces` method.
type ListTracesResponse struct {
	// List of trace records returned.
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
	// If defined, indicates that there are more traces that match the request
	// and that this value should be passed to the next request to continue
	// retrieving additional traces.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token" json:"next_page_token,omitempty"`
}

func (m *ListTracesResponse) Reset()                    { *m = ListTracesResponse{} }
func (m *ListTracesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTracesResponse) ProtoMessage()               {}
func (*ListTracesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListTracesResponse) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// The request message for the `GetTrace` method.
type GetTraceRequest struct {
	// ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id" json:"project_id,omitempty"`
	// ID of the trace to return.
	TraceId string `protobuf:"bytes,2,opt,name=trace_id" json:"trace_id,omitempty"`
}

func (m *GetTraceRequest) Reset()                    { *m = GetTraceRequest{} }
func (m *GetTraceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTraceRequest) ProtoMessage()               {}
func (*GetTraceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// The request message for the `PatchTraces` method.
type PatchTracesRequest struct {
	// ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id" json:"project_id,omitempty"`
	// The body of the message.
	Traces *Traces `protobuf:"bytes,2,opt,name=traces" json:"traces,omitempty"`
}

func (m *PatchTracesRequest) Reset()                    { *m = PatchTracesRequest{} }
func (m *PatchTracesRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchTracesRequest) ProtoMessage()               {}
func (*PatchTracesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PatchTracesRequest) GetTraces() *Traces {
	if m != nil {
		return m.Traces
	}
	return nil
}

func init() {
	proto.RegisterType((*Trace)(nil), "google.devtools.cloudtrace.v1.Trace")
	proto.RegisterType((*Traces)(nil), "google.devtools.cloudtrace.v1.Traces")
	proto.RegisterType((*TraceSpan)(nil), "google.devtools.cloudtrace.v1.TraceSpan")
	proto.RegisterType((*ListTracesRequest)(nil), "google.devtools.cloudtrace.v1.ListTracesRequest")
	proto.RegisterType((*ListTracesResponse)(nil), "google.devtools.cloudtrace.v1.ListTracesResponse")
	proto.RegisterType((*GetTraceRequest)(nil), "google.devtools.cloudtrace.v1.GetTraceRequest")
	proto.RegisterType((*PatchTracesRequest)(nil), "google.devtools.cloudtrace.v1.PatchTracesRequest")
	proto.RegisterEnum("google.devtools.cloudtrace.v1.TraceSpan_SpanKind", TraceSpan_SpanKind_name, TraceSpan_SpanKind_value)
	proto.RegisterEnum("google.devtools.cloudtrace.v1.ListTracesRequest_ViewType", ListTracesRequest_ViewType_name, ListTracesRequest_ViewType_value)
}

var fileDescriptor0 = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0xce, 0x8f, 0x9b, 0x9e, 0xf4, 0xa6, 0xe9, 0x48, 0xb7, 0x37, 0x04, 0x10, 0xc5, 0x02,
	0x29, 0xfc, 0xd4, 0x6e, 0xd2, 0xa2, 0x42, 0x17, 0xa0, 0x36, 0x35, 0x55, 0xd4, 0x34, 0xb5, 0x92,
	0x50, 0xc4, 0xca, 0x72, 0x92, 0x69, 0x31, 0x4d, 0x6c, 0x63, 0x4f, 0x02, 0xa1, 0xea, 0x86, 0x25,
	0x5b, 0xc4, 0x0a, 0xf1, 0x50, 0x88, 0x57, 0x40, 0xe2, 0x35, 0x98, 0x19, 0xdb, 0x69, 0x94, 0xa0,
	0x3a, 0x65, 0x13, 0x65, 0xce, 0x9c, 0x9f, 0xef, 0x7c, 0xdf, 0x99, 0x63, 0xb8, 0x77, 0x62, 0xdb,
	0x27, 0x5d, 0xac, 0x74, 0xf0, 0x80, 0xd8, 0x76, 0xd7, 0x53, 0xda, 0x5d, 0xbb, 0xdf, 0x21, 0xae,
	0xd1, 0xc6, 0xca, 0xa0, 0xa8, 0xf0, 0x3f, 0xb2, 0xe3, 0xda, 0xc4, 0x46, 0x37, 0x7d, 0x57, 0x39,
	0x74, 0x95, 0x2f, 0x5c, 0xe5, 0x41, 0x31, 0x7f, 0x23, 0xc8, 0x64, 0x38, 0xa6, 0x62, 0x58, 0x96,
	0x4d, 0x0c, 0x62, 0xda, 0x96, 0xe7, 0x07, 0xe7, 0xaf, 0x07, 0xb7, 0xfc, 0xd4, 0xea, 0x1f, 0x2b,
	0xb8, 0xe7, 0x90, 0x61, 0x70, 0x79, 0x6b, 0xf2, 0x92, 0x98, 0x3d, 0xec, 0x11, 0xa3, 0xe7, 0xf8,
	0x0e, 0xd2, 0x31, 0x24, 0x9b, 0xac, 0x0e, 0x42, 0x00, 0xd4, 0xf2, 0x06, 0xb7, 0x89, 0x6e, 0x76,
	0x72, 0xc2, 0x8a, 0x50, 0x98, 0x47, 0x59, 0x48, 0x71, 0x10, 0xcc, 0x12, 0xe3, 0x96, 0x4d, 0x48,
	0x7a, 0x8e, 0x61, 0x79, 0xb9, 0xf8, 0x4a, 0xbc, 0x90, 0x2e, 0x15, 0xe4, 0x4b, 0x91, 0xcb, 0x3c,
	0x75, 0x83, 0x06, 0x48, 0x4f, 0x41, 0xe4, 0x07, 0x0f, 0x6d, 0x80, 0xc8, 0xef, 0x3d, 0x5a, 0x84,
	0xe5, 0xb8, 0x33, 0x4b, 0x0e, 0xe9, 0x5b, 0x1c, 0xe6, 0x47, 0xd9, 0xd0, 0x22, 0xcc, 0x31, 0x18,
	0x21, 0x52, 0x11, 0x3d, 0x83, 0xc4, 0xa9, 0x69, 0xf9, 0x28, 0x33, 0xa5, 0xe2, 0xac, 0xb0, 0x64,
	0xf6, 0xb3, 0x4f, 0x03, 0xd1, 0x02, 0x24, 0x2c, 0xa3, 0x87, 0x69, 0x5f, 0xac, 0x4d, 0x19, 0x80,
	0x92, 0xe4, 0x12, 0x9d, 0xd1, 0x95, 0x4b, 0x50, 0x5b, 0xba, 0x94, 0x0f, 0x93, 0x86, 0x5c, 0xca,
	0xcd, 0x90, 0x4b, 0xf4, 0x10, 0x52, 0xd8, 0xea, 0xf8, 0xde, 0xc9, 0x48, 0xef, 0x65, 0xc8, 0x38,
	0x86, 0x8b, 0x2d, 0xa2, 0x87, 0x4d, 0x88, 0xbc, 0x89, 0x5d, 0x10, 0xbb, 0x46, 0x0b, 0x77, 0xbd,
	0xdc, 0x1c, 0x67, 0x66, 0x63, 0xe6, 0x36, 0xaa, 0x3c, 0x4c, 0xb5, 0x88, 0x3b, 0xcc, 0xaf, 0x42,
	0x7a, 0xec, 0x88, 0xd2, 0x10, 0x3f, 0xc5, 0xc3, 0x40, 0xd0, 0x7f, 0x21, 0x39, 0x30, 0xba, 0x7d,
	0xec, 0xab, 0xb9, 0x15, 0x7b, 0x2c, 0x48, 0x2a, 0xa4, 0x46, 0x24, 0x5c, 0x83, 0xff, 0x1a, 0xda,
	0x76, 0x4d, 0xdf, 0xaf, 0xd4, 0x76, 0xf5, 0x17, 0xb5, 0x86, 0xa6, 0x96, 0x2b, 0xcf, 0x2b, 0xea,
	0x6e, 0xf6, 0x1f, 0x94, 0x01, 0xa8, 0x6b, 0x65, 0xbd, 0xa1, 0xd6, 0x8f, 0xd4, 0x7a, 0x56, 0x08,
	0xcf, 0xe5, 0x6a, 0x45, 0xad, 0x35, 0xb3, 0x31, 0xe9, 0x57, 0x0c, 0x96, 0xaa, 0xa6, 0x47, 0x7c,
	0x91, 0xeb, 0xf8, 0x6d, 0x9f, 0x36, 0xfb, 0xc7, 0xa1, 0xda, 0x83, 0xc4, 0xc0, 0xc4, 0xef, 0x02,
	0xa9, 0x9e, 0x44, 0xf4, 0x38, 0x95, 0x53, 0x3e, 0xa2, 0xc1, 0xcd, 0xa1, 0x83, 0xd1, 0x12, 0xcc,
	0x3b, 0xc6, 0x09, 0xd6, 0x3d, 0xf3, 0x83, 0xaf, 0x5b, 0x92, 0xd7, 0x63, 0x26, 0x62, 0x9f, 0x62,
	0x8b, 0xeb, 0x36, 0xa9, 0x65, 0xf2, 0x4a, 0x5a, 0x8a, 0x91, 0xde, 0x19, 0x10, 0x8f, 0xcd, 0x2e,
	0xc1, 0x2e, 0xd5, 0x2c, 0x78, 0x32, 0xb6, 0xdb, 0xc1, 0xae, 0xde, 0x1a, 0xe6, 0x52, 0xcc, 0x22,
	0xd5, 0x20, 0x35, 0x82, 0x4c, 0x09, 0x3e, 0xaa, 0xa8, 0x2f, 0xf5, 0xe6, 0x2b, 0x4d, 0x9d, 0x20,
	0x38, 0x0d, 0x73, 0x07, 0x95, 0x5a, 0xe5, 0x60, 0xbb, 0x4a, 0xd9, 0x5d, 0x80, 0x54, 0xfd, 0xf0,
	0xb0, 0xc9, 0xc4, 0xc8, 0xc6, 0xd8, 0xa9, 0x7c, 0x78, 0xa0, 0x55, 0xd5, 0xa6, 0x9a, 0x8d, 0x4b,
	0x6d, 0x40, 0xe3, 0xa4, 0x78, 0x0e, 0x5d, 0x05, 0xf8, 0xef, 0x5e, 0x15, 0xfa, 0x1f, 0x16, 0x2d,
	0xfc, 0x9e, 0xe8, 0x63, 0xa4, 0xf1, 0xc9, 0x90, 0x36, 0x61, 0x71, 0x0f, 0xfb, 0x35, 0x2e, 0xd3,
	0x72, 0x6a, 0x41, 0x48, 0x3a, 0x20, 0xcd, 0x20, 0xed, 0xd7, 0xd1, 0x73, 0xf0, 0x68, 0x84, 0x38,
	0xc6, 0x59, 0xbe, 0x3b, 0x0b, 0x62, 0xaf, 0xf4, 0x3d, 0x0e, 0x0b, 0xfe, 0xe0, 0x63, 0x77, 0x60,
	0xd2, 0x1e, 0xbe, 0x0a, 0x00, 0x17, 0x84, 0xa0, 0xb5, 0xab, 0x0e, 0x54, 0xbe, 0x78, 0x85, 0x08,
	0x9f, 0x6d, 0xa9, 0xf0, 0xf1, 0xc7, 0xcf, 0xcf, 0x31, 0x09, 0xad, 0xb0, 0x4d, 0x1e, 0x74, 0xe6,
	0x29, 0x67, 0x17, 0x3d, 0x9e, 0xfb, 0xfb, 0xdd, 0x43, 0x5f, 0x04, 0x48, 0x85, 0x4c, 0x22, 0x39,
	0xa2, 0xd2, 0x04, 0xe5, 0xf9, 0xd9, 0x56, 0xe3, 0x3a, 0x07, 0xb3, 0x8a, 0x1e, 0x44, 0x81, 0x51,
	0xce, 0x42, 0xb1, 0xce, 0xd1, 0x27, 0x01, 0xd2, 0x63, 0x42, 0xa1, 0x28, 0x12, 0xa6, 0x45, 0xcd,
	0x2f, 0x4f, 0x3d, 0x0b, 0x95, 0x7d, 0x79, 0xa4, 0x35, 0x8e, 0xe7, 0x7e, 0x29, 0x92, 0x9c, 0xad,
	0x60, 0x00, 0x76, 0x8a, 0x70, 0xbb, 0x6d, 0xf7, 0x2e, 0x47, 0xb0, 0x03, 0xbc, 0xba, 0xc6, 0x6a,
	0x69, 0x42, 0x4b, 0xe4, 0x45, 0xd7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xb8, 0x89, 0xba,
	0x66, 0x07, 0x00, 0x00,
}
