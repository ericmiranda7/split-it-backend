package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

var userService *UserService

type UserService struct {
	conn *pgx.Conn
}

func GetUserService(conn *pgx.Conn) *UserService {
	if userService == nil {
		initUserService(conn)
	}
	return userService
}

func initUserService(conn *pgx.Conn) {
	userService = &UserService{
		conn: conn,
	}
}

// CreateUser creates a user in the database if not exists
func (us *UserService) CreateUser(sub string) error {
	_, err := us.conn.Exec(context.Background(), `INSERT INTO users ("google_id") VALUES ($1) ON CONFLICT DO NOTHING`, sub)
	if err != nil {
		return errors.New("unable to create user")
	}
	return nil
}

func (us *UserService) ValidUser(sub string) bool {
	row := us.conn.QueryRow(context.Background(), `SELECT google_id FROM users WHERE google_id = $1`, sub)

	var dbSub string
	err := row.Scan(&dbSub)
	if err != nil {
		return false
	}

	return true
}
