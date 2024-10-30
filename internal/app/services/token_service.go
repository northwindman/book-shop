package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/northwindman/book-shop/internal/app/domain"
)

// TODO: move to secrets
var jwtSecretKey = []byte("very-secret-key")

// TokenService is a token service
type TokenService struct {
	ttl time.Duration
}

// NewTokenService creates a new token service
func NewTokenService(ttl time.Duration) TokenService {
	return TokenService{
		ttl: ttl,
	}
}

type UserClaims struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

// GenerateToken generates a token
func (s TokenService) GenerateToken(user domain.User) (string, error) {
	payload := UserClaims{
		UserID:   user.ID(),
		UserName: user.Username(),
		Admin:    user.Admin(),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return t, nil
}

func (s TokenService) GetUser(token string) (domain.User, error) {
	var userClaims UserClaims
	t, err := jwt.ParseWithClaims(token, &userClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to parse a token: %w", err)
	}
	if !t.Valid {
		return domain.User{}, errors.New("invalid token")
	}
	user, err := userClaimsToDomainUser(userClaims)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to convert user claims to domain user: %w", err)
	}
	return user, nil
}

func userClaimsToDomainUser(claims UserClaims) (domain.User, error) {
	return domain.NewUser(domain.NewUserData{
		ID:       claims.UserID,
		Username: claims.UserName,
		Admin:    claims.Admin,
	})
}
