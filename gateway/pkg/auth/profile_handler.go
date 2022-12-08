package auth

import (
	"__lib/exceptions"
	userpb "__user/gen/pb/user/v1"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h handler) profile(c echo.Context) error {
	payload := c.Get("claims").(*jwt.Token)
	claims := payload.Claims.(jwt.MapClaims)

	usr, err := h.userClient.GetUser(c.Request().Context(), &userpb.GetUserRequest{
		Id: claims["sub"].(string),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, exceptions.Unauthorized{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		})
	}

	return c.JSON(200, usr.User)
}
