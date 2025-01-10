// cmd/sso/main.go
package main

import (
	"fmt"
	"log/slog"
	"os"

	"sso/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	// TEMP solution
	fmt.Println("Config:", cfg)
	fmt.Println("Logger:", log)
}

// setupLogger creates a logger based on the given environment.
//
// The returned logger always logs to stdout. The log format and level
// depend on the environment:
//
// - local: text format, debug level
// - dev: JSON format, debug level
// - prod: JSON format, info level
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
