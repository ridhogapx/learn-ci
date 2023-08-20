package repository

import (
	"errors"
	"learn-ci/entity"

	"github.com/stretchr/testify/mock"
)

type MemoRepositoryMock struct {
	Mock mock.Mock
}

func (repository *MemoRepositoryMock) Save(req *entity.MemoEntity) error {
	arg := repository.Mock.Called(req)

	if arg.Get(0) == nil {
		return nil
	}

	return arg.Error(0)

}

func (repository *MemoRepositoryMock) Read(author string) (*entity.MemoEntity, error) {
	arg := repository.Mock.Called(author)

	if arg.Get(0) == nil {
		return nil, errors.New("not found")
	}

	return arg.Get(0).(*entity.MemoEntity), nil
}
