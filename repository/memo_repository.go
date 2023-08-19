package repository

import "learn-ci/entity"

type MemoRepository interface {
	Save(memo *entity.MemoEntity) error
	Read(author string) (*entity.MemoEntity, error)
}
