package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

type FillResponse struct {
	Fill int `json:"fill"`
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	file, err := os.Open("state.cfg")
	defer file.Close()
	if err != nil {
		http.Error(w, "file opening error", 500)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading file", 500)
		return
	}
	if len(data) == 0 {
		http.Error(w, "file is empty", 500)
		return
	}
	fill := ""
	for _, val := range data {
		num := string(val)
		_, err := strconv.Atoi(num)
		if err != nil {
			break
		}
		fill += num
		if len(fill) > 3 {
			http.Error(w, "invalid fill value", 500)
			return
		}
	}
	a, err := strconv.Atoi(fill)
	if err != nil || a > 100 {
		http.Error(w, "invalid fill value", 500)
		return
	}
	if err = json.NewEncoder(w).Encode(&FillResponse{a}); err != nil {
		http.Error(w, "error writing json in response", 500)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/reset", resetHandler)
	http.ListenAndServe(":8081", mux)
}
