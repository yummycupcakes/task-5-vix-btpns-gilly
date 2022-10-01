package database

import (
	"fmt"
	"log"

	"github.com/yummycupcakes/task-5-vix-btpns-gilly/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )

var (
	DB_USERNAME = "root"
 	DB_PASSWORD = ""
	DB_NAME = "vixbtpns"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	db       *gorm.DB
)
  
func DatabaseConnection() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	db.AutoMigrate(models.Photo{}, models.User{})
	println("Database connected!")
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	dbSQL.Close()
}

func GetDB() *gorm.DB {
	return db
}