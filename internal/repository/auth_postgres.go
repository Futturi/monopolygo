package repository

import (
	"awesomeProject/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignUp(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name, username, email, password_hash) values($1, $2, $3, $4) RETURNING user_id", userTale)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
