Write-Output "Scry Info.  All rights reserved."
Write-Output "license that can be found in the license file."

Start-Sleep -Milliseconds 1000

geth attach "ipc:\\.\pipe\geth.ipc" --exec "loadScript('./create_account.js')"
