package service

import (
	"learn-ci/model"
	"learn-ci/repository"
)

type memoServiceImpl struct {
	Repository repository.MemoRepository
}

func NewMemoService(repository repository.MemoRepository) MemoService {
	return &memoServiceImpl{
		Repository: repository,
	}
}

func (service *memoServiceImpl) GetMemo(request model.GetMemoRequest) (model.GetMemoResponse, error) {
	author := request.Author
	req, err := service.Repository.Read(author)

	if err != nil {
		return nil, err
	}

	return model.GetMemoResponse{
		Title:  req.Title,
		Body:   req.Body,
		Author: req.Author,
	}, nil
}
