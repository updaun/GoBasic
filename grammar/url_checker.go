package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	var results = make(map[string]string)
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
