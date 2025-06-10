package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var messages = []Message{
	{"Alice", "Hello!"},
	{"Bob", "Hi!"},
}

func main() {
	fmt.Println("Server is running on http://localhost:8080")

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	})

	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			msg := r.FormValue("message")
			messages = append(messages, Message{name, msg})
			w.Write([]byte("OK"))
		}
	})

	http.ListenAndServe(":8080", nil)
}
