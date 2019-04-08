cd $PSScriptRoot
$Command = Read-Host "What do you want to do? (wb: webpack UI and bundle Go with Js, b: bundle Go with Js only)"
if ($Command -eq "wb") {& ".\UI\webpackUI.ps1"}
echo "|-> * ------------------------------"

cd $PSScriptRoot
echo "|-> * Go-astilectron bundle prepared. " 
Start-Sleep -Milliseconds 3000

astilectron-bundler -v
echo "|-> * Go-astilectron bundle finished. "
Start-Sleep -Milliseconds 20000
