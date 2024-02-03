package controller

import (
	"net/http"
	"split-that.com/split-that/v2/src/logger"
)

func InitHandlers() {
	http.Handle("/api/health", NecessaryMiddleware(http.HandlerFunc(healthHandler)))
	http.Handle("/", NecessaryMiddleware(http.HandlerFunc(rootHandler)))
}

func NecessaryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Received request at", r.URL)
		// TODO(eric): be strict
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Healthy af"))
	if err != nil {
		logger.Error.Println(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Hmm, not quite sure what you're looking for here."))
	if err != nil {
		logger.Error.Println(err)
	}
}
