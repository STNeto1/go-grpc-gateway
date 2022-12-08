package main

import (
	"__gateway/pkg/auth"
	"__gateway/pkg/common/utils"
	"__gateway/pkg/product"
	v "__lib/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	uConn, uc := utils.InitGrpcUserClient()
	pConn, pc := utils.InitGrpcProductClient()

	defer uConn.Close()
	defer pConn.Close()

	e := echo.New()
	e.Validator = &v.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	auth.RegisterHandlers(e, uc)
	product.RegisterHandlers(e, uc, pc)

	e.Logger.Fatal(e.Start(":8080"))
}
