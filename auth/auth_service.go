package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type AuthServices interface {
	GetAccessToken(userId uint, userRole string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authService struct {
}

func AuthService() *authService {
	return &authService{}
}

func (s *authService) GetAccessToken(userId uint, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["user_role"] = userRole

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
