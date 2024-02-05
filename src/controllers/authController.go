package controllers

import (
	"context"
	"golang.org/x/oauth2"
	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"net/http"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/service"
)

type AuthController struct {
	as *service.AuthService
}

var authController *AuthController

func GetAuthController(as *service.AuthService) *AuthController {
	if authController == nil {
		authController = initializeAuthController(as)
	}

	return authController
}

func initializeAuthController(as *service.AuthService) *AuthController {
	return &AuthController{as: as}
}

func (ac *AuthController) GetOauthHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("code")

	// exchange code for access token
	token, err := ac.as.Cfg.Exchange(context.TODO(), authCode)
	if err != nil {
		logger.Error.Println("Error while exchanging for token with:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get userprof details from token
	client := oauth2.NewClient(context.TODO(), ac.as.Cfg.TokenSource(context.TODO(), token))
	srv, err := goauth2.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		logger.Error.Println("Error while creating service with client with:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo, err := srv.Userinfo.Get().Do()
	if err != nil {
		logger.Error.Println("Cannot get userInfo with ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	logger.Info.Println("USERINFO IS: ", userInfo.Name, userInfo.Id)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(userInfo.Email))
	if err != nil {
		return
	}
}
