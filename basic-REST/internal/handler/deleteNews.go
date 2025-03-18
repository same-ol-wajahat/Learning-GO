package handler

import (
	"net/http"
	"newsapi/internal/logger"

	"github.com/google/uuid"
)

func DeleteNewsByID(ns NewStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request recieved")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id is not valid uuid", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := ns.DeleteByID(newsUUID); err != nil {
			logger.Error("news not found", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
