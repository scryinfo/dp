Write-Output "Scry Info.  All rights reserved."
Write-Output "license that can be found in the license file."

Set-Location $PSScriptRoot
Set-Location ..

Write-Output "|-> * Please compile contract and copy 'abi' files like scryprotocol.json to scryprotocol.abi first please. (common  2 files)"

Start-Sleep -Milliseconds 300

$confirm = Read-Host "|-> * Make sure you read the tip above, do you want to continue now? [Y/N] "

if ($confirm -eq "Y") {
    Write-Output "|-> * Genarate abi files: "
    if (Test-Path ".\*.go") {
        Remove-Item ".\*.go"
    }
    
    abigen --abi ".\build\contracts\ScryProtocol.abi" --type ScryProtocol --pkg contractinterface --out ScryProtocolInterface.go
    abigen --abi ".\build\contracts\ScryToken.abi" --type ScryToken --pkg contractinterface --out ScryTokenInterface.go
    Write-Output "|-> * Finish. "
    
    Write-Output "|-> * Move go interface to it's position"
    if (Test-Path "..\..\dots\binary\stub\contract\*.go") {
        Remove-Item "..\..\dots\binary\stub\contract\*.go"
    }
    
    Copy-Item ".\*.go" "..\..\dots\binary\stub\contract\"
    Remove-Item ".\*.go"
    Write-Output "|-> * Finish"
}

Write-Output "|-> * End. "

Start-Sleep -Milliseconds 5000
