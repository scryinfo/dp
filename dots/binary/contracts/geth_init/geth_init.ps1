cd $PSScriptRoot

Remove-Item "./chain/*" -Recurse

geth --datadir "chain" init genesis.json
geth --datadir "chain" --rpc --rpcapi "db,eth,net,web3" --nodiscover console
