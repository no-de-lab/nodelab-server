package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/no-de-lab/nodelab-server/api"
	"github.com/no-de-lab/nodelab-server/graphql"
	"github.com/no-de-lab/nodelab-server/internal/logger"
	log "github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

func main() {
	flag.Parse()

	container := InitializeDIContainer()
	rootResolver := InitializeResolver()
	logger.InitLogging()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}     ${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = api.ErrorHandler

	graphql.SetupGraphQL(e, rootResolver)

	for _, h := range container.Handlers {
		h.SetupRoutes(e)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", container.Config.Server.Port),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 120,
	}

	log.Fatal(e.StartServer(server))
	log.Fatal(graceful.ListenAndServe(server, time.Second*30))
}
