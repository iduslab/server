package db

import (
	"fmt"
	"log"
	"time"

	"github.com/gangjun06/iduslab/structure"
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

func AddBox(text, description string) error {
	result := db.Create(&structure.Box{Text: text, Description: description, Timestamp: time.Now()})
	return result.Error
}

func GetAllBox() (data []structure.Box, err error) {
	result := db.Where("").Find(&data)
	err = result.Error
	return
}

func GetBox(id int) (data structure.Box, err error) {
	result := db.Where("id = ?", id).First(&data)
	err = result.Error
	return
}

func AddMemo(box int, userid string, anon bool, text string) error {
	result := db.Create(&structure.Note{BoxNum: box, Author: userid, Anon: anon, Text: text, Timestamp: time.Now()})
	return result.Error
}

func PickMemo(box, count int) (data []structure.Note, err error) {
	result := db.Where("box_num = ?", box).Order("RANDOM()").Limit(count).Find(&data)
	err = result.Error
	return
}
