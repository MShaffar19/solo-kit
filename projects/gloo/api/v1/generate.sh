#!/usr/bin/env bash

OUT=../../pkg/api/v1/
mkdir -p ${OUT}
protoc -I=./ \
    -I=${GOPATH}/src/github.com/gogo/protobuf/ \
    -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
    -I=${GOPATH}/src \
    --gogo_out=Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types:${GOPATH}/src/ \
    --plugin=protoc-gen-solo-kit=${GOPATH}/bin/protoc-gen-solo-kit --solo-kit_out=${OUT} \
    *.proto