package gorm

import (
	"cointrading/app/config"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteConnection() *gorm.DB {
	path := fmt.Sprintf("%s/db/cointrading.db", config.DockerWorkDir())

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
