package user

import (
	"github.com/hpazk/go-booklib/auth"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserApp struct {
	Db *gorm.DB
}

func InitApp(db *gorm.DB) *UserApp {
	return &UserApp{db}
}

var handler *userHandler

func (r *UserApp) UseApp() {
	repository := userRepository(r.Db)
	userservice := UserService(repository)
	authService := auth.AuthService()

	handler = UserHandler(userservice, authService)
}

func (r *UserApp) Route() []helper.Route {

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
