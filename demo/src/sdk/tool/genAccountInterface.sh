#!/bin/sh

PRJ_PATH=D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/iscap/demo/src
PROTO_PATH="$PRJ_PATH/sdk/interface/accountinterface"
PROTO_FILE="account-service.proto"
OUTPUT_FILEPATH="$PRJ_PATH/sdk/interface/accountinterface"

protoc -I$PROTO_PATH $PROTO_FILE  --go_out=plugins=grpc:$OUTPUT_FILEPATH

