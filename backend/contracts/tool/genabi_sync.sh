#!/bin/sh

WORK_PATH=../
echo "work path:"$WORK_PATH

cd $WORK_PATH

rm *.go

abigen --abi ./build/contracts/ScryProtocol.abi --type ScryProtocol --pkg contract --out ScryProtocolInterface.go
abigen --abi ./build/contracts/ScryToken.abi --type ScryToken --pkg contract --out ScryTokenInterface.go

cp -f *.go ../../dots/binary/stub/contract/

echo "end."
