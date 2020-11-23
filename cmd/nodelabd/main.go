package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/no-de-lab/nodelab-server/internal/logger"

	log "github.com/sirupsen/logrus"
)

var (
	addr         = flag.String("addr", "0.0.0.0:9090", "Default HTTP address")
	idleTimeout  = time.Second * 120
	writeTimeout = time.Second * 30
	readTimeout  = time.Second * 30
)

func main() {
	flag.Parse()

	container := InitializeDIContainer()
	config := container.Config

	logLevel := config.Log.Level
	sentryDSN := config.Log.SentryDSN
	phase := config.Phase.Level

	err := logger.InitLogging(logLevel, phase, sentryDSN)
	if err != nil {
		log.Errorf("Failed to setup logger with sentry")
	}

	mainRouter := mux.NewRouter()

	for _, h := range container.Handlers {
		h.SetupRoutes(mainRouter)
	}

	//Setting timeouts and handlers for http server
	s := &http.Server{
		Addr:         *addr,
		Handler:      mainRouter,
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	// Without graceful shutdown, the server might shutdown while handling important requests
	go func() {
		log.Printf("Running server on port %s", *addr)
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	// Notify when there is a os interrupt/kill command
	signal.Notify(sigChan, os.Interrupt)

	// Block the channel here, waiting to receive that os.Interrupt or os.Kill
	sig := <-sigChan
	log.Println("Received terminate signal, gracefully shutting down", sig)

	// Wait for 30 seconds for all handlers to finish
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shut down after completing all hanging requests
	_ = s.Shutdown(tc)
}
