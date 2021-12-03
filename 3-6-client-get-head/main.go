package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, err := http.Head("http://localhost:8888")
	if err != nil {
		panic(1)
	}
	fmt.Println(r.Status)
	fmt.Println(r.Header)
}
