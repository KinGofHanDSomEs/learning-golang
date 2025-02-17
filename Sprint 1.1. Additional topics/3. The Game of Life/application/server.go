package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type FillRequest struct {
	Fill int `json:"fill"`
}

func StateHandler(w http.ResponseWriter, r *http.Request) {
	var request FillRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid json request", 400)
		return
	}
	defer r.Body.Close()
	fill := request.Fill
	if fill < 0 || fill > 100 {
		http.Error(w, "invalid value of fill", 400)
		return
	}
	file, err := os.Create("state")
	if err != nil {
		http.Error(w, "error opening of file", 500)
		return
	}
	defer file.Close()
	file.WriteString(fmt.Sprint(fill) + " relevant")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func RunServer(port string) { // curl -L 'http://localhost:8081/setstate' --header 'Content-Type: application/json' --data '{"fill":40}'
	mux := http.NewServeMux()
	mux.HandleFunc("/setstate", StateHandler)
	http.ListenAndServe(":"+port, mux)
}
