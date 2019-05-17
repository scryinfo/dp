cd $PSScriptRoot

echo "|-> * npm install prepared. "
Start-Sleep -Milliseconds 1000

npm install zeppelin-solidity
npm install ethereumjs-wallet
npm install -g truffle

echo "|-> * npm install finish. "
Start-Sleep -Milliseconds 15000
