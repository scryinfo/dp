#!/bin/sh

PRJ_PATH=`pwd`/../../
PROTO_PATH="./"
PROTO_FILE="*.proto"
GO_OUTPUT_FILEPATH="$PRJ_PATH/api/go"
JS_OUTPUT_FILEPATH="$PRJ_PATH/api/js"

protoc --go_out=plugins=grpc:$GO_OUTPUT_FILEPATH $PROTO_FILE
protoc --js_out="import_style=commonjs,binary:${JS_OUTPUT_FILEPATH}" --ts_out="service=true:${JS_OUTPUT_FILEPATH}" $PROTO_FILE