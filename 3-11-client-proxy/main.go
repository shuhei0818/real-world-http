package main

import (
	"net/http"
	"net/url"
	"realhttpworld"
)

func main() {
	url, err := url.Parse("http://localhost:8888")
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(url),
		},
	}
	resp, err := client.Get("http://github.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	realhttpworld.Write(resp)

}
