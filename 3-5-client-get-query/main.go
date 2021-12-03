package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"realhttpworld"
)

func main() {

	values := url.Values{
		"query": {"hello world"},
	}

	r, err := http.Get("http://localhost:8888/" + "?" + values.Encode())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer r.Body.Close()

	realhttpworld.Write(r)
}
