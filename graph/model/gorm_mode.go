package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type BioData struct {
	ID   int    `gorm:"primary key;autoIncrement" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Database() *gorm.DB {
	// function body goes here
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&BioData{})
	if err != nil {
		log.Fatal("failed to migrate")
	}

	return db
}

func NewBookService(db *gorm.DB) *BioService {
	return &BioService{
		db: db,
	}
}
