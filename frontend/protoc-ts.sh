#!/bin/bash

PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
PROTOC_OUT_DIR="./src/generated"
mkdir -p ${PROTOC_OUT_DIR}
protoc \
       --proto_path="../proto" \
       --proto_path="../third_party" \
       --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
       --js_out="import_style=commonjs,binary:${PROTOC_OUT_DIR}" \
       --ts_out="service=grpc-web:${PROTOC_OUT_DIR}" \
       service.proto
