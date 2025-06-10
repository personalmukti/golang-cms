package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
)

// Permission represents a named permission.
type Permission string

// Role holds a set of permissions.
type Role struct {
	Name        string
	Permissions []Permission
}

// User represents an application user.
type User struct {
	Username string
	Password string
	Email    string
	Roles    []string
}

var (
	mu     sync.Mutex
	users  = make(map[string]*User)
	roles  = make(map[string]*Role)
	tokens = make(map[string]string)
)

// generateToken creates a random hex token.
func generateToken() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// Register adds a new user.
func Register(username, password, email string) error {
	mu.Lock()
	defer mu.Unlock()
	if username == "" || password == "" {
		return errors.New("username and password required")
	}
	if _, exists := users[username]; exists {
		return errors.New("user already exists")
	}
	users[username] = &User{Username: username, Password: password, Email: email}
	return nil
}

// Login validates credentials and returns a token.
func Login(username, password string) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	u, ok := users[username]
	if !ok || u.Password != password {
		return "", errors.New("invalid credentials")
	}
	token, err := generateToken()
	if err != nil {
		return "", err
	}
	tokens[token] = username
	return token, nil
}

// ForgotPassword is a stub that always succeeds.
func ForgotPassword(email string) error {
	if email == "" {
		return errors.New("email required")
	}
	// In a real application, send an email here.
	return nil
}

// ResetPassword updates the user's password.
func ResetPassword(username, newPassword string) error {
	mu.Lock()
	defer mu.Unlock()
	u, ok := users[username]
	if !ok {
		return errors.New("user not found")
	}
	u.Password = newPassword
	return nil
}

// LoginByGoogle is a stub for Google OAuth login.
func LoginByGoogle(googleToken string) (string, error) {
	if googleToken == "" {
		return "", errors.New("token required")
	}
	// Assume googleToken is the username for demo purposes.
	username := googleToken
	if _, exists := users[username]; !exists {
		users[username] = &User{Username: username}
	}
	token, err := generateToken()
	if err != nil {
		return "", err
	}
	tokens[token] = username
	return token, nil
}

// AddRole creates a role with permissions.
func AddRole(name string, perms []Permission) {
	mu.Lock()
	defer mu.Unlock()
	roles[name] = &Role{Name: name, Permissions: perms}
}

// AssignRole assigns a role to a user.
func AssignRole(username, roleName string) error {
	mu.Lock()
	defer mu.Unlock()
	u, ok := users[username]
	if !ok {
		return errors.New("user not found")
	}
	if _, ok := roles[roleName]; !ok {
		return errors.New("role not found")
	}
	for _, r := range u.Roles {
		if r == roleName {
			return nil
		}
	}
	u.Roles = append(u.Roles, roleName)
	return nil
}

// HasPermission checks if a user has a permission.
func HasPermission(username string, perm Permission) bool {
	mu.Lock()
	defer mu.Unlock()
	u, ok := users[username]
	if !ok {
		return false
	}
	for _, rname := range u.Roles {
		role, ok := roles[rname]
		if !ok {
			continue
		}
		for _, p := range role.Permissions {
			if p == perm {
				return true
			}
		}
	}
	return false
}
