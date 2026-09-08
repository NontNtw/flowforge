package logger

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/NontNtw/flowforge/platform/config"
)

func TestNewAddsApplicationFields(t *testing.T) {
	var output bytes.Buffer

	log := New(Options{
		Writer:      &output,
		AppName:     "flowforge-api",
		Environment: config.EnvironmentStaging,
		Level:       slog.LevelInfo,
	})

	log.Info("application starting")

	var entry map[string]any

	if err := json.Unmarshal(output.Bytes(), &entry); err != nil {
		t.Fatalf("decode log entry: %v", err)
	}

	if got := entry["msg"]; got != "application starting" {
		t.Fatalf("expected message %q, got %q", "application starting", got)
	}

	if got := entry["app"]; got != "flowforge-api" {
		t.Fatalf("expected app %q, got %q", "flowforge-api", got)
	}

	if got := entry["environment"]; got != "staging" {
		t.Fatalf("expected environment %q, got %q", "staging", got)
	}

	if got := entry["level"]; got != "INFO" {
		t.Fatalf("expected level %q, got %q", "INFO", got)
	}
}

func TestNewRespectsLogLevel(t *testing.T) {
	var output bytes.Buffer

	log := New(Options{
		Writer:      &output,
		AppName:     "flowforge",
		Environment: config.EnvironmentLocal,
		Level:       slog.LevelWarn,
	})

	log.Info("ignored message")

	if output.Len() != 0 {
		t.Fatalf("expected info log to be filtered, got %q", output.String())
	}

	log.Warn("warning message")

	if output.Len() == 0 {
		t.Fatal("expected warning log")
	}
}