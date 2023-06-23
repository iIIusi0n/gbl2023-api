package data

import (
	"gbl-api/config"
	"gbl-api/controllers/booth"
	"gbl-api/controllers/problem"
	"gbl-api/controllers/score"
	"gbl-api/controllers/user"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {
	if db != nil {
		return db
	}

	switch config.DbType {
	case "sqlite":
		db = getSqliteDatabase()
	default:
		panic("Unknown database type")
	}

	db.AutoMigrate(&booth.Booth{})
	db.AutoMigrate(&score.Participation{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&problem.Problem{})

	return db
}
