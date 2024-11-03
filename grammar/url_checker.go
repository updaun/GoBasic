package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.daum.net",
		"https://www.nate.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
		"https://www.yandex.com",
		"https://www.baidu.com",
		"https://www.sogou.com",
	}
	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
