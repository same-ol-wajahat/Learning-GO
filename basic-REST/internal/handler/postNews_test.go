package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"newsapi/internal/handler"
	"strings"
	"testing"
)

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewStorer
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			body:           strings.NewReader(`{`),
			store:          MockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid reuqest body",
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
			name: "Success",
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
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", tc.body)

			// Act
			handler.PostNews(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
