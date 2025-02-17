package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Greeting struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type RPCprotocol struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	jsonBytes, err := json.Marshal(RPCprotocol{Status: "ok", Result: Greeting{Greetings: "hello", Name: name}})
	if err != nil {
		http.Error(w, "invalid translate to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(jsonBytes)
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if res, _ := regexp.MatchString("^[a-zA-Z]+$", name); !res {
			panic("invalid name")
		}
		next.ServeHTTP(w, r)
	}

}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.Header().Set("Content-type", "application/json")
			jsonBytes, err := json.Marshal(RPCprotocol{Status: "ok", Result: Greeting{Greetings: "hello", Name: "stranger"}})
			if err != nil {
				http.Error(w, "invalid translate to json", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-type", "application/json")
			w.Write(jsonBytes)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil { //{"status": "error", "result": {}}
				jsonBytes, err := json.Marshal(RPCprotocol{Status: "error", Result: make(map[string]interface{})})
				if err != nil {
					http.Error(w, "panic", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-type", "application/json")
				w.Write(jsonBytes)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
}
