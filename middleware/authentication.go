package middleware

import (
	"net/http"

	"echo/helpers"

	"github.com/labstack/echo/v4"
)

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			verifyToken, err := helpers.VerifyToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Unauthenticated",
					"message": err.Error(),
				})
			}

			// Log or print verifyToken for debugging
			// fmt.Println(verifyToken["id"])

			c.Set("userData", verifyToken)

			return next(c)
		}
	}
}
