package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("secret-key-change-this"))

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode, // ← Cookie制限を緩和
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Password hashing error", 500)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", u.Username, string(hash))
	if err != nil {
		http.Error(w, "Username already taken?", 500)
		return
	}

	w.Write([]byte("registered"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	var hashed string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username=$1", u.Username).Scan(&hashed)
	if err != nil {
		http.Error(w, "User not found", 400)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(u.Password)) != nil {
		http.Error(w, "Invalid password", 400)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["username"] = u.Username
	session.Save(r, w)

	w.Write([]byte("logged in"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.Write([]byte("logged out"))
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username, ok := session.Values["username"].(string)
	if !ok {
		http.Error(w, "Not logged in", 401)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"username": username})
}
