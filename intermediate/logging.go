package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	// Basic log
	log.Println("Starting application...")

	// File logging
	f, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(f)
	log.Println("This is logged to a file")
	defer f.Close()

	// Structured logging with slog
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("User login", "user", "radoslaw", "status", "success")
	logger.Error("Database connection failed", "retry", true)
	logger.Warn("Disk space low", "remaining", "500MB")

	// JSON logging with slog
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("Service started", "port", 8080)
	jsonLogger.Error("Failed to process request", "error", "timeout")
	jsonLogger.Warn("High memory usage", "usage", "1.5GB")

	// Custom log level
	type customLevel int
	const (
		LevelDebug customLevel = iota
		LevelInfo
		LevelWarn
		LevelError
	)

	var levelMap = map[customLevel]slog.Level{
		LevelDebug: slog.LevelDebug,
		LevelInfo:  slog.LevelInfo,
		LevelWarn:  slog.LevelWarn,
		LevelError: slog.LevelError,
	}

	customLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: levelMap[LevelInfo],
	}))

	customLogger.Info("This is an info message")
	customLogger.Warn("This is a warning message")
	customLogger.Error("This is an error message")
	customLogger.Debug("This debug message will not be shown")

}
