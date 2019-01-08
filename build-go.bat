@echo off
SETLOCAL

PUSHD %~dp0

set GO111MODULE=on

PUSHD src

SET SRCFILES=mumble-www.go githubcache.go snapshotcache.go config.go downloads.go programflags.go

go build -o ..\mumble-www.exe %SRCFILES%
IF ERRORLEVEL 1 (EXIT /B 1)

SET GOOS=linux
go build -o ..\mumble-www %SRCFILES%
IF ERRORLEVEL 1 (EXIT /B 1)

POPD

POPD
