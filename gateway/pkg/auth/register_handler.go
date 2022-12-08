package auth

import (
	"__lib/exceptions"
	v "__lib/validator"
	userpb "__user/gen/pb/v1"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RegisterRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h handler) register(c echo.Context) error {
	body := new(RegisterRequestBody)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, exceptions.BadRequest{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
		})
	}
	if err := c.Validate(body); err != nil {
		errors := v.ParseValidatorErrors(err)
		return echo.NewHTTPError(http.StatusBadRequest, exceptions.BadValidation{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
			Errors:     errors,
		})
	}

	res, err := h.userClient.Register(c.Request().Context(), &userpb.RegisterRequest{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, exceptions.Unauthorized{
			Message:    "Invalid credentials",
			StatusCode: http.StatusUnauthorized,
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Subject:   res.User.Id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	t, err := token.SignedString([]byte(h.secret))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, exceptions.InternalServerError{
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, AuthResponse{
		Token: t,
	})
}
