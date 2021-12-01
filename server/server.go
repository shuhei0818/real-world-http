package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
