#!/usr/bin/env bash

# we assume geth has been installed correctly
# and truffle.js also has been configured
# the script aims to initialize geth

echo "start initializing geth..."

# test and deploy contracts to geth
WORK_PATH=./
echo "work path:"$WORK_PATH
cd $WORK_PATH

# start geth
GETH_DATA_DIR="$HOME/gethdata/chain/"
echo $GETH_DATA_DIR

if [ ! -d $GETH_DATA_DIR ]; then
  mkdir -p $GETH_DATA_DIR
else
  rm $GETH_DATA_DIR/* -rf
fi

geth --datadir $GETH_DATA_DIR init "./genesis.json"
if [ $? -ne 0 ]; then
  echo "failed to init geth."
  exit -1
fi

nohup geth --datadir $GETH_DATA_DIR --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi eth,web3,net,personal,ssh,db,debug &
if [ $? -ne 0 ]; then
  echo "failed to start geth."
  exit -1
fi

sleep 20

# create account
geth attach "$GETH_DATA_DIR/geth.ipc" --exec 'loadScript("./create_account.js")'
if [ $? -ne 0 ]; then
  echo "failed to attach to geth."
  exit -1
fi

# wait for mining
sleep 120

truffle migrate --reset --network geth
if [ $? -ne 0 ]
then
    echo "truffle migrate failed."
    exit -1
fi


echo "Initialize geth successed."
