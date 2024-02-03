package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"split-that.com/split-that/v2/src/logger"
)

var conn *pgx.Conn

func GetDb() *pgx.Conn {
	if conn != nil {
		return conn
	}

	return initDb()
}

func initDb() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://splitthat:DTCzWjVdF2mxBvrrdOPUm53MynB6oqWt@dpg-cmu41eq1hbls73d5t330-a.singapore-postgres.render.com/splitthat")

	if err != nil {
		logger.Error.Println("Could not connect to db with: ", err)
	}

	return conn
}
