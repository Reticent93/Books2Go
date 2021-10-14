package main

import (
	"fmt"
	"github.com/Reticent93/Books2Go/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("user=%s", "pass=%s", "dbname=%s", "host=%s", "sslmode=disable", username, password, dbName, dbHost)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}

	db = conn
	db.Debug().AutoMigrate(&models.Book{}, &models.Author{})
}
