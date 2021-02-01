package main

import (
	"flag"
	"fmt"
	"github.com/sudot/m3u8/dl"
	"github.com/sudot/m3u8/parse"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	url      string
	file     string
	headers  string
	output   string
	chanSize int
	help     bool
)

func init() {
	flag.BoolVar(&help, "h", false, "查看命令帮助")
	flag.StringVar(&url, "u", "", "M3U8 地址, url 和 file 必须至少有一个")
	flag.StringVar(&file, "f", "", "M3U8 URL and Headers files, url 和 file 必须至少有一个")
	flag.StringVar(&headers, "headers", "", "请求需要携带的头信息, 若 file 中也有 Headers,则此 Headers 会附加在 file 中的 Headers 后面")
	flag.IntVar(&chanSize, "c", 25, "下载协程并发数")
	flag.StringVar(&output, "o", "./out", "文件保存目录, 默认输出到当前目录下的 out 目录中")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	if url == "" && file == "" {
		panicParameter("u OR f")
	}
	if output == "" {
		panicParameter("o")
	}
	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}

	var lines string
	if file != "" {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			panic("file '" + file + "' read error : " + err.Error())
		}
		lines = string(bytes)
	} else if url != "" {
		lines += url
	}
	if headers != "" {
		lines += headers
	}
	links := parse.PickURL(lines)
	header := parse.PickHeader(lines)

	var splitCompile = regexp.MustCompile("[\\s|\\t]+")
	for _, link := range links {
		if len(link) <= 4 {
			continue
		}

		split := splitCompile.Split(link, -1)
		url := split[0]
		fileName := strings.Join(split[1:], "")
		downloader, err := dl.NewTask(output, url, fileName, header)
		if err != nil {
			panic(err)
		}
		if err := downloader.Start(chanSize, header); err != nil {
			panic(err)
		}
	}
	fmt.Println("Done!")
}

func panicParameter(name string) {
	panic("parameter '" + name + "' is required")
}
