package server

import (
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

func sign(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
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
