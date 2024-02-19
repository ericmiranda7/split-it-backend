package controllers

import (
	"fmt"
	"net/http"
	"os"
	"split-that.com/split-that/v2/src/constants"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/service"
	"split-that.com/split-that/v2/src/util"
)

type AuthController struct {
	us *service.UserService
}

var authController *AuthController

func GetAuthController(us *service.UserService) *AuthController {
	if authController == nil {
		authController = initializeAuthController(us)
	}

	return authController
}

func initializeAuthController(us *service.UserService) *AuthController {
	return &AuthController{us: us}
}

func (ac *AuthController) GetOauthHandler(w http.ResponseWriter, r *http.Request) {
	// todo(): probably validate the token from google
	jwtToken := r.FormValue("credential")
	claims, err := util.GetClaims(jwtToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
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
	redirectUrl := os.Getenv(constants.FrontendRedirectURL) +
		fmt.Sprintf("/login?name=%s&redirect=true&token=%s", name, jwtToken)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)

	if err != nil {
		return
	}
}
