package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func runServer() error {
	logger := log.Default()
	userStore := make([]string, 0)
	appServer := NewServer(logger, userStore)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: appServer,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func NewServer(logger *log.Logger, userStore []string) http.Handler {
	mux := http.NewServeMux()

	balHandler := http.HandlerFunc(handleBalance)
	// handlers
	mux.Handle("/balance", balHandler)
	mux.Handle("POST /users", handleCreateUser(&userStore, logger))
	mux.Handle("GET /users", handleReadUser(&userStore))

	return loggingMiddleware(logger, mux)
}

func handleBalance(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("2.3"))
	if err != nil {
		return
	}
}

func handleCreateUser(userStore *[]string, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bb, _ := io.ReadAll(r.Body)
		logger.Println("got", string(bb[:]))
		*userStore = append(*userStore, string(bb[:]))
	})
}

func handleReadUser(userStore *[]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(strings.Join(*userStore, ", ")))
		if err != nil {
			return
		}
	})
}

func loggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Logging middleware for %v", r.URL)
		next.ServeHTTP(w, r)
	})
}
