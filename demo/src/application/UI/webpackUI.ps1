cd $PSScriptRoot
Remove-Item ".\resources\app\*" -Recurse
echo "|-> * Vue.js webpack prepared. "
Start-Sleep -Milliseconds 3000

npm run build
echo "|-> * Update UI files to UI directory: "
Remove-Item "..\resources\app" -Recurse
Copy-Item ".\resources\app" "..\resources" -Recurse
echo "|-> * finish." 

echo "|-> * Webpack UI files finished. "
echo "|-> * End. "
Start-Sleep -Milliseconds 15000
