package main

import (
	"github.com/alexpojman/go-starter/cmd/api-server/routes"
	"github.com/alexpojman/go-starter/internal/queue"
	"github.com/labstack/echo/v4"
)

func main() {
	queue.Queue()
	
	e := echo.New()
	e.HideBanner = true

	e.GET("/", routes.Hello)
	e.POST("/save", routes.SaveUser)
	e.Logger.Fatal(e.Start(":1323"))
}