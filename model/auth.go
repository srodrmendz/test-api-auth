package model

import "time"

type AuthResponse struct {
	Token     string    `json:"token"`
	ExpiresIn int64     `json:"expires_in"`
	ExpiresAt time.Time `json:"expires_at"`
}
