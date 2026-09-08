package database

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestOpenRejectsEmptyDatabaseURL(t *testing.T) {
	_, err := Open(context.Background(), "")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "database URL must not be empty") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOpenRejectsInvalidDatabaseURL(t *testing.T) {
	_, err := Open(context.Background(), "://invalid")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "parse database configuration") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOpenPropagatesContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := Open(
		ctx,
		"postgres://flowforge:secret@localhost:5432/flowforge?sslmode=disable",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled error, got %v", err)
	}
}
