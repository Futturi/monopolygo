package repository

import (
	"awesomeProject/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authentification interface {
	SignUp(user models.User) (int, error)
	GetUser(username, password string, refresh models.RefreshToken) (int, bool, error)
	VerifyEmail(token string) (int, error)
	GetUserByToken(token string) (int, error)
	GetRefreshToken(input models.SignInInput) (string, error)
}

type Hub interface {
	AllServers() ([]models.Room, error)
	GetServerById(id_room int) (models.Room, error)
	CreateServer(username string) (int, error)
	GetUserIdByUsername(username string) (int, error)
}
type Repository struct {
	Authentification
	Hub
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authentification: NewAuthPostgres(db), Hub: NewHubPostgres(db)}
}
