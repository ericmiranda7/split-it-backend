package service

import (
	"context"
	"github.com/jackc/pgx/v5"
)

var userService *UserService

type UserService struct {
	conn *pgx.Conn
}

func initUserService(conn *pgx.Conn) {
	userService = &UserService{
		conn: conn,
	}
}

func GetUserService(conn *pgx.Conn) *UserService {
	initUserService(conn)
	return userService
}

func (us *UserService) CreateUser(name string) (bool, error) {
	_, err := us.conn.Exec(context.Background(), "INSERT INTO users (\"name\") VALUES ($1)", name)
	if err != nil {
		return false, err
	}
	return true, nil
}
