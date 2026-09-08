package config

import (
	"os"
	"testing"
)

func TestGetEnvUsesFallbackWhenUnset(t *testing.T) {
	const key = "FLOWFORGE_TEST_UNSET"

	previous, existed := os.LookupEnv(key)

	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("unset environment variable: %v", err)
	}

	t.Cleanup(func() {
		if existed {
			_ = os.Setenv(key, previous)
			return
		}

		_ = os.Unsetenv(key)
	})

	got := getEnv(key, "fallback")

	if got != "fallback" {
		t.Fatalf("expected fallback value, got %q", got)
	}
}

func TestGetEnvUsesEnvironmentValue(t *testing.T) {
	const key = "FLOWFORGE_TEST_VALUE"

	t.Setenv(key, "configured")

	got := getEnv(key, "fallback")

	if got != "configured" {
		t.Fatalf("expected configured value, got %q", got)
	}
}

func TestLoadFromEnvironment(t *testing.T) {
	t.Setenv("FLOWFORGE_APP_NAME", "flowforge-api")
	t.Setenv("FLOWFORGE_ENV", "staging")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.AppName != "flowforge-api" {
		t.Fatalf("expected app name %q, got %q", "flowforge-api", cfg.AppName)
	}

	if cfg.Environment != EnvironmentStaging {
		t.Fatalf(
			"expected environment %q, got %q",
			EnvironmentStaging,
			cfg.Environment,
		)
	}
}

func TestLoadRejectsEmptyAppName(t *testing.T) {
	t.Setenv("FLOWFORGE_APP_NAME", "")
	t.Setenv("FLOWFORGE_ENV", "local")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestLoadRejectsUnsupportedEnvironment(t *testing.T) {
	t.Setenv("FLOWFORGE_APP_NAME", "flowforge")
	t.Setenv("FLOWFORGE_ENV", "development")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
