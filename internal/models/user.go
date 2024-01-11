package models

type User struct {
	Userid            int    `json:"user_id" db:"user_id"`
	Name              string `json:"name" db:"name"`
	Username          string `json:"username" db:"username"`
	Email             string `json:"email" db:"email"`
	Password          string `json:"password" db:"password_hash"`
	Token             string `db:"token"`
	Is_email_verified bool   `db:"is_email_verified"`
}

type SignInInput struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}
