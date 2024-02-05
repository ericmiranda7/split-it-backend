package controllers

import (
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
