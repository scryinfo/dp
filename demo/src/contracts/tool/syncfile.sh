#!/bin/sh

WORK_PATH=../
echo "work path:"$WORK_PATH

cd $WORK_PATH

cp -f build/contracts/*.abi ../testconsole/
cp -f *.go ../sdk/interface/contractinterface/

echo "end."
