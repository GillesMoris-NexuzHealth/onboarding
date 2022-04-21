#!/bin/bash

protoc --proto_path="../proto" --proto_path="../third_party" --go_out="../proto" service.proto

protoc --proto_path="../proto" --proto_path="../third_party" --go-grpc_out="../proto" service.proto
