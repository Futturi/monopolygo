package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"time"
)

const (
	signingKey = "qrkjk#4#35FSFJlja#4353KSFjH"
	salt       = "2jkhjojkhsdfkjghkjlfshngbkiwhjeoir"
)

type AuthService struct {
	repo repository.Authentification
}

func NewAuthService(repo repository.Authentification) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) SignUp(user models.User, cfg ConfigEmail) (int, error) {
	user.Token = CreateTokenForAccess(user)
	SendMail(cfg, user.Email, user.Token)
	user.Password = hashedPassword(user.Password)
	return a.repo.SignUp(user)
}

func hashedPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (a *AuthService) Token(user models.SignInInput) (string, error) {
	user1, verifyEmail, err := a.repo.GetUser(user.Username, hashedPassword(user.Password), models.NewRefreshToken())
	if err != nil {
		return "", err
	}
	if !verifyEmail {
		return "", errors.New("your email is not verified")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(50 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user1,
	})
	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) RefreshToken(token string) (string, error) {
	user, err := a.repo.GetUserByToken(token)
	if err != nil {
		return "", err
	}
	accesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user})
	return accesstoken.SignedString([]byte(signingKey))
}

func (r *AuthService) GetRefresh(input models.SignInInput) (string, error) {

	refresh_token, err := r.repo.GetRefreshToken(input)
	if err != nil {
		return "", err
	}
	return refresh_token, nil
}
func (r *AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (r *AuthService) VerifyEmail(token string) (int, error) {
	return r.repo.VerifyEmail(token)
}
func SendMail(cfg ConfigEmail, tomail string, token string) {
	sender := NewGmailSender(cfg.EmailSenderName, cfg.EmailSenderAddress, cfg.EmailSenderPassword)

	text := "Authorization in monopoly"
	content := fmt.Sprintf("Authorization link: localhost:8000/auth/%s", token)

	to := []string{tomail}
	err := sender.Sendmail(text, content, to, nil, nil, nil)

	require.NoError(nil, err)
}

func CreateTokenForAccess(user models.User) string {
	token := make([]byte, 32)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}
