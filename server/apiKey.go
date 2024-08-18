package server

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func NewApiKey(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*userClaims)
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365))

	newToken, err := Sign(claims)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": newToken,
	})
}
