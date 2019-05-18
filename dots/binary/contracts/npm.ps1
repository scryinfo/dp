cd $PSScriptRoot

echo "|-> * npm install prepared. "
echo ""
Start-Sleep -Milliseconds 1000

npm install zeppelin-solidity
npm install ethereumjs-wallet
npm install -g truffle@4.1.14

echo "|-> * npm install finished. "
echo ""
Start-Sleep -Milliseconds 15000
