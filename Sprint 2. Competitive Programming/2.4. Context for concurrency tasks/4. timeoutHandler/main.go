package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	http.HandleFunc("/readSource", func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{}
		timedHandler := http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp, err := client.Get("http://localhost:8081/provideData")
			if err != nil {
				http.Error(w, "Failed to reach source server", http.StatusServiceUnavailable)
				return
			}
			defer resp.Body.Close()
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, "Failed to read response", http.StatusInternalServerError)
				return
			}
			w.Write(data)
		}), maxTimeout, "Timeout waiting for data")
		timedHandler.ServeHTTP(w, r)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func main() {
	maxTimeout := 2 * time.Second
	StartServer(maxTimeout)
}
