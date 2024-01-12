// 実行時にCGO_CFLAGS="-D_LARGEFILE64_SOURCE"を設定
// https://github.com/mattn/go-sqlite3/issues/1164のissue対策

package main

import (
	"cointrading/app/config"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
	mygorm "cointrading/app/infrastructure/gorm"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func init() {
	config.LoggingSettings(config.LogFileName())
}

func main() {
	switch config.ORM() {
	case 1:
		gormMigrate()
	default:
		gormMigrate()
	}

	log.Println("Migrate successfully")
}

func gormMigrate() {
	var db *gorm.DB

	switch config.DB() {
	case 1:
		db = mygorm.SQLiteConnection()
	default:
		db = mygorm.SQLiteConnection()
	}

	exchanges := valueobjects.Exchanges()
	symbols := valueobjects.Symbols()
	durations := valueobjects.Durations()

	originDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get origin DB: %s", err)
	}

	for _, e := range exchanges {
		for _, s := range symbols {
			for _, d := range durations {
				candle := entities.NewCandle(e, s, d, nil, 0, 0, 0, 0, 0)

				tableName := candle.GetTableName()

				cmd := fmt.Sprintf(`
				CREATE TABLE IF NOT EXISTS %s (
					time DATETIME PRIMARY KEY NOT NULL,
					open FLOAT,
					close FLOAT,
					high FLOAT,
					low FLOAT,
					volume FLOAT
					)`, tableName)

				log.Println(cmd)

				if _, err := originDB.Exec(cmd); err != nil {
					log.Fatalf("Failed to execute query: %s", err)
				}
			}
		}
	}

	// if err := db.AutoMigrate(); err != nil {
	// 	log.Fatalf("Failed to migrate: %s", err)
	// }
}
