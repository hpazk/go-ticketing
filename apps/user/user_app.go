package user

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

var handler *userHandler

func (a *App) UseApp() {
	repository := userRepository(a.Db)
	userservice := UserService(repository)
	authService := auth.AuthService()

	handler = UserHandler(userservice, authService)
}

func (a *App) Route() []helper.Route {

	// TODO jwt: OK
	// TODO jwt: custom-error
	// TODO token-validation

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/registration",
			Handler: handler.PostUserRegistration,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: handler.PostUserLogin,
		},
		{
			Method:     echo.POST,
			Path:       "/logout",
			Handler:    handler.PostUserLogout,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		// {
		// 	Method:  echo.GET,
		// 	Path:    "/users",
		// 	Handler: handler.GetUsers,
		// 	Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		// },
		// {
		// 	Method:     echo.GET,
		// 	Path:       "/users/:id",
		// 	Handler:    handler.GetUser,
		// 	Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		// },
		// {
		// 	Method:     echo.PUT,
		// 	Path:       "/users/:id",
		// 	Handler:    handler.PutUser,
		// 	Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		// },
		// {
		// 	Method:     echo.DELETE,
		// 	Path:       "/users/:id",
		// 	Handler:    handler.DeleteUser,
		// 	Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		// },
	}
}
