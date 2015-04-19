@echo off
pushd %~dp0
call :Build windows amd64 tree.exe
call :Build windows 386 tree.exe
pause
exit

:Build
set GOOS=%~1
set GOARCH=%~2
set out_dir=build\%GOOS%\%GOARCH%
if not exist %out_dir%\ ( mkdir %out_dir% )
echo GOOS=%GOOS% GOARCH=%GOARCH% build
go build -o %out_dir%\%~3
exit /b
