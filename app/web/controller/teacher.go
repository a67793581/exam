package controller

import (
	"exam/app/service/token_jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	key := c.FormValue("key")

	if key != "carlo" {
		return echo.ErrUnauthorized
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, token_jwt.Claims{
		ID:       0,
		Identity: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	})

	// Generate encoded token and send it as response.
	//t, err := token.SigningString()
	t, err := token.SignedString([]byte(token_jwt.Key))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"token": t})
}
