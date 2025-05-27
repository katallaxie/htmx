#!/bin/bash
# This script is executed after the creation of a new project.

git submodule update --init --recursive
sudo apt update && \
    sudo apt install -y direnv protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/yoheimuta/protolint/cmd/protolint@latest