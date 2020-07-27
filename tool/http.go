package tool

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string, header http.Header) (io.ReadCloser, error) {
	c := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	request, err := http.NewRequest("GET", url, nil)
	if header != nil {
		request.Header = header
	}
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
