package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
    GenerateToken(email string) (string, error)
}

type authService struct {
    jwtSecret string
}

func NewAuthService(secret string) AuthService {
    return &authService{jwtSecret: secret}
}

func (s *authService) GenerateToken(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString([]byte(s.jwtSecret))
}
