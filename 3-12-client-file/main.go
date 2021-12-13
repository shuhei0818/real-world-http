package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))

	c := http.Client{
		Transport: t,
	}
	resp, err := c.Get("file://./3-12-client-file/main.go")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
