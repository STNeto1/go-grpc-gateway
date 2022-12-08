package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUserIDFromToken(c echo.Context) string {
	payload := c.Get("claims").(*jwt.Token)
	claims := payload.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}
