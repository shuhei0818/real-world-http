package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rtime, err := time.Parse(time.RFC1123, r.Header.Get("If-Modified-Since"))
	now := time.Now()
	if err != nil {
		fmt.Println("first")
		w.Header().Add("Last-Modified", now.Format(time.RFC1123))
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", "receive first request.\n")
		return
	}

	if rtime.Before(now) {
		fmt.Println("304")
		w.WriteHeader(http.StatusNotModified)
		return
	} else {
		fmt.Println("200")
		w.Header().Add("Last-Modified", now.Format(time.RFC1123))
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", "receive new request.\n")
		return
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
