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

type Repository struct {
	Authentification
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authentification: NewAuthPostgres(db)}
}
