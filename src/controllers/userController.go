package controllers

import (
	"net/http"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/service"
)

var userController *UserController

type UserController struct {
	userService *service.UserService
}

func GetUserController(us *service.UserService) *UserController {
	if userController == nil {
		initUserController(us)
	}

	return userController
}

func initUserController(us *service.UserService) {
	userController = &UserController{userService: us}
}

func (uc *UserController) UserHandler(w http.ResponseWriter, r *http.Request) {
	// todo(eric): don't hardcode name
	created, err := uc.userService.CreateUser("Eric")
	if created {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err)
	}
}
