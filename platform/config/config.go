package config

import (
	"fmt"
	"os"
	"strings"
)

type Environment string

const (
	EnvironmentLocal      Environment = "local"
	EnvironmentStaging    Environment = "staging"
	EnvironmentProduction Environment = "production"
)

type Config struct {
	AppName     string
	Environment Environment
}

func Load() (Config, error) {
	cfg := Config{
		AppName:     getEnv("FLOWFORGE_APP_NAME", "flowforge"),
		Environment: Environment(getEnv("FLOWFORGE_ENV", string(EnvironmentLocal))),
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) Validate() error {
	if strings.TrimSpace(c.AppName) == "" {
		return fmt.Errorf("app name must not be empty")
	}

	switch c.Environment {
	case EnvironmentLocal, EnvironmentStaging, EnvironmentProduction:
		return nil
	default:
		return fmt.Errorf("unsupported environment %q", c.Environment)
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}
