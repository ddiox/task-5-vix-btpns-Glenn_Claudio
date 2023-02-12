package database

import (
	"fmt"
	"log"

	"github.com/ddiox/task-5-vix-btpns-Glenn_Claudio/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "tcp(localhost:3306)"
	user     = "root"
	password = "123456"
	dbname   = "finalassignment"
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error saat menghubungkan ke database: ", err)
	}

	fmt.Println("Berhasil terhubung ke database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
