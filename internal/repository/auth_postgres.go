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
	query := fmt.Sprintf("INSERT INTO %s(name, username, email, password_hash, token) values($1, $2, $3, $4, $5) RETURNING user_id", userTale)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Email, user.Password, user.Token)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (int, bool, error) {
	var result int
	var email bool
	query := fmt.Sprintf("SELECT user_id, is_email_verified FROM %s WHERE username = $1 AND password_hash = $2", userTale)
	row := r.db.QueryRow(query, username, password)
	if err := row.Scan(&result, &email); err != nil {
		return 0, false, err
	}
	return result, email, nil
}

func (r *AuthPostgres) VerifyEmail(token string) (int, error) {
	var result int
	query := fmt.Sprintf("UPDATE %s SET is_email_verified = true WHERE token = $1 RETURNING user_id", userTale)
	row := r.db.QueryRow(query, token)
	if err := row.Scan(&result); err != nil {
		return 0, err
	}
	return result, nil
}
