package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func database() *gorm.DB {
	db, err := gorm.Open("sqlite3", "images.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Create the Image model in the database
	db.AutoMigrate(&Image{})

	return db
}
