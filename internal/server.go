package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() error {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: dbconn
	// dbconn, err := db.NewConn()
	// if err != nil {
	// 	 return err
	// }

	// TODO: handler
	// handlers := interface.NewHandlers(dbconn)
	// oas.RegisterHandlers(e, handlers)

	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
