package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type JwtCustomClaims struct {
// 	Id   uint   `json:"id"`
// 	Role string `json:"role"`
// 	jwt.StandardClaims
// }

func JwtMiddleWare() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      jwt.MapClaims{},
		SigningKey:  []byte(key),
		ContextKey:  "user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	})
}
