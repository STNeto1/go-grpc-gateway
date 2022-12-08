package product

import (
	"__lib/auth"
	"__lib/exceptions"
	v "__lib/validator"
	productpb "__product/gen/pb/v1"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CreateProductRequestBody struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
}

func (h handler) createProduct(c echo.Context) error {
	id := auth.GetUserIDFromToken(c)

	usr, err := auth.GetUser(id, h.userClient, c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, exceptions.Unauthorized{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		})
	}

	body := new(CreateProductRequestBody)
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

	success, err := h.productClient.CreateProduct(c.Request().Context(), &productpb.CreateProductRequest{
		UserId:      usr.Id,
		UserName:    usr.Name,
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	})
	if err != nil || !success.Success {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, exceptions.InternalServerError{
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		})
	}

	return c.NoContent(201)
}
