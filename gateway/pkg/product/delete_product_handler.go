package product

import (
	"__lib/auth"
	"__lib/exceptions"
	productpb "__product/gen/pb/v1"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h handler) deleteProduct(c echo.Context) error {
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

	res, err := h.productClient.DeleteProduct(c.Request().Context(), &productpb.DeleteProductRequest{
		Id:     id.String(),
		UserId: usr.Id,
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
