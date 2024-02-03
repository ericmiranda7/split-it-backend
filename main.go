package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"split-that.com/split-that/v2/src/constants"
	"split-that.com/split-that/v2/src/controller"
	"split-that.com/split-that/v2/src/logger"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file")
	}

	ServerPort := os.Getenv(constants.ServerPort)
	if ServerPort == "" {
		ServerPort = "8080"
	}

	fmt.Println("Running server on port:", ServerPort)

	server := http.Server{
		Addr:     ":" + ServerPort,
		Handler:  nil,
		ErrorLog: logger.Error,
	}
	controller.InitHandlers()

	err = server.ListenAndServe()
	if err != nil {
		return
	}

}
