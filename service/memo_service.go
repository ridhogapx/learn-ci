package service

import (
	"learn-ci/model"
)

type MemoService interface {
	GetMemo(model.GetMemoRequest) (model.GetMemoResponse, error)
	AddMemo(model.AddMemoRequest) (bool, error)
}
