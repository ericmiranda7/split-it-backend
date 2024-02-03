package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"split-that.com/split-that/v2/src/logger"
)

var conn *pgx.Conn

func GetDb(connString string) *pgx.Conn {
	if conn != nil {
		return conn
	}

	return initDb(connString)
}

func initDb(connString string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		logger.Error.Println("Could not connect to database with: ", err)
	}

	return conn
}
