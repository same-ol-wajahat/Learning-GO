package handler_test

import (
	"newsapi/internal/handler"
	"testing"
)

func TestNewsPostReqBody_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		req         handler.NewsPostReqBody
		expectedErr bool
	}{
		{
			name:        "auther empty",
			req:         handler.NewsPostReqBody{},
			expectedErr: true,
		},
		{
			name: "title empty",
			req: handler.NewsPostReqBody{
				Author: "son wukong",
			},
			expectedErr: true,
		},
		{
			name: "summary empty",
			req: handler.NewsPostReqBody{
				Author: "son goku",
				Title:  "DBZ",
			},
			expectedErr: true,
		},
		{
			name: "time invalid",
			req: handler.NewsPostReqBody{
				Author:    "madara",
				Title:     "naruto",
				Summary:   "dattay bayu",
				CreatedAt: "invalid time",
			},
			expectedErr: true,
		},
		{
			name: "source invalid",
			req: handler.NewsPostReqBody{
				Author:    "madara",
				Title:     "naruto",
				Summary:   "dattay bayu",
				CreatedAt: "2002-10-02T10:00:00-05:00",
				Source:    "http://test-news.com",
			},
			expectedErr: true,
		},
		{
			name: "tags empty",
			req: handler.NewsPostReqBody{
				Author:    "madara",
				Title:     "naruto",
				Summary:   "dattay bayu",
				CreatedAt: "2002-10-02T10:00:00-05:00",
				Source:    "http://test-news.com",
			},
			expectedErr: true,
		},
		{
			name: "valid",
			req: handler.NewsPostReqBody{
				Author:    "madara",
				Title:     "naruto",
				Summary:   "dattay bayu",
				CreatedAt: "2002-10-02T10:00:00-05:00",
				Source:    "http://test-news.com",
				Tags:      []string{"test-tags"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()

			if tc.expectedErr && err == nil {
				t.Fatal("expected error but got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Fatal("expected nil but got error")
			}
		})
	}
}
