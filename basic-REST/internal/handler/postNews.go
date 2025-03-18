package handler

import (
	"encoding/json"
	"net/http"
	"newsapi/internal/logger"
)

func PostNews(ns NewStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request reccieved")

		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the reuqest", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := newsRequestBody.Validate(); err != nil {
			logger.Error("reuquest validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(newsRequestBody); err != nil {
			logger.Error("error creating the news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}
