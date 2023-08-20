package model

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Title  string
	Body   string
	Author string
}

type GetMemoRequest struct {
	Author string `json:"author"`
}

type GetMemoResponse struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type AddMmemoRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}
