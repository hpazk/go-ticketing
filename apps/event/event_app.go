package event

import (
	"github.com/hpazk/go-ticketing/auth"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
)

type App struct {
}

var handlers *handler

func (a *App) UseApp() {

	service := EventService()
	authService := auth.AuthService()

	handlers = eventHandler(service, authService)
}

func (a *App) Route() []helper.Route {

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/events",
			Handler: handlers.PostEvent,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/events",
			Handler: handlers.GetEvents,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/events/:id",
			Handler: handlers.GetEvent,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},

		},
		{
			Method:  echo.PUT,
			Path:    "/events",
			Handler: handlers.PutEvent,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},

		},
		{
			Method:  echo.DELETE,
			Path:    "/events",
			Handler: handlers.DeleteEvent,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/events/:id/participant",
			Handler: handlers.DeleteEvent,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/events/:id/report",
			Handler: handlers.GetEventReport,
			// Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
	}
}
