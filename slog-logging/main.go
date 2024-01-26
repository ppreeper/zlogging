package main

import (
	"log/slog"
	"os"
)

func main() {
	logs := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logs)
	// slog.Debug("Debug message")
	// slog.Info("Info message")
	// slog.Warn("Warning message")
	// slog.Error("Error message")
	tFunc()
}

func tFunc() {
	// slog.Debug("Debug message")
	// slog.Info("Info message")
	// slog.Warn("Warning message")
	// slog.Error("Error message")
}
