package main

import (
	"log/slog"
	"net/http"
	"newsapi/internal/logger"
	"newsapi/internal/router"
	"os"
)

func main() {

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	r := router.New(nil)
	wrappedRouter := logger.AddLoggerMid(log, logger.LoggerMid(r))

	log.Info("Server starting on port 3000")

	if err := http.ListenAndServe(":3000", wrappedRouter); err != nil {
		log.Error("Failed to to Start the Server", "error", err)
	}
}
