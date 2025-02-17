package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	a = 0
	b = 1
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, min(a, b))
	if a > b {
		b += a
	} else {
		a += b
	}
}

func StartServer(t time.Duration) {
	http.ListenAndServe(":8080", nil)
}
