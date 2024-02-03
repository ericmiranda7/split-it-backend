package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"split-that.com/split-that/v2/src/constants"
	"split-that.com/split-that/v2/src/controllers"
	"split-that.com/split-that/v2/src/database"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/service"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file")
	}
	env := os.Getenv(constants.Env)
	serverPort := os.Getenv(constants.ServerPort)
	if serverPort == "" {
		serverPort = "8080"
	}

	var connString string
	if env == "dev" {
		connString = os.Getenv(constants.LocalConnString)
	} else {
		connString = os.Getenv(constants.ConnString)
	}
	logger.Debug.Println("connstring", connString, serverPort)

	conn := database.GetDb(connString)
	us := service.GetUserService(conn)
	uc := controllers.GetUserController(us)
	var handlers []http.Handler
	handlers = append(handlers, http.HandlerFunc(uc.UserHandler))
	controllers.InitHandlers(handlers)

	fmt.Println("Running server on port:", serverPort)
	server := http.Server{
		Addr:     ":" + serverPort,
		Handler:  nil,
		ErrorLog: logger.Error,
	}

	err = server.ListenAndServe()
	if err != nil {
		return
	}
}
