#!/bin/sh

WORK_PATH=../
echo "work path:"$WORK_PATH

cd $WORK_PATH

rm *.go

abigen --abi ./build/contracts/ScryProtocol.abi --type ScryProtocol --pkg contractinterface --out ScryProtocolInterface.go
abigen --abi ./build/contracts/ScryToken.abi --type ScryToken --pkg contractinterface --out ScryTokenInterface.go

cp -f build/contracts/*.abi ../testconsole/
cp -f *.go ../sdk/interface/contractinterface/

echo "end."
