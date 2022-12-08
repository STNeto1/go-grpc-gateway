package lib

import (
	"__lib/exceptions"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ParseValidatorErrors(err error) []exceptions.ValidationError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]exceptions.ValidationError, len(ve))
		for i, fe := range ve {
			out[i] = exceptions.ValidationError{Param: fe.Field(), Message: exceptions.MsgForTag(fe)}
		}

		return out
	}

	return nil
}
