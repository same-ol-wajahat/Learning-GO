package handler_test

import (
	"errors"
	"newsapi/internal/store"

	"github.com/google/uuid"
)

type MockNewsStore struct {
	errState bool
}

func (m MockNewsStore) Create(_ store.News) (news store.News, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m MockNewsStore) FindByID(_ uuid.UUID) (news store.News, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m MockNewsStore) FindAll() (news []store.News, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m MockNewsStore) DeleteByID(_ uuid.UUID) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

func (m MockNewsStore) UpdateByID(_ store.News) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}
