package routes

import (
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

func DefineApiRoutes(e *echo.Echo, handlers []helper.Handler) {
	var routes []helper.Route

	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}

	api := e.Group("/api/v1")

	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
			}
		}
	}
}
