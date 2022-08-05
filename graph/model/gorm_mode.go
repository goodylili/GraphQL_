package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

)

type BioData struct {
	ID   int    `gorm:"primary key;autoIncrement" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}


func Database(DNS string) {

	// function body goes here
	DB, err = gorm.Open(sqlite.Open("test.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Human{})
	if err != nil {
		return
	}

}