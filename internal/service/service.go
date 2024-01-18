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
	Hub
}

type EmailSender interface {
	Sendmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error
}

type Hub interface {
	AllServers() ([]models.Room, error)
	GetServerById(id_room int) (models.Room, error)
	CreateServer(username string) (int, error)
	Connect(room_id int, username string) (models.Room, error)
	Disconnect(room_id int, username string) (models.Room, error)
	IsServerFull(id_room int) (bool, error)
	GetUsernameById(id int) (string, error)
	GetUsersByRoomId(id int) []int
}
type Authentification interface {
	SignUp(user models.User, cfg ConfigEmail) (int, error)
	Token(user models.SignInInput) (string, error)
	ParseToken(accesstoken string) (int, error)
	VerifyEmail(token string) (int, error)
	RefreshToken(token string) (string, error)
	GetRefresh(input models.SignInInput) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authentification: NewAuthService(repo.Authentification), Hub: (repo.Hub)}
}
