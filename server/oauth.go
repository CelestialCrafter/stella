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

type User struct {
	ID string
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Could not verify state",
		})
	}
	token, err := config.Exchange(context.Background(), c.QueryParam("code"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	client := config.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	user.ID = fmt.Sprint("google-", user.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}
