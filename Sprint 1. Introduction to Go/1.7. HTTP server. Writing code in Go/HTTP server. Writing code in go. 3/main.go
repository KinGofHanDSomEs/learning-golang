package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	k = 0
	a = 0
	b = 1
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, min(a, b))
	k++
	if a > b {
		b += a
	} else {
		a += b
	}
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "rpc_duration_milliseconds_count ", k)
}

func StartServer(t time.Duration) {
	http.ListenAndServe(":8080", nil)
}