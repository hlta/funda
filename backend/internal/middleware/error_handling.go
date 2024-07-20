package middleware

import (
	"funda/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandlingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			var fieldErrors []model.FieldError

			switch e := err.(type) {
			case model.FieldError:
				fieldErrors = append(fieldErrors, e)
				return c.JSON(http.StatusConflict, ErrorResponse{
					Code:    http.StatusConflict,
					Message: "Validation Error",
					Errors:  fieldErrors,
				})
			default:
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
		}
		return nil
	}
}
