package repository

import "learn-ci/entity"

type MemoRepository interface {
	Save(*entity.MemoEntity) error
	Read(author string) *entity.MemoEntity
}
