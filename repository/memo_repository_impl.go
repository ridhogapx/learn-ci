package repository

import (
	"learn-ci/entity"
	"learn-ci/model"

	"gorm.io/gorm"
)

type memoRepositoryImpl struct {
	Query *gorm.DB
}

func NewMemoRepository(db *gorm.DB) MemoRepository {
	return memoRepositoryImpl{
		Query: db,
	}
}

func (repository memoRepositoryImpl) Save(memo *entity.MemoEntity) error {
	newMemo := model.Memo{
		Title:  memo.Title,
		Body:   memo.Body,
		Author: memo.Author,
	}

	repository.Query.Create(&newMemo)
	return nil
}

func (repository memoRepositoryImpl) Read(author string) (*entity.MemoEntity, error) {
	var memo model.Memo
	result := repository.Query.Find(&memo, "author = ?", author)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity.MemoEntity{
		Title:  memo.Title,
		Body:   memo.Body,
		Author: memo.Author,
	}, nil

}
