package main

import (
	"log/slog"
	"net/http"
	"os"
	"tasks/internal/di"
)

func main() {
	router, port, err := di.Initialize()
	if err != nil {
		slog.Error("Failed to initialize dependencies", "ERROR", err)
		os.Exit(1)
	}

	slog.Info("Server is running", "port", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
