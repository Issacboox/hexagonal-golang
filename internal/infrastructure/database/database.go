package database

import (
	"bam/internal/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.User{}, &model.Product{}); err != nil {
		return nil, err
	}

	return db, nil
}
