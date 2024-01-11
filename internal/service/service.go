package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
)

type ConfigEmail struct {
	EmailSenderName     string
	EmailSenderAddress  string
	EmailSenderPassword string
}
type Service struct {
	Authentification
}

type EmailSender interface {
	Sendmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error
}
type Authentification interface {
	SignUp(user models.User, cfg ConfigEmail) (int, error)
	Token(user models.SignInInput) (string, error)
	ParseToken(accesstoken string) (int, error)
	VerifyEmail(token string) (int, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authentification: NewAuthService(repo.Authentification)}
}
