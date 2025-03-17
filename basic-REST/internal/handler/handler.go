package handler

import (
	"encoding/json"
	"net/http"
	"newsapi/internal/logger"

	"github.com/google/uuid"
)

type NewStorer interface {
	// create news
	Create(NewsPostReqBody) (NewsPostReqBody, error)
	// Find new by its ID
	FindByID(uuid.UUID) (NewsPostReqBody, error)
	// FindAll returns all news in the store
	FindAll() ([]NewsPostReqBody, error)
	// Delete news by its specific ID
	DeleteByID(uuid.UUID) error
	// Update news By its specific ID
	UpdateByID(uuid.UUID, NewsPostReqBody) error
}

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

func GetAllNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func GetNewsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func UpdateNewsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func DeleteNewsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
