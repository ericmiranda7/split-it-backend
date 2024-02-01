package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"split-that.com/split-that/v2/src/constants"
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

	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// consider cors from localhost:5173
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// write hi to body
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, "20")
		if err != nil {
			return
		}
	})

	fmt.Println("Running server on port:", ServerPort)
	err = http.ListenAndServe(fmt.Sprint(":", ServerPort), rootHandler)
	if err != nil {
		log.Fatal(err)
	}
}
