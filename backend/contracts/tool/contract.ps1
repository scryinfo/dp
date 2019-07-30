Write-Output "Scry Info.  All rights reserved."
Write-Output "license that can be found in the license file."

Set-Location $PSScriptRoot
Set-Location ..

Write-Output "|-> * npm install prepared. "
Write-Output ""
Start-Sleep -Milliseconds 1000

npm install zeppelin-solidity
npm install -g truffle@4.1.14

Write-Output "|-> * npm install finished. "
Write-Output ""

truffle version
Write-Output "|-> * truffle migrate prepared. "
Write-Output ""
Start-Sleep -Milliseconds 1000

Start-Transcript "migrate.log" -Force
truffle migrate --network geth
Stop-Transcript 
Write-Output "|-> * truffle migrate finished. "
Write-Output "|-> * End. "

Start-Sleep -Milliseconds 5000
