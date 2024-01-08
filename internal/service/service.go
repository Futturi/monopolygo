package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
)

type Service struct {
	Authentification
}

type Authentification interface {
	SignUp(user models.User) (int, error)
	Token(user models.SignInInput) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authentification: NewAuthService(repo.Authentification)}
}
