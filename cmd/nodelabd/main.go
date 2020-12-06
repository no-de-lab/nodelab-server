package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/no-de-lab/nodelab-server/api"
	"github.com/no-de-lab/nodelab-server/internal/logger"
	log "github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
	"net/http"
	"time"
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

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", container.Config.Server.Port),
		ReadTimeout:  time.Second * 120,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 30,
	}

	log.Fatal(e.StartServer(server))
	log.Fatal(graceful.ListenAndServe(server, time.Second*30))
}
