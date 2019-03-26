$WorkPath = $PSScriptRoot
cd $WorkPath
Remove-Item .\UI\resources\app\* -Recurse

cd .\UI
echo "Vue.js build prepared. "
Start-Sleep -Milliseconds 2000

npm run build
Remove-Item ..\resources\app -Recurse
Copy-Item .\resources\app ..\resources -Recurse
Remove-Item "C:\Users\马同帅\AppData\Roaming\My Astilectron demo\resources\app" -Recurse
Copy-Item .\resources\app "C:\Users\马同帅\AppData\Roaming\My Astilectron demo\resources" -Recurse

cd ..
echo "Go-astilectron bundle prepared. "
Start-Sleep -Milliseconds 2000

astilectron-bundler -v
Start-Sleep -Milliseconds 5000