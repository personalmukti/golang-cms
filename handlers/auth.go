package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/golang-cms/auth"
)

type AuthHandler struct {
	Users auth.UserStore
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	salt, err := auth.GenerateSalt()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hash := auth.HashPassword(req.Password, salt)
	user := &auth.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hash,
		Salt:         salt,
	}
	if err := h.Users.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.Users.GetByUsername(req.Username)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	if !auth.CheckPasswordHash(user.PasswordHash, req.Password, user.Salt) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
