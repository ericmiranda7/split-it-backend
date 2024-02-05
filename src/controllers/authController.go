package controllers

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
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
	jwtToken := r.FormValue("credential")
	token, err := jwt.Parse(jwtToken, nil)
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		logger.Info.Println("Name is", claims["name"])
	} else {
		log.Printf("Invalid JWT Token")
		http.Error(w, "Invalid JWT Token", http.StatusUnauthorized)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(`{"status": "success", "user": {...}}`))
	if err != nil {
		return
	}
}
