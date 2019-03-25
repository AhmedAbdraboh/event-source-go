package routes

import (
	"net/http"
	"roosh-app/controllers"
)

var accountController controllers.AccountController
var accountRoutes = []Route{
	{Path: "/accounts", Method: http.MethodGet, Handler: accountController.List},
	{Path: "/accounts", Method: http.MethodPost, Handler: accountController.Create},
	{Path: "/accounts/{id}", Method: http.MethodPut, Handler: accountController.UpdateOwner},
	{Path: "/accounts/{id}", Method: http.MethodGet, Handler: accountController.Aggregate},
	{Path: "/accounts/withdrawal", Method: http.MethodPost, Handler: accountController.Withdrawal},
	{Path: "/accounts/deposit", Method: http.MethodPost, Handler: accountController.Deposit},
}
