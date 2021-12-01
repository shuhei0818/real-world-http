package main

import (
	"fmt"
	"net/http"
	"os"
	"realhttpworld"
)

func main() {
	res, err := http.Get("http://localhost:8888/")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer res.Body.Close()

	realhttpworld.Write(res)

}
