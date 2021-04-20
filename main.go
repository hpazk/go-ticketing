package main

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/hpazk/go-booklib/apps"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Custom Validator
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Trailing slash
	e.Pre(middleware.RemoveTrailingSlash())

	// Static folder images
	e.Static("/", "public")

	// Main root
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.M{"message": "success"})
	})

	// App initialization
	apps.AppInit(e)

	// Run server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
