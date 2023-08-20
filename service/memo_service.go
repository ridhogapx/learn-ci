package service

import "learn-ci/entity"

type MemoService interface {
	GetMemo(author string) (*entity.MemoEntity, error)
	AddMemo(entity.MemoEntity) (bool, error)
}
