#!/bin/sh

set -e

DIR="$(cd "$(dirname "${0}")/../.." && pwd)"
cd "${DIR}"

export CSV_FILE="etc/data/country-codes.csv"
export GO_TMPL_FILE="etc/tmpl/pb/geo/geo.gen.go.tmpl"
export GO_FILE="go/pb/geo/geo.gen.go"
export GOGO_FILE="gogo/pb/geo/geo.gen.go"
export PB_TMPL_FILE="etc/tmpl/pb/geo/geo.gen.proto.tmpl"
export PB_FILE="proto/pb/geo/geo.gen.proto"

mkdir -p go/pb/geo
mkdir -p gogo/pb/geo
go run "etc/cmd/gen-geo/main.go"
gofmt -s -w "${GO_FILE}"
gofmt -s -w "${GOGO_FILE}"
