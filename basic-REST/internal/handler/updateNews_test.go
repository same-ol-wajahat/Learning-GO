package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"newsapi/internal/handler"
	"strings"
	"testing"
)

func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewStorer
		expectedStatus int
	}{
		{
			name:           "invalid json request body",
			body:           strings.NewReader(`{`),
			store:          MockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request body",
			body: strings.NewReader(`
			{
			"id": "550e8400-e29b-41d4-a716-446655440000",
			"author": "son goku",
			"title": "DBZ",
			"summary": "phewwwww",
			"created_at": "2002-10-02T10:00:00-05:00",
			"source": "https://news-api.com"
			}
			`),
			store:          MockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "db error",
			body: strings.NewReader(`
			{
			"id": "550e8400-e29b-41d4-a716-446655440000",
			"author": "son goku",
			"title": "DBZ",
			"summary": "phewwwww",
			"created_at": "2002-10-02T10:00:00-05:00",
			"source": "https://news-api.com",
			"tags": ["anime"]
			}
			`),
			store:          MockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			body: strings.NewReader(`
			{
			"id": "550e8400-e29b-41d4-a716-446655440000",
			"author": "son goku",
			"title": "DBZ",
			"summary": "phewwwww",
			"created_at": "2002-10-02T10:00:00-05:00",
			"source": "https://news-api.com",
			"tags": ["anime"]
			}
			`),
			store:          MockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/", tc.body)

		// Act
		handler.UpdateNewsByID(tc.store)(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}
