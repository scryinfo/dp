Write-Output "Scry Info.  All rights reserved."
Write-Output "license that can be found in the license file."

Set-Location $PSScriptRoot

Write-Output "|-> * Modify contract address, make sure you have copy it into ./config.txt first please. "

Start-Sleep -Milliseconds 300

$confirm = Read-Host "|-> * Make sure you read the tip above, do you want to continue now? [Y/N] "

if ($confirm -eq "Y") { # no matter if the character is big-case or small-case
    $tokenAddr = (Get-Content .\config.txt -TotalCount 1).Split(": ")[3]
    $protocolAddr = (Get-Content .\config.txt -TotalCount 2)[-1].Split(": ")[3]

    $dirs = Get-ChildItem "..\users" -Recurse *.json    # dirs type: System.IO.DirectoryInfo

    Write-Output ""

    $dirs | ForEach-Object -Process {    # for each type: System.IO.FileInfo
        Set-Location $_.DirectoryName

        if (Test-Path -Path ".\main.exe") {
            Remove-Item ".\main.exe" -Force
        }
        # go build ".\main.go"

        $conf = Get-Content $_.FullName | ConvertFrom-Json
        $conf.dots[0].lives[0].json.tokenContractAddr = $tokenAddr.Replace("`"", "")
        $conf.dots[0].lives[0].json.protocolContractAddr = $protocolAddr.Replace("`"", "")
        $conf = $conf | ConvertTo-Json -Depth 5
        Set-Content -path $_.FullName -Value $conf -Force

        $t = $_.DirectoryName
        Write-Output "|-> * $t finished. "
    }

    Set-Location $PSScriptRoot
    go build ".\main.go"

    Write-Output ""
    Write-Output "|-> * Generate executeable file finished. "
    Write-Output "|-> * Modify config file finished. "
    Write-Output "|-> * End. "
}

Start-Sleep -Milliseconds 5000
