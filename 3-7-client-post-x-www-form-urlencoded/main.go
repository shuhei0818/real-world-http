package main

import (
	"net/http"
	"net/url"
	"realhttpworld"
)

func main() {

	values := url.Values{
		"test": {"hello world"},
	}

	// PostForm RFC3986でエンコード
	r, err := http.PostForm("http://localhost:8888/", values)
	if err != nil {
		panic(1)
	}

	realhttpworld.Write(r)
}
