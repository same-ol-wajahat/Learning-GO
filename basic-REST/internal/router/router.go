package router

import (
	"net/http"
	"newsapi/internal/handler"
)

// Create and return new router with all routes configured
func New() *http.ServeMux {

	r := http.NewServeMux()

	// Get all news
	r.HandleFunc("GET /news", handler.GetAllNews())
	// Get specific news by id
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID())
	// Update specific news by id
	r.HandleFunc("POST /news", handler.PostNews())
	// Update News By id
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID())
	// Delete specific news by id
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID())

	return r
}
