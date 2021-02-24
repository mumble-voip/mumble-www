@echo off
SETLOCAL

PUSHD %~dp0

PUSHD hugo
hugo
IF ERRORLEVEL 1 (EXIT /B 1)
POPD

DEL mumble-www.7z 1>NUL
IF ERRORLEVEL 1 EXIT /B 1

7z.exe a mumble-www.7z public
IF ERRORLEVEL 1 EXIT /B 1

POPD
