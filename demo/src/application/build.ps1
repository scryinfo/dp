$WorkPath = $PSScriptRoot
cd $WorkPath
Remove-Item .\UI\resources\app\* -Recurse

$AppName = (Get-Content ".\bundler.json" -TotalCount 2)[-1].Split(":")[1].Split("`"")[1]
echo " * Name in bundle file: " $AppName

cd .\UI
echo " * Vue.js build prepared. "
Start-Sleep -Milliseconds 2000

npm run build
echo " * Copy UI files to UI directory: "
Remove-Item ..\resources\app -Recurse
Copy-Item .\resources\app ..\resources -Recurse
echo " * finish." 
echo " * Copy UI files to bundle directory. "
Remove-Item "C:\Users\马同帅\AppData\Roaming\$AppName\resources\app" -Recurse
Copy-Item .\resources\app "C:\Users\马同帅\AppData\Roaming\$AppName\resources" -Recurse
echo " * finish. "

cd ..
echo " * Go-astilectron bundle prepared. " 
Start-Sleep -Milliseconds 2000

astilectron-bundler -v
Start-Sleep -Milliseconds 5000
