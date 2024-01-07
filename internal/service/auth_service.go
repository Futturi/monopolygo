package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
	"crypto/sha1"
	"fmt"
)

const (
	salt = "2jkhjojkhsdfkjghkjlfshngbkiwhjeoir"
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
