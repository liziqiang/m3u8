package parse

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sudot/m3u8/tool"
)

type Result struct {
	URL  *url.URL
	M3u8 *M3u8
	Keys map[int]string
}

func FromURL(link string, header http.Header) (*Result, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	link = u.String()
	body, err := tool.Get(link, header)
	if err != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %s", err.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer body.Close()
	m3u8, err := parse(body)
	if err != nil {
		return nil, err
	}
	if len(m3u8.MasterPlaylist) != 0 {
		sf := m3u8.MasterPlaylist[0]
		return FromURL(tool.ResolveURL(u, sf.URI), header)
	}
	if len(m3u8.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}
	result := &Result{
		URL:  u,
		M3u8: m3u8,
		Keys: make(map[int]string),
	}

	urlKeyMap := make(map[string]string)
	for idx, key := range m3u8.Keys {
		switch {
		case key.Method == "" || key.Method == CryptMethodNONE:
			continue
		case key.Method == CryptMethodAES:
			// Request URL to extract decryption key
			keyURL := key.URI
			keyURL = tool.ResolveURL(u, keyURL)
			value := urlKeyMap[keyURL]
			if value == "" {
				resp, err := tool.Get(keyURL, header)
				if err != nil {
					return nil, fmt.Errorf("extract key failed: %s", err.Error())
				}
				keyByte, err := ioutil.ReadAll(resp)
				_ = resp.Close()
				if err != nil {
					return nil, err
				}
				value = string(keyByte)
				urlKeyMap[keyURL] = value
				fmt.Println("decryption key: ", value)
			}
			result.Keys[idx] = value
		default:
			return nil, fmt.Errorf("unknown or unsupported cryption method: %s", key.Method)
		}
	}
	return result, nil
}

func PickURL(context string) []string {
	lines := strings.Split(context, "\n")
	urls := make([]string, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "http") {
			urls = append(urls, strings.TrimSpace(line))
		}
	}
	return urls
}

func PickHeader(context string) http.Header {
	lines := strings.Split(context, "\n")
	header := make(http.Header)
	for _, line := range lines {
		if strings.HasPrefix(line, "http") {
			continue
		}
		index := strings.Index(line, ":")
		if index < 1 {
			// header第一位字符是冒号,或者不包含冒号做忽略处理
			continue
		}
		key := strings.TrimSpace(line[:index])
		value := strings.TrimSpace(line[index+1:])
		header.Add(key, value)
	}
	return header
}
