package server

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/CelestialCrafter/stella/common"
	"github.com/CelestialCrafter/stella/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const redirectUrl = "http://localhost:8000/api/auth/callback"

var config = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OAUTH_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_SECRET"),
	Scopes:       []string{"openid", "email"},
	RedirectURL:  redirectUrl,
	Endpoint:     google.Endpoint,
}

type userClaims struct {
	ID    string `json:"id"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type user struct {
	ID string `json:"id"`
}

func Login(c echo.Context) error {

	state := hex.EncodeToString(common.Hash())
	url := config.AuthCodeURL(state)

	c.SetCookie(&http.Cookie{
		Name:  "state",
		Value: state,

		MaxAge:   int((time.Minute * 5).Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return c.Redirect(http.StatusSeeOther, url)
}

func Callback(c echo.Context) error {
	originalState, err := c.Cookie("state")
	state := c.QueryParam("state")

	if err != nil || originalState.Value != state {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Could not verify state",
		})
	}
	oauthToken, err := config.Exchange(context.Background(), c.QueryParam("code"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	client := config.Client(context.Background(), oauthToken)
	res, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	claims := &userClaims{
		"",
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}

	err = json.Unmarshal(body, &claims)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	// @TODO refresh token
	claims.ID = fmt.Sprint("google-", claims.ID)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	err = db.CreateUser(claims.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
