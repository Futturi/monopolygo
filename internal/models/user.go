package models

type User struct {
	Userid   int    `json:"user_id" db:"user_id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}

type SignInInput struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}
