package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "qrkjk#4#35FSFJlja#4353KSFjH"
	salt       = "2jkhjojkhsdfkjghkjlfshngbkiwhjeoir"
)

type AuthService struct {
	repo repository.Authentification
}

func NewAuthService(repo repository.Authentification) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) SignUp(user models.User) (int, error) {
	user.Password = hashedPassword(user.Password)
	return a.repo.SignUp(user)
}

func hashedPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (a *AuthService) Token(user models.SignInInput) (string, error) {
	user1, err := a.repo.GetUser(user.Username, hashedPassword(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user1,
	})
	return token.SignedString([]byte(signingKey))
}
