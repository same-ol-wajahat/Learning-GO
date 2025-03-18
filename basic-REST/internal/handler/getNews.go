package handler

import (
	"encoding/json"
	"net/http"
	"newsapi/internal/logger"

	"github.com/google/uuid"
)

func GetAllNews(ns NewStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request reccieved")

		news, err := ns.FindAll()
		if err != nil {
			logger.Error("failed to fetch all teh news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		allNewsResponse := AllNewsRespose{News: news}
		if err := json.NewEncoder(w).Encode(allNewsResponse); err != nil {
			logger.Error("falied to write response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetNewsByID(ns NewStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id is not valid uuid", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		news, err := ns.FindByID(newsUUID)
		if err != nil {
			logger.Error("news not found", "newsID", newsID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&news); err != nil {
			logger.Error("failed to encode", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
