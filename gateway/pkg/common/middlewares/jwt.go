package middlewares

// echo middleware to check jwt

import (
	"__lib/exceptions"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWT(secret string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get header from header
			header := c.Request().Header.Get("Authorization")
			if header == "" {
				return c.JSON(http.StatusUnauthorized, exceptions.Unauthorized{
					Message:    "Unauthorized",
					StatusCode: http.StatusUnauthorized,
				})
			}

			splits := strings.Split(header, " ")
			if len(splits) != 2 {
				return c.JSON(http.StatusUnauthorized, exceptions.Unauthorized{
					Message:    "Unauthorized",
					StatusCode: http.StatusUnauthorized,
				})
			}

			// validate token
			claims, err := jwt.Parse(splits[1], func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, exceptions.Unauthorized{
					Message:    "Unauthorized",
					StatusCode: http.StatusUnauthorized,
				})
			}

			if !claims.Valid {
				return c.JSON(http.StatusUnauthorized, exceptions.Unauthorized{
					Message:    "Unauthorized",
					StatusCode: http.StatusUnauthorized,
				})
			}

			// set claims in context
			c.Set("claims", claims)

			return next(c)
		}
	}
}
