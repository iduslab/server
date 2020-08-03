package db

import (
	"fmt"
	"log"

	"github.com/gangjun06/bot01/structure"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func Init() {
	connection, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln(err)
		return
	}

	db = connection
	fmt.Println("Successful to Connect database")

	var models = []interface{}{&structure.Box{}, &structure.Note{}}
	connection.AutoMigrate(models...)
	fmt.Println("Successfully performed AutoMigrate")
}

func CloseDB() {
	db.Close()
}

func AddBox(text string) error {
	result := db.Create(&structure.Box{Text: text})
	return result.Error
}

func GetAllBox() (data []structure.Box, err error) {
	result := db.Where("").Find(&data)
	err = result.Error
	return
}
