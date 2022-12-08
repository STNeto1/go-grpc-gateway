package product

import (
	"__gateway/pkg/common/middlewares"
	productpb "__product/gen/pb/v1"
	userpb "__user/gen/pb/v1"

	"github.com/labstack/echo/v4"
)

type handler struct {
	userClient    userpb.UserServiceClient
	productClient productpb.ProductServiceClient
}

func RegisterHandlers(e *echo.Echo, userClient userpb.UserServiceClient, productClient productpb.ProductServiceClient) {
	h := &handler{
		userClient:    userClient,
		productClient: productClient,
	}

	e.POST("/products/create", h.createProduct, middlewares.JWT())
	e.GET("/products/list", h.listProducts)
	e.GET("/products/show/:id", h.getProduct)
	e.PUT("/products/update/:id", h.updateProduct, middlewares.JWT())
	e.DELETE("/products/delete/:id", h.deleteProduct, middlewares.JWT())
	//e.GET("/products", h.GetProducts
}
