package handler_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"newsapi/internal/handler"
	"strings"
	"testing"

	"github.com/google/uuid"
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
			store:          mockNewsStore{},
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
			store:          mockNewsStore{},
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
			store:          mockNewsStore{errState: true},
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
			store:          mockNewsStore{},
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

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		// Act
		handler.GetAllNews()(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}

func Test_GetNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		// Act
		handler.GetNewsByID()(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}

func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/", nil)

		// Act
		handler.UpdateNewsByID()(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}

func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		// Arrange
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/", nil)

		// Act
		handler.DeleteNewsByID()(w, r)

		// Assert
		if w.Result().StatusCode != tc.expectedStatus {
			t.Errorf("expected: %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
		}
	}
}

type mockNewsStore struct {
	errState bool
}

func (m mockNewsStore) Create(_ handler.NewsPostReqBody) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindByID(_ uuid.UUID) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindAll() (news []handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) DeleteByID(_ uuid.UUID) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

func (m mockNewsStore) UpdateByID(_ uuid.UUID, _ handler.NewsPostReqBody) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}
