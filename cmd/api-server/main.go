package main

import (
	"github.com/alexpojman/go-starter/internal/config"
	"github.com/alexpojman/go-starter/internal/logger"
	"github.com/alexpojman/go-starter/internal/queue"
	"github.com/alexpojman/go-starter/internal/routes"
	server "github.com/alexpojman/go-starter/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	// 1. Load Config
	c := config.NewConfig()
	err := c.LoadConfigFile(".", "env", ".env")

	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	// 2. Initialize Logger
	logger.InitLogger(c.GetString("ENVIRONMENT"))
	
	queue.Queue()
	
	e := server.InitEchoServer()

	e.GET("/", routes.Hello)
	e.POST("/save", routes.SaveUser)
	e.Logger.Fatal(e.Start(":1323"))
}