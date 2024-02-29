package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPost := os.Getenv("DB_PORT")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPost + ")/" + dbName + "?parseTime=true"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	Db.AutoMigrate(&Product{}, &User{})
}
