package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"train-http/internal/handlers"
	"train-http/internal/utils"
)

func CheckAccess(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		cookie, err := c.Request().Cookie("JWT")
		if cookie == nil {
			return c.JSON(http.StatusUnauthorized, handlers.Response{
				Message: "JWT is Empty",
			})
		}
		checkToken := utils.ValidateJWT(cookie.Value)
		if err != nil || !checkToken {
			cookie.Value = ""
			return c.JSON(http.StatusUnauthorized, handlers.Response{
				Message: "Unauthorized",
			})
		}
		return next(c)
	}
}
