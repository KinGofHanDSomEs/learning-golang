package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello, stranger!")
		return
	}
	if res, _ := regexp.MatchString("^[a-zA-Z]+$", name); !res {
		fmt.Fprintf(w, "Hello, dirty hacker!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if res, _ := regexp.MatchString("^[a-zA-Z]+$", name); !res {
			r.URL.Query().Set("name", "dirty hacker")
		}
		next.ServeHTTP(w, r)
	}

}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			r.URL.Query().Set("name", "stranger")
		}
		next.ServeHTTP(w, r)
	}
}
