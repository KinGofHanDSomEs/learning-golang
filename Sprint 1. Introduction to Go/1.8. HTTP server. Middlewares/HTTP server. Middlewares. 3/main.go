package main

import (
	"net/http"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "Basic userid:password" {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The answer is 42"))
}
