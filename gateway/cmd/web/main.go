package main

import (
	"__gateway/pkg/auth"
	"__gateway/pkg/common/utils"
	v "__lib/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	grpcConn := utils.InitGrpcConn()
	uc := utils.InitGrpcUserClient(grpcConn)

	defer grpcConn.Close()

	e := echo.New()
	e.Validator = &v.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	auth.RegisterHandlers(e, uc)

	e.Logger.Fatal(e.Start(":8080"))
}
