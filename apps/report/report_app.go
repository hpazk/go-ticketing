package report

import (
	"github.com/hpazk/go-ticketing/auth"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
)

type App struct {
}

var handlers *handler

func (a *App) UseApp() {

	service := ReportService()
	authService := auth.AuthService()

	handlers = reportHandler(service, authService)
}

func (a *App) Route() []helper.Route {

	return []helper.Route{
		{
			Method:     echo.GET,
			Path:       "/report/participant",
			Handler:    handlers.GetEventParticipant,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		// {
		// 	Method:     echo.GET,
		// 	Path:       "/events/:id/report",
		// 	Handler:    handlers.GetEventReport,
		// 	Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		// },
	}
}
