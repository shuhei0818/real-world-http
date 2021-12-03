package main

import (
	"net/http"
	"realhttpworld"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")
	resp, err := http.Post("http://localhost:8888/", "text/plain", r)
	if err != nil {
		panic(1)
	}
	defer resp.Body.Close()

	realhttpworld.Write(resp)

}
