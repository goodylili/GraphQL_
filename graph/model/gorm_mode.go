package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

type BioData struct {
	ID   int    `gorm:"primary key;autoIncrement" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}


func Database(DNS string) {
	// function body goes here
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&BioData{})
	if err != nil {
		return
	}

}