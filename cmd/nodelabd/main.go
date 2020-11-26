package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/no-de-lab/nodelab-server/api"
	"github.com/no-de-lab/nodelab-server/internal/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	flag.Parse()

	container := InitializeDIContainer()
	logger.InitLogging()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = api.ErrorHandler

	for _, h := range container.Handlers {
		h.SetupRoutes(e)
	}

	log.Fatal(e.Start(":9090"))

	sigChan := make(chan os.Signal, 1)
	// Notify when there is a os interrupt/kill command
	signal.Notify(sigChan, os.Interrupt)

	// Block the channel here, waiting to receive that os.Interrupt or os.Kill
	sig := <-sigChan
	log.Println("Received terminate signal, gracefully shutting down", sig)
}
