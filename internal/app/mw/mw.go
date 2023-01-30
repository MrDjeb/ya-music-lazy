package mw

import (
	"log"

	"github.com/labstack/echo/v4"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		headRole := ctx.Request().Header.Get("User-Role")

		if headRole == "admin" {
			log.Print("Admin detected!")
		}

		if err := next(ctx); err != nil {
			return err
		}
		return nil
	}
}
