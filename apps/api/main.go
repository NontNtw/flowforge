package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/NontNtw/flowforge/platform/config"
	"github.com/NontNtw/flowforge/platform/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load configuration: %v", err)
	}

	log := logger.New(logger.Options{
		Writer:      os.Stdout,
		AppName:     cfg.AppName,
		Environment: cfg.Environment,
		Level:       slog.LevelInfo,
	})

	log.Info("api starting")
}
