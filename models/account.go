package models

import "roosh-app/helpers"

type Account struct {
	Id      int    `json:"id"`
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

var Accounts []*Account

func NewAccount() *Account {
	//should be generated from data store
	newId := helpers.GetNextAccountId()
	return &Account{Id: newId}

}
func (account *Account) Create() {
	Accounts = append(Accounts, account)
}

func (account *Account) Update() {
	for _, currentAccount := range Accounts {
		if currentAccount.Id == account.Id {
			currentAccount = account
			break
		}
	}
}

func (account *Account) Get(id int) *Account {
	for _, currentAccount := range Accounts {
		if currentAccount.Id == id {
			return currentAccount
		}
	}
	return nil
}

func (account *Account) Aggregate() *Account {
	var accountEvent *AccountEvent
	AccountEvents := accountEvent.GetAccountEvents(account.Id)
	if AccountEvents == nil {
		return nil
	}
	newAccount := *account
	for _, currentAccountEvent := range AccountEvents {
		newAccount.Balance += currentAccountEvent.Amount
	}
	return &newAccount
}

func (account *Account) List() []*Account {
	return Accounts
}