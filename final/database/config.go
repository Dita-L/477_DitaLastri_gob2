package database

import (
	"final/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	username = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "myGram"
)

func GetDB() *gorm.DB {
	//startDB
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:@tcp(%s)/%s?parseTime=true&loc=Local", username, hostname, dbname))
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.Socmed{})
	return db
}
