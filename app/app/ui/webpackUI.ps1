echo "Scry Info.  All rights reserved."
echo "license that can be found in the license file."

cd $PSScriptRoot

echo "|-> * Install dependences in 'package.json' "
echo ""
Start-Sleep -Milliseconds 1000

npm install

echo ""
echo "|-> * Install dependences finished. "


Remove-Item ".\resources\app\*" -Recurse
echo "|-> * Webpack UI files prepared. "
echo ""
Start-Sleep -Milliseconds 1000

npm run build --report

echo ""
echo "|-> * Webpack UI files finished. "
echo "|-> * End. "
echo ""
Start-Sleep -Milliseconds 15000
