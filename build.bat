@echo off
SETLOCAL

PUSHD %~dp0

PUSHD hugo
hugo
IF ERRORLEVEL 1 (PAUSE && EXIT /B 1)
POPD

go build src/mumble-www.go src/githubcache.go src/snapshotcache.go
IF ERRORLEVEL 1 (PAUSE && EXIT /B 1)

SET GOOS=linux
go build src/mumble-www.go src/githubcache.go src/snapshotcache.go
IF ERRORLEVEL 1 (PAUSE && EXIT /B 1)

DEL mumble-www.7z 1>NUL
7z.exe a mumble-www.7z public mumble-www

POPD
