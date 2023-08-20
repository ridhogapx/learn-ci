package service

import (
	"learn-ci/entity"
	"learn-ci/model"
	"learn-ci/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var memoRepository = repository.MemoRepositoryMock{Mock: mock.Mock{}}
var memoService = memoServiceImpl{Repository: &memoRepository}

func TestMemoService(t *testing.T) {
	expect := entity.MemoEntity{
		Title:  "My Activity",
		Body:   "Slept",
		Author: "rneko2006",
	}

	memoRepository.Mock.On("Read", "rneko2006").Return(expect)

	res, err := memoService.GetMemo(model.GetMemoRequest{
		Author: "rneko2006",
	})

	assert.Nil(t, err)
	assert.Equal(t, expect.Title, res.Title)
}
