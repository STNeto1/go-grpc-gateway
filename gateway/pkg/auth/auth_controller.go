package auth

import (
	"__gateway/pkg/common/middlewares"
	userpb "__user/gen/pb/user/v1"

	"github.com/labstack/echo/v4"
)

type handler struct {
	userClient userpb.UserServiceClient

	secret string
}

func RegisterHandlers(e *echo.Echo, uc userpb.UserServiceClient) {
	h := handler{
		userClient: uc,
		secret:     "some secret",
	}

	authGroup := e.Group("/auth")

	authGroup.POST("/login", h.login)
	authGroup.POST("/register", h.register)
	authGroup.GET("/profile", h.profile, middlewares.JWT(h.secret))
}
