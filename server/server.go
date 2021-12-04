package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	fmt.Fprintf(w, "Hello World\n")
}

func handlerCookie(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	w.Header().Add("Set-Cookie", "Visit=True")
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cookie", handlerCookie)
	http.ListenAndServe(":8888", nil)
}
