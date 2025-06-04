package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yourusername/golang-cms/auth"
	"github.com/yourusername/golang-cms/handlers"
)

func main() {
	userStore := auth.NewInMemoryUserStore()
	authHandler := &handlers.AuthHandler{Users: userStore}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Golang CMS")
	})

	http.HandleFunc("/api/users/register", authHandler.Register)
	http.HandleFunc("/api/auth/login", authHandler.Login)

	http.Handle("/api/users/me", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value(auth.UserContextKey).(string)
		json.NewEncoder(w).Encode(map[string]string{"username": username})
	})))

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
