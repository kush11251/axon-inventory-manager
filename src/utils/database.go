package utils

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDatabase() (
	*gorm.DB,
	error,
) {
	var err error
	db, err = gorm.Open("sqlite3", "./inventory.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
