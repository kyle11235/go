package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo" // https://echo.labstack.com/guide
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("body=\n\n%s", reqBody)
	}))

	// Routes
	e.POST("/", fn)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}

func fn(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

/*
go run main.go
*/
