package logger

import (
	"io"
	"log/slog"

	"github.com/NontNtw/flowforge/platform/config"
)

type Options struct {
	Writer      io.Writer
	AppName     string
	Environment config.Environment
	Level       slog.Level
}

func New(options Options) *slog.Logger {
	handler := slog.NewJSONHandler(
		options.Writer,
		&slog.HandlerOptions{
			Level: options.Level,
		},
	)

	return slog.New(handler).With(
		slog.String("app", options.AppName),
		slog.String("environment", string(options.Environment)),
	)
}