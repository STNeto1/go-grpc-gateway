package product

import (
	"__lib/exceptions"
	productpb "__product/gen/pb/v1"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h handler) getProduct(c echo.Context) error {
	strId := c.Param("id")
	id, err := uuid.Parse(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, exceptions.BadRequest{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid id",
		})
	}

	product, err := h.productClient.GetProduct(c.Request().Context(), &productpb.GetProductRequest{
		Id: id.String(),
	})
	if err != nil {
		return c.JSON(http.StatusNotFound, exceptions.NotFound{
			StatusCode: http.StatusNotFound,
			Message:    "Product was not found",
		})
	}

	return c.JSON(http.StatusOK, product.Product)
}
