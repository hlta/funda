package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandlingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			httpError, ok := err.(*echo.HTTPError)
			if !ok {
				httpError = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			code := httpError.Code
			message := httpError.Message
			if he, ok := message.(string); ok {
				return c.JSON(code, ErrorResponse{Code: code, Message: he})
			}
			return err
		}
		return nil
	}
}
