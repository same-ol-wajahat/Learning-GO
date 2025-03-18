package handler_test

import (
	"net/http"
	"net/http/httptest"
	"newsapi/internal/handler"
	"testing"

	"github.com/google/uuid"
)

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewStorer
		expectedStatus int
	}{
		{
			name:           "insternal server error",
			store:          MockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          MockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		// Act
		handler.GetAllNews(tc.store)(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}

func Test_GetNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewStorer
		newsID         string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			store:          MockNewsStore{},
			newsID:         "inavlid-uuid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			store:          MockNewsStore{errState: true},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          MockNewsStore{},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.SetPathValue("news_id", tc.newsID)

		// Act
		handler.GetNewsByID(tc.store)(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}
