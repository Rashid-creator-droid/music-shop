package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("musicshop.db?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	return DB
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("failed to get sql.DB:", err)
		return
	}
	sqlDB.Close()
}
