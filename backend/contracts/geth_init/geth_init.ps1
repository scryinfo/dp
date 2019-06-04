cd $PSScriptRoot

Remove-Item "./chain/*" -Recurse
echo "Scry Info.  All rights reserved."
echo "license that can be found in the license file."



geth --datadir "chain" init genesis.json
geth --datadir "chain" --rpc --rpcapi "db,eth,net,web3" --nodiscover console
