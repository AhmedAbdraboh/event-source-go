package models

import "roosh-app/helpers"

type AccountEvent struct {
	Id        int `json:"id"`
	AccountId int `json:"accountId"`
	Amount    int `json:"amount"`
}

var AccountEvents []*AccountEvent

func NewAccountEvent(accountId, amount int) *AccountEvent {
	id := helpers.GetNextAccountEventId()
	return &AccountEvent{Id: id, AccountId: accountId, Amount: amount}
}

func (accountEvent *AccountEvent) Create()  {
	AccountEvents = append(AccountEvents, accountEvent)
}

func (accountEvent *AccountEvent) GetAccountEvents(accountId int) []*AccountEvent {
	var resultAccountEvents []*AccountEvent
	for _, currentAccountEvent := range AccountEvents{
		if currentAccountEvent.AccountId == accountId {
			resultAccountEvents = append(resultAccountEvents, currentAccountEvent)
		}
	}
	return resultAccountEvents
}