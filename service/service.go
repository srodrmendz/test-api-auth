package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/srodrmendz/api-auth/model"
	"github.com/srodrmendz/api-auth/repository"
)

func New(repository repository.Repository, jwtSecretKey string) *AuthService {
	return &AuthService{
		repository:   repository,
		jwtSecretKey: jwtSecretKey,
	}
}

func (s *AuthService) Authenticate(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	user, err := s.repository.Authenticate(ctx, email, password)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	expires := now.Add(expiresAt)

	token, err := s.generateJwt(user.Email, user.UserName, now, expires)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token:     *token,
		ExpiresIn: expires.Unix(),
		ExpiresAt: expires,
	}, nil
}

func (s *AuthService) generateJwt(email string, username string, issuedAt, expires time.Time) (*string, error) {
	claims := jwtClaim{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expires.Unix(),
			Id:        uuid.NewString(),
			Issuer:    "test_app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tk, err := token.SignedString([]byte(s.jwtSecretKey))
	if err != nil {
		return nil, fmt.Errorf("signing jwt token %w", err)
	}

	return &tk, nil
}
