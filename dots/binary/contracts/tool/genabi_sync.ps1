cd $PSScriptRoot
cd ..
pwd
echo "|-> * Please compile contract and copy 'abi' in scryprotocol.json to scryprotocol.abi first! "
Start-Sleep -Milliseconds 3000

echo "|-> * Genarate abi files: "
Remove-Item *.go
abigen --abi ".\build\contracts\ScryProtocol.abi" --type ScryProtocol --pkg contractinterface --out ScryProtocolInterface.go
abigen --abi ".\build\contracts\ScryToken.abi" --type ScryToken --pkg contractinterface --out ScryTokenInterface.go
echo "|-> * finish. "
Start-Sleep -Milliseconds 3000

echo "|-> * Copy go files: "
Copy-Item ".\*.go" "..\sdk\interface\contractinterface"
Copy-Item ".\build\contracts\*.abi" "..\testconsole"
echo "|-> * finish. "
echo "|-> * End. "
Start-Sleep -Milliseconds 15000
