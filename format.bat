@echo off
cd src\github.com\mikan
for /r %%i in (*.go) do ..\..\..\bin\goimports -l -w "%%i"
cd ..\..\..
