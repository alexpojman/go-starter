package main

import (
	"os"

	"github.com/alexpojman/go-starter/internal/queue"
	"github.com/alexpojman/go-starter/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)
	queue.Queue(logger)
	
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")
	
			return nil
		},
	}))

	e.GET("/", routes.Hello)
	e.POST("/save", routes.SaveUser)
	e.Logger.Fatal(e.Start(":1323"))
}