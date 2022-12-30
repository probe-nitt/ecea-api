package middlewares

import "github.com/labstack/echo/v4"

func Responder(c echo.Context, code int, response interface{}) error {
	return c.JSON(code, map[string]interface{}{
		"response": response,
	})
}
