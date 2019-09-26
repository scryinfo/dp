Write-Output "Scry Info.  All rights reserved."
Write-Output "license that can be found in the license file."

Set-Location $PSScriptRoot
Set-Location ..

Write-Output "|-> * Next step is 'npm install', if you have already install successfully one time, you can skip it and run webpack directly.  "

Start-Sleep -Milliseconds 300

$confirm = Read-Host "|-> * Make sure you read the tip above, do you want to install? [Y/N] "

if ($confirm -eq "Y") {
    Write-Output "|-> * Install dependences in 'package.json' "
    Write-Output ""
    Start-Sleep -Milliseconds 1000

    npm install

    Write-Output "|-> * Install dependences finished. "
    Write-Output ""
}

if (Test-Path ".\resources\app\*") {
    Remove-Item ".\resources\app\*" -Recurse
}

Write-Output "|-> * Webpack UI files prepared. "
Write-Output ""
Start-Sleep -Milliseconds 1000

npm run build --report

Write-Output "|-> * Webpack UI files finished. "
Write-Output "|-> * End. "
Write-Output ""

Start-Sleep -Milliseconds 5000
