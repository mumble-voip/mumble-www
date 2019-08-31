@echo off
SETLOCAL

PUSHD %~dp0

PUSHD hugo
hugo
IF ERRORLEVEL 1 (EXIT /B 1)
POPD

call build-go.bat
IF ERRORLEVEL 1 (EXIT /B 1)

DEL mumble-www.7z 1>NUL
IF ERRORLEVEL 1 EXIT /B 1

7z.exe a mumble-www.7z public mumble-www
IF ERRORLEVEL 1 EXIT /B 1

POPD
