package data

import (
	"gbl-api/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSqliteDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.SqliteFilename), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
