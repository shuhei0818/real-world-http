package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, strings.Join(v, ";"))
	}

	fmt.Fprintf(w, "\n")

	rc := r.Body
	defer rc.Close()
	b, _ := io.ReadAll(rc)

	if b != nil {
		fmt.Fprintf(w, "%s\n", string(b))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
