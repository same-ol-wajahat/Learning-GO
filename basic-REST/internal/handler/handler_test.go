package handler_test

import (
	"errors"
	"newsapi/internal/handler"

	"github.com/google/uuid"
)

type MockNewsStore struct {
	errState bool
}

func (m MockNewsStore) Create(_ handler.NewsPostReqBody) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m MockNewsStore) FindByID(_ uuid.UUID) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m MockNewsStore) FindAll() (news []handler.NewsPostReqBody, err error) {
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

func (m MockNewsStore) UpdateByID(_ handler.NewsPostReqBody) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}
