package logger

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/no-de-lab/nodelab-server/container"
	log "github.com/sirupsen/logrus"
)

func InitLogging() {
	config := container.Container.Config

	level := config.Log.Level
	dsn := config.Log.SentryDSN
	phase := config.Phase.Level

	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}

	client, err := raven.New(dsn)
	if err != nil {
		log.Panicf("Failed to initialize raven client %s", err)
	}

	if phase == "local" {
		return
	}

	hook, err := logrus_sentry.NewWithClientSentryHook(client, []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	})

	if err != nil {
		log.Panicf("Failed to initialize sentry hooks, err: %s", err)
	} else {
		log.AddHook(hook)
	}
}
