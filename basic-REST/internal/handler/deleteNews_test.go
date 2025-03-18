package handler_test

import (
	"net/http"
	"net/http/httptest"
	"newsapi/internal/handler"
	"testing"

	"github.com/google/uuid"
)

func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewStorer
		newsID         string
		expectedStatus int
	}{
		{
			name:           "invalid newsID",
			store:          MockNewsStore{},
			newsID:         "invalid UUID",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			store:          MockNewsStore{errState: true},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Success",
			store:          MockNewsStore{},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusNoContent,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		r.SetPathValue("news_id", tc.newsID)

		// Act
		handler.DeleteNewsByID(tc.store)(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}
