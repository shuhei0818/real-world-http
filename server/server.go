package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	fmt.Fprintf(w, "Defalut\n")
}

func handlerCookie(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	w.Header().Add("Set-Cookie", "Visit=True")
	fmt.Fprintf(w, "Cookie\n")
}

func handlerSlow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "Slow\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cookie", handlerCookie)
	http.HandleFunc("/slow", handlerSlow)
	http.ListenAndServe(":8888", nil)
}
