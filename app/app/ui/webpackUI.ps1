cd $PSScriptRoot

echo "|-> * Install dependences in 'package.json' "
Start-Sleep -Milliseconds 1000

npm install

echo "|-> * Install dependences finished. "


Remove-Item ".\resources\app\*" -Recurse
echo "|-> * Webpack UI files prepared. "
Start-Sleep -Milliseconds 1000

npm run build --report

echo "|-> * Webpack UI files finished. "
echo "|-> * End. "
Start-Sleep -Milliseconds 15000
