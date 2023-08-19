package model

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Title  string
	Body   string
	Author string
}
