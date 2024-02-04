package service

import (
	"github.com/jackc/pgx/v5"
	"split-that.com/split-that/v2/src/models"
)

var accountService *AccountService

type AccountService struct {
	conn *pgx.Conn
}

func GetAccountService(conn *pgx.Conn) *AccountService {
	if accountService == nil {
		accountService = initAccountService(conn)
	}

	return accountService
}

func initAccountService(conn *pgx.Conn) *AccountService {
	return &AccountService{conn: conn}
}

func (as *AccountService) GetTotalOwe(user models.User) float64 {
	println("lols")
	return -2.3
}
