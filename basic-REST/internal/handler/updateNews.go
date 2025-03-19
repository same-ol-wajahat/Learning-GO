package handler

import (
	"encoding/json"
	"net/http"
	"newsapi/internal/logger"
)

func UpdateNewsByID(ns NewStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")

		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		n, err := newsRequestBody.Validate()
		if err != nil {
			logger.Error("request valiadation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := ns.UpdateByID(n); err != nil {
			logger.Error("error updating the news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
