#!/bin/sh

set -e

DIR="$(cd "$(dirname "${0}")/../.." && pwd)"
cd "${DIR}"

export CSV_FILE="etc/data/country-codes.csv"
export GO_TMPL_FILE="etc/tmpl/pb/money/money.gen.go.tmpl"
export GO_FILE="go/pb/money/money.gen.go"
export GOGO_FILE="gogo/pb/money/money.gen.go"
export PB_TMPL_FILE="etc/tmpl/pb/money/money.gen.proto.tmpl"
export PB_FILE="proto/pb/money/money.gen.proto"

mkdir -p go/pb/money
mkdir -p gogo/pb/money
go run "etc/cmd/gen-money/main.go"
gofmt -s -w "${GO_FILE}"
gofmt -s -w "${GOGO_FILE}"
