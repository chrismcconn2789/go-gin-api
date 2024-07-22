package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Transaction struct {
	gorm.Model
	Name   string `json:"name"`
	Amount uint   `json:"amount"`
}

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	err = database.AutoMigrate(&Transaction{})

	if err != nil {
		return
	}

	DB = database
}
