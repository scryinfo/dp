cd $PSScriptRoot
Remove-Item ".\resources\app\*" -Recurse

$AppName = (Get-Content "..\bundler.json" -TotalCount 2)[-1].Split(":")[1].Split("`"")[1]
echo "|-> * Name in bundle file: $AppName"
echo "|-> * Vue.js webpack prepared. "
Start-Sleep -Milliseconds 3000

npm run build
echo "|-> * Update UI files to UI directory: "
Remove-Item "..\resources\app" -Recurse
Copy-Item ".\resources\app" "..\resources" -Recurse
echo "|-> * finish." 

echo "|-> * Update UI files to bundle directory: "
Remove-Item "C:\Users\马同帅\AppData\Roaming\$AppName\resources\app" -Recurse
Copy-Item .\resources\app "C:\Users\马同帅\AppData\Roaming\$AppName\resources" -Recurse
echo "|-> * finish. "

echo "|-> * Webpack UI files finished. "
