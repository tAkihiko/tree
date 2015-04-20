@echo off
set dist=tree
pushd %~dp0
call :Build windows amd64 %dist%.exe
call :Build windows 386 %dist%.exe
pause
exit

:Build
set GOOS=%~1
set GOARCH=%~2
set out_dir=%dist%\%GOOS%\%GOARCH%
if not exist %out_dir%\ ( mkdir %out_dir% )
echo GOOS=%GOOS% GOARCH=%GOARCH% build
go build -o %out_dir%\%~3
exit /b
