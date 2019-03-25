package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"roosh-app/base"
	"roosh-app/models"
	"strconv"
)

type AccountController struct {
	base.Controller
}

func (accountController *AccountController) List(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	accounts := account.List()
	if accounts == nil {
		accountController.RespondWithJson(w, "No accounts created yet", http.StatusOK)
		return
	}
	accountController.RespondWithJson(w, accounts, http.StatusOK)
	return
}

func (accountController *AccountController) Aggregate(w http.ResponseWriter, r *http.Request) {
	var account *models.Account
	accountId := chi.URLParam(r, "id")
	accountIdInt, err := strconv.Atoi(accountId)
	if err != nil {
		accountController.RespondWithJson(w, "Error while converting id, use int id", http.StatusBadRequest)
		return
	}
	account = account.Get(accountIdInt)
	if account == nil {
		accountController.RespondWithJson(w, "account not found ", http.StatusBadRequest)
		return
	}
	aggregatedAccount := account.Aggregate()
	accountController.RespondWithJson(w, aggregatedAccount, http.StatusOK)
	return
}
func (accountController *AccountController) Create(w http.ResponseWriter, r *http.Request) {
	account := models.NewAccount()
	account.Create()
	accountController.RespondWithJson(w, "Account Created with id "+strconv.Itoa(account.Id), http.StatusOK)
	return
}

func (accountController *AccountController) UpdateOwner(w http.ResponseWriter, r *http.Request) {

	var account *models.Account
	accountId := chi.URLParam(r, "id")
	accountIdInt, err := strconv.Atoi(accountId)
	if err != nil {
		accountController.RespondWithJson(w, "Error while converting id, use int id", http.StatusBadRequest)
		return
	}
	account = account.Get(accountIdInt)
	if account == nil {
		accountController.RespondWithJson(w, "account not found ", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		accountController.RespondWithJson(w, "Error while decoding body", http.StatusBadRequest)
		return
	}
	account.Update()
	accountController.RespondWithJson(w, "Owner updated to "+account.Owner, http.StatusOK)
	return
}

func (accountController *AccountController) Withdrawal(w http.ResponseWriter, r *http.Request) {
	var accountEvent models.AccountEvent
	if err := json.NewDecoder(r.Body).Decode(&accountEvent); err != nil {
		accountController.RespondWithJson(w, "Error while decoding body", http.StatusBadRequest)
		return
	}
	var account *models.Account
	account = account.Get(accountEvent.AccountId)
	if account == nil {
		accountController.RespondWithJson(w, "account not found ", http.StatusBadRequest)
		return
	}
	aggregatedAccount := account.Aggregate()
	if aggregatedAccount.Balance < accountEvent.Amount {
		accountController.RespondWithJson(w, "Not enough balance", http.StatusBadRequest)
		return
	}
	accountEvent.Amount = -accountEvent.Amount
	accountEvent.Create()
	accountController.RespondWithJson(w, "Withdrawal from account "+strconv.Itoa(account.Id)+" amount "+strconv.Itoa(-accountEvent.Amount), http.StatusOK)
	return
}

func (accountController *AccountController) Deposit(w http.ResponseWriter, r *http.Request) {
	var accountEvent models.AccountEvent
	if err := json.NewDecoder(r.Body).Decode(&accountEvent); err != nil {
		accountController.RespondWithJson(w, "Error while decoding body", http.StatusBadRequest)
		return
	}
	var account *models.Account
	account = account.Get(accountEvent.AccountId)
	if account == nil {
		accountController.RespondWithJson(w, "account not found ", http.StatusBadRequest)
		return
	}
	accountEvent.Create()
	accountController.RespondWithJson(w, "Account with id "+strconv.Itoa(accountEvent.AccountId)+"deposited with amount"+strconv.Itoa(accountEvent.Amount), http.StatusOK)
	return
}
