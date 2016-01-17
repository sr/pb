[![CircleCI](https://circleci.com/gh/peter-edge/pb/tree/master.png)](https://circleci.com/gh/peter-edge/pb/tree/master)
[![Go Report Card](http://goreportcard.com/badge/peter-edge/pb)](http://goreportcard.com/report/peter-edge/pb)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/go.pedge.io/pb/go)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/go.pedge.io/pb/gogo)
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/peter-edge/pb/blob/master/LICENSE)

## pb

Compiled well-known types from google/protobuf, compiled types from googleapis, and my protocol buffer standard library.

Compiled for both golang/protobuf and gogo/protobuf.

### Exceptions

#### golang/protobuf

* google/protobuf/descriptor.proto: github.com/golang/protobuf/protoc-gen-go/descriptor
* google/api/annotations.proto: github.com/gengo/grpc-gateway/third_party/googleapis/google/api
* google/api/http.proto: github.com/gengo/grpc-gateway/third_party/googleapis/google/api

#### gogo/protobuf

* google/protobuf/descriptor.proto: github.com/gogo/protobuf/protoc-gen-gogo/descriptor
* google/api/annotations.proto: github.com/peter-edge/grpc-gateway-gogo/third_party/googleapis/google/api
* google/api/http.proto: github.com/peter-edge/grpc-gateway-gogo/third_party/googleapis/google/api
