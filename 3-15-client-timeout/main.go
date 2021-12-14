package main

import (
	"context"
	"net/http"
	"realhttpworld"
	"time"
)

func main() {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8888/slw", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	realhttpworld.Write(res)
}
