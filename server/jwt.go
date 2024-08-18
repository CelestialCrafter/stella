package server

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type userClaims struct {
	ID    string `json:"id"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func Sign(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func Verify(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unsupported signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(userClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	return echojwt.WithConfig(config)
}
