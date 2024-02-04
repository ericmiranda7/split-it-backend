package controllers

import (
	"net/http"
	"split-that.com/split-that/v2/src/logger"
	"split-that.com/split-that/v2/src/models"
	"split-that.com/split-that/v2/src/service"
	"strconv"
)

var accountController *AccountController

type AccountController struct {
	accountService *service.AccountService
}

func GetAccountController(as *service.AccountService) *AccountController {
	if accountController == nil {
		accountController = initAccountController(as)
	}

	return accountController
}

func initAccountController(as *service.AccountService) *AccountController {
	return &AccountController{accountService: as}
}

func (ac *AccountController) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	// todo(): confirm get request

	// todo(): extract user from request
	user := models.User{Name: "Eric"}

	println("am fine")
	owe := ac.accountService.GetTotalOwe(user)

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(strconv.FormatFloat(owe, 'f', 2, 64)))

	if err != nil {
		logger.Error.Println(err)
	}
}
