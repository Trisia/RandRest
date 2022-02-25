MD target
@echo off

SET GOARCH=386
go build -ldflags="-s -w -H windowsgui" -o "target\StandRest_386.exe"
echo [+] target\StandRest_386.exe

SET GOARCH=amd64
go build -ldflags="-s -w -H windowsgui" -o "target\StandRest.exe"
echo [+] target\StandRest.exe