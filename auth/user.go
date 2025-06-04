package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"sync"
)

type User struct {
	Username     string
	PasswordHash string
	Salt         string
	Email        string
}

type UserStore interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
}

type InMemoryUserStore struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{users: make(map[string]*User)}
}

func (s *InMemoryUserStore) Create(user *User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[user.Username]; exists {
		return errors.New("user exists")
	}
	s.users[user.Username] = user
	return nil
}

func (s *InMemoryUserStore) GetByUsername(username string) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func GenerateSalt() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func HashPassword(password, salt string) string {
	h := sha256.New()
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckPasswordHash(hash, password, salt string) bool {
	return hash == HashPassword(password, salt)
}
