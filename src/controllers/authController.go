package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"split-that.com/split-that/v2/src/constants"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/service"
)

type AuthController struct {
	as *service.AuthService
	us *service.UserService
}

var authController *AuthController

func GetAuthController(as *service.AuthService, us *service.UserService) *AuthController {
	if authController == nil {
		authController = initializeAuthController(as, us)
	}

	return authController
}

func initializeAuthController(as *service.AuthService, us *service.UserService) *AuthController {
	return &AuthController{as: as, us: us}
}

func (ac *AuthController) GetOauthHandler(w http.ResponseWriter, r *http.Request) {
	// todo(): probably validate the token from google
	jwtToken := r.FormValue("credential")
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})

	if err != nil {
		logger.Error.Println("Error parsing JWT Token", err.Error())
		return
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		logger.Error.Println("Invalid JWT token")
		http.Error(w, "Invalid JWT Token", http.StatusUnauthorized)
		return
	}

	name := claims["name"].(string)
	sub := claims["sub"].(string)

	// create the user in db
	err = ac.us.CreateUser(name, sub)
	if err != nil {
		logger.Error.Println("Couldn't create user due to", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// redirect back to frontend
	// todo(): prolly generate and sign your own token for the frontend?
	w.Header().Set("Location", os.Getenv(constants.FrontendRedirectURL)+fmt.Sprintf("login?token=%s&name=%s", jwtToken, name))

	w.WriteHeader(http.StatusSeeOther)

	if err != nil {
		return
	}
}
