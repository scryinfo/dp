#!/bin/sh

WORK_PATH=../
echo "work path:"$WORK_PATH

cd $WORK_PATH

cp build/contracts/*.abi ../testconsole/
cp *.go ../sdk/contractinterface/

echo "end."
