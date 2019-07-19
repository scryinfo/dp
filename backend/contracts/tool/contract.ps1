echo "Scry Info.  All rights reserved."
echo "license that can be found in the license file."

cd $PSScriptRoot
cd ..

echo "|-> * npm install prepared. "
echo ""
Start-Sleep -Milliseconds 1000

npm install zeppelin-solidity
npm install -g truffle@4.1.14

echo "|-> * npm install finished. "
echo ""

truffle version
echo "|-> * truffle migrate prepared. "
echo ""
Start-Sleep -Milliseconds 1000

Start-Transcript "migrate.log" -Force
truffle migrate --network geth
Stop-Transcript 
echo "|-> * truffle migrate finished. "
echo "|-> * End. "

Start-Sleep -Milliseconds 15000
