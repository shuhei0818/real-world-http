package realhttpworld

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Write(resp *http.Response) {
	fmt.Println("---- Request ----")

	path := resp.Request.URL.Path
	if len(resp.Request.URL.RawQuery) > 0 {
		path += "?" + resp.Request.URL.RawQuery
	}

	fmt.Printf("%s %s %s\n", resp.Request.Method, path, resp.Request.Proto)

	for k, v := range resp.Request.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ";"))
	}

	fmt.Println()

	var qb []byte
	if resp.Request.Body != nil {
		b, _ := resp.Request.GetBody()
		qb, _ = io.ReadAll(b)
	}

	fmt.Println(string(qb))

	fmt.Println("---- Response ----")

	fmt.Printf("%s %s", resp.Proto, resp.Status)

	fmt.Println()

	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ";"))
	}

	fmt.Println()

	b, _ := io.ReadAll(resp.Body)

	fmt.Println(string(b))
}
