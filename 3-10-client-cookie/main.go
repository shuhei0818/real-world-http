package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"realhttpworld"
	"strconv"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Jar: jar,
	}
	for i := 0; i < 2; i++ {
		resp, err := client.Get("http://localhost:8888/cookie")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println(strconv.Itoa(i + 1))
		realhttpworld.Write(resp)
	}
}
