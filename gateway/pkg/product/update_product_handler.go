package product

import (
	"__lib/auth"
	"__lib/exceptions"
	v "__lib/validator"
	productpb "__product/gen/pb/v1"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UpdateProductRequestBody struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
}

func (h handler) updateProduct(c echo.Context) error {
	strId := c.Param("id")
	id, err := uuid.Parse(strId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, exceptions.BadRequest{
			Message:    "Invalid id",
			StatusCode: http.StatusBadRequest,
		})
	}

	usrId := auth.GetUserIDFromToken(c)

	usr, err := auth.GetUser(usrId, h.userClient, c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, exceptions.Unauthorized{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		})
	}

	body := new(UpdateProductRequestBody)
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

	res, err := h.productClient.UpdateProduct(c.Request().Context(), &productpb.UpdateProductRequest{
		Id:          id.String(),
		UserId:      usr.Id,
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, exceptions.InternalServerError{
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		})
	}
	if !res.Success {
		return echo.NewHTTPError(http.StatusBadRequest, exceptions.InternalServerError{
			Message:    "Product was not found",
			StatusCode: http.StatusBadRequest,
		})
	}

	return c.NoContent(204)
}
