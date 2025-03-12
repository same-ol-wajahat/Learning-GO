package main

import (
	"log/slog"
	"net/http"
	"newsapi/internal/router"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: false}))

	logger.Info("Server starting on port 3000")

	r := router.New()

	if err := http.ListenAndServe(":3000", r); err != nil {
		logger.Error("Failed to to Start the Server", "error", err)
	}
}
