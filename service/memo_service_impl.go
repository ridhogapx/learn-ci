package service

import (
	"learn-ci/entity"
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
		return model.GetMemoResponse{}, err
	}

	return model.GetMemoResponse{
		Title:  req.Title,
		Body:   req.Body,
		Author: req.Author,
	}, nil
}

func (service *memoServiceImpl) AddMemo(request model.AddMemoRequest) (bool, error) {
	err := service.Repository.Save(&entity.MemoEntity{
		Title:  request.Title,
		Body:   request.Body,
		Author: request.Author,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
