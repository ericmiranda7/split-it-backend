package controllers

import (
	"net/http"
	"os"
	"split-that.com/split-that/v2/src/constants"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/service"
	"split-that.com/split-that/v2/src/util"
)

type RootController struct {
	userService *service.UserService
}

var rootController *RootController

func GetRootController(us *service.UserService) *RootController {
	if rootController == nil {
		initRootController(us)
	}
	return rootController
}

func initRootController(us *service.UserService) {
	rootController = &RootController{userService: us}
}

func (rc *RootController) InitHandlers(handlers []http.Handler) {
	http.Handle("/", rc.necessaryMiddleware(http.HandlerFunc(rc.rootHandler)))
	http.Handle("GET /api/health", rc.necessaryMiddleware(http.HandlerFunc(rc.healthHandler)))

	http.Handle("GET /api/accounts/", rc.authTokenExtractMiddleware(rc.necessaryMiddleware(handlers[0])))
	http.Handle("/api/auth", rc.necessaryMiddleware(handlers[1]))
}

// necessaryMiddleware Adds metadata to requests required by all requests
func (rc *RootController) necessaryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Received request at", r.URL)
		// TODO(eric): be strict
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv(constants.FrontendRedirectURL))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		next.ServeHTTP(w, r)
	})
}

// authTokenExtractMiddleware extracts the bearer token from the request header
func (rc *RootController) authTokenExtractMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug.Println("Extracting auth token from header")
		authHeader, ok := r.Header["Authorization"]
		if !ok {
			logger.Info.Println("Auth header not available")
			http.Error(w, "Invalid authorization", http.StatusUnauthorized)
			return
		}

		// verify bearer
		if authHeader[0][:6] != "Bearer" {
			logger.Debug.Println("Auth not of bearer type")
			http.Error(w, "Auth method not supported", http.StatusUnauthorized)
			return
		}

		authToken := authHeader[0][7:]
		claims, err := util.GetClaims(authToken)
		if err != nil {
			logger.Debug.Println("Unable to parse token")
			http.Error(w, "Unable to parse token", http.StatusUnauthorized)
			return
		}

		sub := claims["sub"].(string)

		if !rc.userService.ValidUser(sub) {
			logger.Debug.Println("No such subject (google_id) exists in db")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rc *RootController) healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Healthy af"))
	if err != nil {
		logger.Error.Println(err)
	}
}

func (rc *RootController) rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Hmm, not quite sure what you're looking for here."))
	if err != nil {
		logger.Error.Println(err)
	}
}
