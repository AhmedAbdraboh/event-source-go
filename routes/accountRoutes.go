package routes

import (
	"net/http"
	"roosh-app/controllers"
)

var accountController controllers.AccountController
var accountRoutes = []Route{
	{Path: "/accounts", Method: http.MethodGet, Handler: accountController.List},
	{Path: "/accounts", Method:http.MethodPost, Handler: accountController.Create},
	{Path: "/accounts/:id/withdrawal", Method: http.MethodPost, Handler: accountController.Withdrawal},
	{Path: "/accounts/:id/deposit", Method: http.MethodPost, Handler: accountController.Deposit},

}
