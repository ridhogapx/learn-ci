package service

import "learn-ci/repository"

type memoServiceImpl struct {
	Repository *repository.MemoRepository
}

func NewMemoService(repository *repository.MemoRepository) MemoService {
	return memoServiceImpl{
		Repository: repository,
	}
}
