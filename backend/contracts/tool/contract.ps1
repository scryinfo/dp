Write-Output "Scry Info.  All rights reserved."
Write-Output "license that can be found in the license file."

Set-Location $PSScriptRoot
Set-Location ..

Write-Output "|-> * Next step is 'npm install', if you are already install successfully one time, you can skip it.  "

Start-Sleep -Milliseconds 300

$confirm = Read-Host "|-> * Make sure you read the tip above, do you want to install? [Y/N] "

if ($confirm -eq "Y") { # no matter if the character is big-case or small-case
    Write-Output "|-> * npm install prepared. "
    Write-Output ""
    Start-Sleep -Milliseconds 1000

    npm install zeppelin-solidity
    npm install -g truffle@4.1.14

    Write-Output "|-> * npm install finished. "
    Write-Output ""
}

truffle version
Write-Output "|-> * truffle migrate prepared. "
Write-Output ""
Start-Sleep -Milliseconds 1000

$Transcript = $PSScriptRoot+"/migrate.txt"
Start-Transcript -Path $Transcript -Force
truffle migrate --network geth
Stop-Transcript

Set-Location $PSScriptRoot
Write-Output "|-> * truffle migrate finished. "
Write-Output "|-> * End. "
