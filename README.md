# M3U8

M3U8 是一个使用了 Go 语言编写的迷你 M3U8 下载工具。你只需指定必要的 flag (`u`、`o`、`c`) 来运行, 工具就会自动帮你解析 M3U8 文件，并将 TS 片段下载下来合并成一个文件。



## 功能

- 下载和解析 M3U8（仅限 VOD 类型）
- 下载 TS 失败重试
- 解析 Master playlist
- 解密 TS
- 合并 TS 片段
- 下载链接和加密 key 的获取链接都支持携带 header 信息 (2020-07-27)

## 用法

### 源码方式

指定下载文件并下载到指定目录
```bash
go run main.go -u=http://example.com/index.m3u8 -o=data/example
```

指定下载文件并下载到当前目录下的out目录中
```bash
go run main.go -u=http://example.com/index.m3u8
```

读取文件中的链接批量下载下载到指定目录
```bash
go run main.go -f=data/test.txt -o=data/example
```

读取文件中的链接批量下载下载到当前目录下的out目录中
```bash
go run main.go -f=data/test.txt
```
### 二进制方式:

Linux 和 MacOS

```
./m3u8 -u=http://example.com/index.m3u8 -o=data/example
```

Windows PowerShell

```
.\m3u8.exe -u="http://example.com/index.m3u8" -o="D:\data\example"
```

参数说明：

```
-c int
    下载协程并发数 (default 25)
-f string
    M3U8 URL and Headers files, url 和 file 必须至少有一个 (default "data\\test.txt")
-h    查看命令帮助
-headers string
    请求需要携带的头信息, 若 file 中也有 Headers,则此 Headers 会附加在 file 中的 Headers 后面
-o string
    文件保存目录, 默认输出到当前目录下的 out 目录中 (default "./out")
-u string
    M3U8 地址, url 和 file 必须至少有一个
```

部分链接可能限制请求频率，可根据实际情况调整 `c` 参数的值。

## 下载

[二进制文件](https://github.com/sudot/m3u8/releases)

## 编译
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o m3u8_v1.2_windows_amd64.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o m3u8_v1.2_darwin_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o m3u8_v1.2_linux_amd64
```

## 截屏

![Demo](./screenshots/demo.gif)

## 参考资料

- [TS科普 2 包头](https://blog.csdn.net/cabbage2008/article/details/49281729)
- [HTTP Live Streaming draft-pantos-http-live-streaming-23](https://tools.ietf.org/html/draft-pantos-http-live-streaming-23#section-4.3.4.2)
- [MPEG transport stream - Wikipedia](https://en.wikipedia.org/wiki/MPEG_transport_stream)


## License

[MIT License](./LICENSE)