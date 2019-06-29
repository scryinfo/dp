#!/bin/sh

PRJ_PATH=/home/waitgordor/project/goproject/src/github.com/scryinfo/dp
PROTO_PATH="$PRJ_PATH/api/proto/"
PROTO_FILE="binary.proto"
OUTPUT_FILEPATH="$PRJ_PATH/api/go_out/"

protoc -I$PROTO_PATH $PROTO_FILE  --go_out=plugins=grpc:$OUTPUT_FILEPATH

