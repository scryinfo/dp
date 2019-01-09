#!/bin/sh

WORK_PATH=../
echo "work path:"$WORK_PATH

cd $WORK_PATH

cp build/contracts/ScryProtocol.abi ../testconsole/
cp ScryProtocolInterface.go ../sdk/contractinterface/

echo "end."
