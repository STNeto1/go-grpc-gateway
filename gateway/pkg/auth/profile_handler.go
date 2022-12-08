package auth

import (
	"__lib/auth"
	"__lib/exceptions"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) profile(c echo.Context) error {
	id := auth.GetUserIDFromToken(c)

	usr, err := auth.GetUser(id, h.userClient, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, exceptions.Unauthorized{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		})
	}

	return c.JSON(200, usr)
}
