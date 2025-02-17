package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func StrangerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	if res, _ := regexp.MatchString("^[a-zA-Z]+$", name); !res {
		name = "dirty hacker"
	}
	fmt.Fprintf(w, fmt.Sprintf("hello %s", name))
}
