package config

import (
	"learn-ci/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	dsn := "host=10.10.10.3 user=root password=root dbname=memo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Memo{})
	return db
}
