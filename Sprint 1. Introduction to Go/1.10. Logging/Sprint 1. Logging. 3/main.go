package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Answer struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	jsonBytes, err := json.Marshal(Answer{Name: name})
	if err != nil {
		log.Printf("invalid json marshall: %v", err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	log.Print(string(jsonBytes))
	w.Write(jsonBytes)
}
