package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var jwtKey = []byte("change_this_secret")

type Claims struct {
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

func GenerateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		Exp:      time.Now().Add(24 * time.Hour).Unix(),
	}
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)
	unsigned := header + "." + payload
	sig := hmacSHA256(unsigned)
	token := unsigned + "." + base64.RawURLEncoding.EncodeToString(sig)
	return token, nil
}

func ParseJWT(tokenStr string) (*Claims, error) {
	parts := strings.Split(tokenStr, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token")
	}
	unsigned := parts[0] + "." + parts[1]
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}
	if !hmac.Equal(sig, hmacSHA256(unsigned)) {
		return nil, errors.New("invalid signature")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}
	if time.Now().Unix() > claims.Exp {
		return nil, errors.New("token expired")
	}
	return &claims, nil
}

func hmacSHA256(data string) []byte {
	h := hmac.New(sha256.New, jwtKey)
	h.Write([]byte(data))
	return h.Sum(nil)
}
