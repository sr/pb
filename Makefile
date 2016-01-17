all: gen test

dlgen: dl gen

gen: genpb proto build

deps:
	go get -d -v ./...

updatedeps:
	go get -d -v -u -f ./...

testdeps:
	go get -d -v -t ./...

updatetestdeps:
	go get -d -v -t -u -f ./...

build: deps
	go build ./...

install: deps
	go install ./...

dl:
	curl -sSL https://raw.githubusercontent.com/datasets/country-codes/master/data/country-codes.csv > etc/data/country-codes.csv
	rm -rf /tmp/googleapis
	git clone https://github.com/google/googleapis.git /tmp/googleapis
	rm -rf proto/google
	mv /tmp/googleapis/google proto/google
	rm -rf /tmp/google-protobuf
	git clone https://github.com/google/protobuf.git /tmp/google-protobuf
	rm -rf proto/google/protobuf
	mkdir -p proto/google/protobuf
	for file in $$(ls /tmp/google-protobuf/src/google/protobuf/*\.proto | grep -v test); do \
		cp $$file proto/google/protobuf/; \
	done

genpb:
	sh -x etc/bin/gen-money.sh
	sh -x etc/bin/gen-geo.sh

proto:
	go get -v go.pedge.io/protoeasy/cmd/protoeasy
	go get -v go.pedge.io/pkg/cmd/strip-package-comments
	protoeasy \
		--go \
		--go-rel-out go \
		--go-no-default-modifiers \
		--go-import-path go.pedge.io/pb/go \
		--go-modifier google/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor \
		--gogo \
		--gogo-rel-out gogo \
		--gogo-no-default-modifiers \
		--gogo-import-path go.pedge.io/pb/gogo \
		--gogo-modifier google/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor \
		--no-default-includes \
		--exclude google/protobuf/descriptor.proto \
		--out . \
		proto
	find . -name *\.pb\*\.go | xargs strip-package-comments

lint: testdeps
	go get -v github.com/golang/lint/golint
	for file in $$(find go -name '*.go' | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go'); do \
		golint $$file; \
		if [ -n "$$(golint $$file)" ]; then \
			exit 1; \
		fi; \
	done

vet: testdeps
	go vet ./...

errcheck: testdeps
	go get -v github.com/kisielk/errcheck
	errcheck ./...

pretest: lint vet errcheck

test: pretest
	go test ./...

clean:
	go clean -i ./...

.PHONY: \
	all \
	dlgen \
	gen \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	install \
	dl \
	genpb \
	proto \
	lint \
	vet \
	errcheck \
	pretest \
	test \
	clean
