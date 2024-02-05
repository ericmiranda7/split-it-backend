package controllers

import (
	"net/http"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/util"
)

func InitHandlers(handlers []http.Handler) {
	// todo(eric): diff between get, post, put?

	http.Handle("/api/health", necessaryMiddleware(http.HandlerFunc(healthHandler)))
	http.Handle("/", necessaryMiddleware(http.HandlerFunc(rootHandler)))

	http.Handle("/api/accounts/", necessaryMiddleware(handlers[0]))
	http.Handle("/api/auth", necessaryMiddleware(handlers[1]))
}

// todo(eric): token parsing middleware

func necessaryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Received request at", r.URL)
		// TODO(eric): be strict
		w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

func authTokenExtractMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug.Println("Extracting token auth from cookie")
		cookie, err := r.Cookie("auth")
		if err != nil {
			http.Error(w, "Invalid authorization", http.StatusUnauthorized)
			return
		}

		jwtTokenString := cookie.Value
		_, err = util.GetClaims(jwtTokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		logger.Debug.Println("Middleware Authorised success")

		next.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Healthy af"))
	if err != nil {
		logger.Error.Println(err)
	}
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Hmm, not quite sure what you're looking for here."))
	if err != nil {
		logger.Error.Println(err)
	}
}
