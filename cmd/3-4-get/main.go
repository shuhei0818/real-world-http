package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	res, err := http.Get("http://localhost:8888")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s %s", res.Proto, res.Status)

	fmt.Println()

	for k, v := range res.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ";"))
	}

	fmt.Println()

	fmt.Println(string(b))

}
