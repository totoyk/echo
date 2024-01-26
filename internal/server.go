package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Run() error {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
