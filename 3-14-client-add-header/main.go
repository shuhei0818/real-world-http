package main

import (
	"net/http"
	"realhttpworld"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8888/", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Key", "Value")
	req.SetBasicAuth("UserName", "Password")
	req.AddCookie(&http.Cookie{Name: "test", Value: "value"})

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	realhttpworld.Write(res)
}
