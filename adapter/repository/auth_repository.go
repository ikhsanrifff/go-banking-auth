package adapter

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ikhsanrifff/go-banking-auth/domain"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	CreateToken(user_id, token, expired_date string) error
	GetAccountByUsername(username string) (*domain.Account, error)
}

type AuthRepositoryDB struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepositoryDB {
	return &AuthRepositoryDB{DB: db}
}

func (a *AuthRepositoryDB) CreateToken(user_id, token, expired_date string) error {
	user_id = uuid.New().String()
	query := "INSERT INTO store (user_id, token, expired_at) VALUES (?, ?, ?);"
	_, err := a.DB.Exec(query, user_id, token, expired_date)
	if err != nil {
		return fmt.Errorf("failed to create token: %v", err)
	}
	return err
}

func (a *AuthRepositoryDB) GetAccountByUsername(username string) (*domain.Account, error) {
	var account domain.Account
	query := "SELECT id, customer_id, username, password, balance, currency, status FROM accounts WHERE username = ?"
	err := a.DB.Get(&account, query, username)
	if err != nil {
		if account == (domain.Account{}) {
			return nil, fmt.Errorf("no accounts found")
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	return &account, nil
}