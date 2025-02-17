package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type FillRequest struct {
	Fill int `json:"fill"`
}

func stateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	var request FillRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	defer r.Body.Close()
	fill := request.Fill
	if fill < 0 || fill > 100 {
		http.Error(w, "invalid fill", 400)
		return
	}
	file, err := os.Create("state.cfg")
	defer file.Close()
	if err != nil {
		http.Error(w, "file creation error", 500)
		return
	}
	file.WriteString(fmt.Sprint(fill) + "%")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/setstate", stateHandler)
	http.ListenAndServe(":8081", mux)
}
