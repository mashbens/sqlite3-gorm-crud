package config

import (
	"sqlite3-gorm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		models.User{},
	)
}

func InitDB() *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	err = MigrateDB(db)
	if err != nil {
		panic(err)
	}

	return db
}
