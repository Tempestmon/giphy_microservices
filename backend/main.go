package main

import (
	"log/slog"
	"os"
)

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
	)
}

func main() {
	log := setupLogger()
}
