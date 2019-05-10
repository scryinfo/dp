cd $PSScriptRoot
Remove-Item ".\resources\app\*" -Recurse
echo "|-> * Webpack UI files prepared. "
Start-Sleep -Milliseconds 3000

npm run build

echo "|-> * Webpack UI files finished. "
echo "|-> * End. "
Start-Sleep -Milliseconds 15000
