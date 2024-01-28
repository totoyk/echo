package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/totoyk/trial-api-golang/internal/infra/db"
	"github.com/totoyk/trial-api-golang/internal/oas"
	"github.com/totoyk/trial-api-golang/internal/ui"
)

func Run() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbconn, err := db.NewConn()
	if err != nil {
		return err
	}
	handlers := ui.NewHandlers(dbconn)
	oas.RegisterHandlers(e, handlers)

	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
