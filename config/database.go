package config

import (
	"fmt"
	"log"
	"os"

	"github.com/baguseka01/golang_echo_RESTAPI/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Not connect to databases")
	} else {
		log.Println("Success connect to database")
	}

	DB.AutoMigrate(
		&models.User{},
	)
}

// database
// var DB *gorm.DB

// func Connect() {
// 	dsn := "root:Nomorsatu01@tcp(localhost:3306)/golang_echo_resapi?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	DB.AutoMigrate(&models.User{})
// }
