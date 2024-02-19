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
	// todo(eric): get that us,as,ac names shiz refactored properly
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file")
	}
	serverPort := os.Getenv(constants.ServerPort)
	if serverPort == "" {
		serverPort = "8080"
	}

	connString := os.Getenv(constants.ConnString)
	logger.Debug.Println("connstring", connString, serverPort)

	conn := database.GetDb(connString)
	us := service.GetUserService(conn)
	as := service.GetAccountService(conn)
	ac := controllers.GetAccountController(as)
	authC := controllers.GetAuthController(us)
	rootC := controllers.GetRootController(us)
	var handlers []http.Handler
	handlers = append(handlers, http.HandlerFunc(ac.GetAccountHandler))
	handlers = append(handlers, http.HandlerFunc(authC.GetOauthHandler))
	rootC.InitHandlers(handlers)

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
