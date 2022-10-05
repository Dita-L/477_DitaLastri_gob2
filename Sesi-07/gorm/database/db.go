package database

import (
	"fmt"
	"gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "Dita2710"
	dbPort   = "5432"
	dbname   = "db_go"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user,
		password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	//method Open digunakan untuk membangun koneksi ke database.
	//ketika koneksi berhasil, maka variabel db akan mengandung referensi dari database dengan tipe data *gorm.DB
	//perlu diingat, variabel db adalah variabel global

	if err != nil {
		log.Fatal("error connecting to database :", err)

	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
	//method Debug digunakan sebagai debugging atau logger.
	//di chain dengan method AutoMigrate
	//AutoMigratedigunakan untuk migrasi otomatis dari struct-struct yang telah dibuat.

}

func GetDB() *gorm.DB {
	return db
}
