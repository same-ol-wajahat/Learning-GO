package handler

import (
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
	UpdateByID(NewsPostReqBody) error
}
