package models

import (
	"fmt"
	"math/rand"
	"time"
)

type RefreshToken struct {
	Token     string `json:"refresh_token" db:"refresh_token"`
	ExpiresAt int64  `json:"expires_at" db:"expires_at"`
}

func NewRefreshToken() RefreshToken {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	_, err := r.Read(b)
	if err != nil {
		return RefreshToken{}
	}
	return RefreshToken{
		Token:     fmt.Sprintf("%x", b),
		ExpiresAt: time.Now().Add(30 * time.Hour).Unix(),
	}
}
