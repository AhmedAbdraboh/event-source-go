package controllers

import (
	"net/http"
	"roosh-app/base"
)

type AccountController struct {
	base.Controller
}

func (accountController *AccountController) List(w http.ResponseWriter, r *http.Request)  {
	accountController.RespondWithJson(w, "Listing all accountskkkk ggggg", http.StatusOK)
	return
}

func (accountController *AccountController) Create(w http.ResponseWriter, r *http.Request)  {
	accountController.RespondWithJson(w, "Listing all accountskkkk ggggg", http.StatusOK)
	return
}

func (accountController *AccountController) Withdrawal(w http.ResponseWriter, r *http.Request)  {
	accountController.RespondWithJson(w, "Listing all accountskkkk ggggg", http.StatusOK)
	return
}

func (accountController *AccountController) Deposit(w http.ResponseWriter, r *http.Request)  {
	accountController.RespondWithJson(w, "Listing all accountskkkk ggggg", http.StatusOK)
	return
}