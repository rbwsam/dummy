package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type response struct {
	Hostname string `json:"hostname"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		hostname, err := os.Hostname()
		if err != nil {
			return err
		}

		r := response{
			Hostname: hostname,
		}
		return c.JSONPretty(http.StatusOK, r, "  ")
	})

	e.Logger.Fatal(e.Start(":80"))
}
