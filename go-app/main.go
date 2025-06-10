package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=postgres dbname=message_board password=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Server is running")

	initDB()
	defer db.Close()

	r := mux.NewRouter()

	// authorization routes
	r.HandleFunc("/api/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/api/login", LoginHandler).Methods("POST")
	r.HandleFunc("/api/logout", LogoutHandler).Methods("POST")
	r.HandleFunc("/api/me", MeHandler).Methods("GET")

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT name, message FROM messages ORDER BY id DESC")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var messages []Message
		for rows.Next() {
			var m Message
			if err := rows.Scan(&m.Name, &m.Message); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			messages = append(messages, m)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	})

	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			msg := r.FormValue("message")

			_, err := db.Exec("INSERT INTO messages (name, message) VALUES ($1, $2)", name, msg)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Write([]byte("OK"))
		}
	})

	http.ListenAndServe(":8080", r)
}
