#!/bin/sh

PRJ_PATH=/home/waitgordor/project/iscap2/iscap/demo/src
PROTO_PATH="$PRJ_PATH/sdk/interface/accountinterface"
PROTO_FILE="account-service.proto"
OUTPUT_FILEPATH="$PRJ_PATH/sdk/interface/accountinterface"

protoc -I$PROTO_PATH $PROTO_FILE  --go_out=plugins=grpc:$OUTPUT_FILEPATH

