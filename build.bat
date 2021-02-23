@ECHO OFF
REM ���� windows �汾
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o m3u8_windows_amd64.exe

REM ���� linux �汾
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o m3u8_linux_amd64

REM ���� mac �汾
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o m3u8_darwin_amd64
