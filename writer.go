package realhttpworld

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Write(resp *http.Response) {
	fmt.Println("---- Request ----")

	fmt.Printf("%s %s %s\n", resp.Request.Method, resp.Request.URL.Path+"?"+resp.Request.URL.RawQuery, resp.Request.Proto)

	for k, v := range resp.Request.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ";"))
	}

	fmt.Println()

	var qb []byte
	if resp.Request.Body != nil {
		qb, _ = io.ReadAll(resp.Request.Body)
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
