package handler

import (
	"errors"
	"fmt"
	"net/url"
	"newsapi/internal/store"
	"time"

	"github.com/google/uuid"
)

type NewsPostReqBody struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt string    `json:"created_at"`
	Content   string    `json:"content"`
	Source    string    `json:"source"`
	Tags      []string  `json:"tags"`
}

type AllNewsRespose struct {
	News []store.News `json:"news"`
}

func (n NewsPostReqBody) Validate() (news store.News, errs error) {
	if n.Author == "" {
		errs = errors.Join(errs, fmt.Errorf("author is empty: %s", n.Author))
	}
	if n.Title == "" {
		errs = errors.Join(errs, fmt.Errorf("title is empty: %s", n.Title))
	}
	if n.Summary == "" {
		errs = errors.Join(errs, fmt.Errorf("summary is empty: %s", n.Summary))
	}
	t, err := time.Parse(time.RFC3339, n.CreatedAt)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	url, err := url.Parse(n.Source)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	if len(n.Tags) == 0 {
		errs = errors.Join(errs, fmt.Errorf("tags should not be empty %s", n.Tags))
	}

	if errs != nil {
		return news, nil
	}
	return store.News{
		ID:        n.ID,
		Author:    n.Author,
		Title:     n.Title,
		Summary:   n.Summary,
		CreatedAt: t,
		Source:    url,
		Tags:      n.Tags,
	}, nil
}
