package user

import (
	"github.com/hpazk/go-ticketing/auth"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
)

type App struct {
}

var handler *userHandler

func (a *App) UseApp() {
	userservice := UserService()
	authService := auth.AuthService()

	handler = UserHandler(userservice, authService)
}

func (a *App) Route() []helper.Route {
	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: handler.PostUserRegistration,
		},
		{
			Method:     echo.POST,
			Path:       "/users/creator",
			Handler:    handler.PostNewCreator,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: handler.PostUserLogin,
		},
		{
			Method:     echo.GET,
			Path:       "/users",
			Handler:    handler.GetUsers,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:     echo.PATCH,
			Path:       "/users/:id",
			Handler:    handler.PatchUser,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:     echo.DELETE,
			Path:       "/users/:id",
			Handler:    handler.DeleteUser,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		// {
		// 	Method:     echo.POST,
		// 	Path:       "/logout",
		// 	Handler:    handler.PostUserLogout,
		// 	Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		// },
		// {
		// 	Method:     echo.GET,
		// 	Path:       "/users/:id",
		// 	Handler:    handler.GetUser,
		// 	Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		// },
	}
}
