cd $PSScriptRoot
cd ..
pwd
echo "|-> * Please compile contract and copy 'abi' files like scryprotocol.json to scryprotocol.abi first! (common  2 files)"
Start-Sleep -Milliseconds 3000

echo "|-> * Genarate abi files: "
Remove-Item *.go
abigen --abi ".\build\contracts\ScryProtocol.abi" --type ScryProtocol --pkg contractinterface --out ScryProtocolInterface.go
abigen --abi ".\build\contracts\ScryToken.abi" --type ScryToken --pkg contractinterface --out ScryTokenInterface.go
echo "|-> * Finish. "

echo "|-> * Move go interface to it's position"
Remove-Item "..\..\dots\binary\stub\contract\*.go"
Copy-Item "*.go" "..\..\dots\binary\stub\contract\"
Remove-Item *.go
echo "|-> * Finish"

echo "|-> * End. "
Start-Sleep -Milliseconds 15000
