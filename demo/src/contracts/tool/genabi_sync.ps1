$WorkPath = $PSScriptRoot
cd $WorkPath
cd ..
pwd
echo "Please compile contract and copy 'abi' in scryprotocol.json to scryprotocol.abi first! "
Start-Sleep -Milliseconds 2000

Remove-Item *.go
abigen --abi ".\build\contracts\ScryProtocol.abi" --type ScryProtocol --pkg contractinterface --out ScryProtocolInterface.go
abigen --abi ".\build\contracts\ScryToken.abi" --type ScryToken --pkg contractinterface --out ScryTokenInterface.go
echo "Genarate abi files finish. "
Start-Sleep -Milliseconds 2000

Copy-Item ".\*.go" "..\sdk\interface\contractinterface"
Copy-Item ".\build\contracts\*.abi" "..\testconsole"
echo "Copy go files finish. "
Start-Sleep -Milliseconds 5000
