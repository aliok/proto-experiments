#!/usr/bin/env bash

# REPO_ROOT_DIR is the `./` directory
REPO_ROOT_DIR=$(dirname "$0")

# Generate proto files
protoc \
  --proto_path=$REPO_ROOT_DIR/contract1 \
  --go_out=$REPO_ROOT_DIR/contract1 \
  --go_opt=paths=source_relative \
  contract1/contract1.proto \
  --go_opt=Mcontract1.proto=$REPO_ROOT_DIR/contract1

protoc \
  --proto_path=$REPO_ROOT_DIR/contract2 \
  --go_out=$REPO_ROOT_DIR/contract2 \
  --go_opt=paths=source_relative \
  contract2/contract2.proto \
  --go_opt=Mcontract2.proto=$REPO_ROOT_DIR/contract2

# Run go mod tidy and vendor
go mod tidy
go mod vendor
