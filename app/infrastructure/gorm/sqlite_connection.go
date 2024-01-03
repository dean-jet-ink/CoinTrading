package gorm

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../db/cointrading.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
