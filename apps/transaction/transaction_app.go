package transaction

import (
	"github.com/hpazk/go-booklib/auth"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	Db *gorm.DB
}

func Init(db *gorm.DB) *App {
	return &App{db}
}

var handlers *handler

func (a *App) UseApp() {
	repo := transactionRepository(a.Db)
	service := transactionService(repo)
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
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/transactions/:id",
			Handler: handlers.GetTransaction,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},

		},
		{
			Method:  echo.PUT,
			Path:    "/transactions",
			Handler: handlers.PutTransaction,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},

		},
		{
			Method:  echo.DELETE,
			Path:    "/transactions",
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
