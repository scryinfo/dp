Start-Sleep -Milliseconds 5000

geth attach "ipc:\\.\pipe\geth.ipc" --exec "loadScript('./create_account.js')"