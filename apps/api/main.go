package main

import (
	"log"

	"github.com/NontNtw/flowforge/platform/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load configuration: %v", err)
	}

	log.Printf(
		"%s api starting environment=%s",
		cfg.AppName,
		cfg.Environment,
	)
}
