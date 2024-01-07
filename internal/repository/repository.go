package repository

import (
	"awesomeProject/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authentification interface {
	SignUp(user models.User) (int, error)
}

type Repository struct {
	Authentification
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authentification: NewAuthPostgres(db)}
}
