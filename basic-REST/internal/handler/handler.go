package handler

import (
	"newsapi/internal/store"

	"github.com/google/uuid"
)

type NewStorer interface {
	// create news
	Create(store.News) (store.News, error)
	// Find new by its ID
	FindByID(uuid.UUID) (store.News, error)
	// FindAll returns all news in the store
	FindAll() ([]store.News, error)
	// Delete news by its specific ID
	DeleteByID(uuid.UUID) error
	// Update news By its specific ID
	UpdateByID(store.News) error
}
