package service

import (
	"fmt"
	adapter "github.com/ikhsanrifff/go-banking-auth/adapter/repository"
	config "github.com/ikhsanrifff/go-banking-auth/config"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LoginAccount(username, password string) (string, error)
	CreateToken(user_id, token, expired_date string) (string, error)
}

type AuthAdapterDB struct {
	repo adapter.AccountRepository
}

func NewAuthService(repo adapter.AccountRepository) *AuthAdapterDB {
	return &AuthAdapterDB{repo: repo}
}

func (u *AuthAdapterDB) LoginAccount(username, password string) (string, error) {
	user, err := u.repo.GetAccountByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password: %v", err)
	}

	token, err := config.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return "", fmt.Errorf("could not generate token: %v", err)
	}

	return token, nil
}

func (a *AuthAdapterDB) CreateToken(user_id, token, expired_date string) (string, error) {
	user, err := a.repo
	if err != nil {
		return user, err
	}
}