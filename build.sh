# 编译 windows 版本
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o m3u8_windows_amd64.exe
# 编译 linux 版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o m3u8_linux_amd64
# 编译 mac 版本
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o m3u8_darwin_amd64
