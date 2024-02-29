package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/srodrmendz/api-auth/model"
	"github.com/srodrmendz/api-auth/repository"
)

// Check on build time that AuthService implement Service interface
var _ Service = (*AuthService)(nil)

const expiresAt = 1 * time.Hour

// Service defines the methods that should be implemented by a auth service.
type Service interface {
	// Authenticate a user
	// Returns error if user is not found exists or there is an error in the system
	Authenticate(ctx context.Context, email string, password string) (*model.AuthResponse, error)
}

type AuthService struct {
	repository   repository.Repository
	jwtSecretKey string
}

type jwtClaim struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}
