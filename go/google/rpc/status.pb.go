// Code generated by protoc-gen-go.
// source: google/rpc/status.proto
// DO NOT EDIT!

package google_rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The `Status` type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs. It is used by
// [gRPC](https://github.com/grpc). The error model is designed to be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error message,
// and error details. The error code should be an enum value of
// [google.rpc.Code][google.rpc.Code], but it may accept additional error codes if needed.  The
// error message should be a developer-facing English message that helps
// developers *understand* and *resolve* the error. If a localized user-facing
// error message is needed, put the localized message in the error details or
// localize it in the client. The optional error details may contain arbitrary
// information about the error. There is a predefined set of error detail types
// in the package `google.rpc` which can be used for common error conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error model, but it
// is not necessarily the actual wire format. When the `Status` message is
// exposed in different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety of
// environments, either with or without APIs, to provide a
// consistent developer experience across different environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the client,
//     it may embed the `Status` in the normal response to indicate the partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch response, the
//     `Status` message should be used directly inside batch response, one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous operation
//     results in its response, the status of those operations should be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message `Status` could
//     be used directly after any stripping needed for security/privacy reasons.
type Status struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A list of messages that carry the error details.  There will be a
	// common set of message types for APIs to use.
	Details []*google_protobuf1.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Status) GetDetails() []*google_protobuf1.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "google.rpc.Status")
}

var fileDescriptor2 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x48, 0xe8, 0x01, 0x25, 0xa4, 0x24, 0xa1, 0x8a, 0xc0, 0x32,
	0x49, 0xa5, 0x69, 0xfa, 0x89, 0x79, 0x95, 0x10, 0x65, 0x4a, 0x7e, 0x5c, 0x6c, 0xc1, 0x60, 0x6d,
	0x42, 0x3c, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x42, 0xfc,
	0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x40, 0x01, 0x4e, 0x21, 0x55,
	0x2e, 0xf6, 0x94, 0xd4, 0x92, 0xc4, 0xcc, 0x9c, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0x6e, 0x23,
	0x11, 0x3d, 0xa8, 0x0d, 0x30, 0x53, 0xf5, 0x1c, 0xf3, 0x2a, 0x9d, 0xe4, 0xb9, 0xf8, 0x92, 0xf3,
	0x73, 0xf5, 0x10, 0x96, 0x3b, 0x71, 0x43, 0xcc, 0x0f, 0x00, 0xa9, 0x0a, 0x60, 0x4c, 0x62, 0x03,
	0x2b, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x54, 0xb4, 0x8a, 0x6b, 0xb9, 0x00, 0x00, 0x00,
}
