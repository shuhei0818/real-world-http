// GET, HEAD, POST 以外のメソッドを利用する場合は、http.Request / http.Client を利用する
// Bodyはurl.Valuesを利用する
package main

import (
	"net/http"
	"net/url"
	"realhttpworld"
	"strings"
)

func main() {
	values := url.Values{
		"key": {"value"},
	}
	c := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8888/", strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	realhttpworld.Write(res)
}
