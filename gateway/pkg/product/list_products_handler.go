package product

import (
	"__lib/exceptions"
	productpb "__product/gen/pb/v1"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) listProducts(c echo.Context) error {
	term := c.QueryParam("term")

	products, err := h.productClient.ListProduct(c.Request().Context(), &productpb.ListProductRequest{
		Term: term,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, exceptions.InternalServerError{
			Message:    "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		})
	}

	if len(products.Products) == 0 {
		return c.JSON(http.StatusOK, []productpb.Product{})
	}

	return c.JSON(http.StatusOK, products.Products)
}
