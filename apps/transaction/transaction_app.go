package transaction

import (
	"github.com/hpazk/go-booklib/auth"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

type App struct {
}

var handlers *handler

func (a *App) UseApp() {
	service := TransactionService()
	authService := auth.AuthService()

	handlers = transactionHandler(service, authService)
}

func (a *App) Route() []helper.Route {

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/transactions",
			Handler: handlers.PostTransaction,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/transactions",
			Handler: handlers.GetTransactions,
			// Middleware: []echo.MiddlewareFunc{handlers.GetTransactionsCached},
		},
		{
			Method:  echo.GET,
			Path:    "/transactions/:id",
			Handler: handlers.GetTransaction,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},

		},
		{
			Method:  echo.PUT,
			Path:    "/transactions/:id",
			Handler: handlers.PutTransaction,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},

		},
		{
			Method:  echo.DELETE,
			Path:    "/transactions/:id",
			Handler: handlers.DeleteTransaction,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/transactions/event/:id",
			Handler: handlers.GetTransactionsByEvent,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
	}
}
