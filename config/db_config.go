package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbConfig() (*gorm.DB, error) {
	errorEnvFIle := godotenv.Load(".env")
	if errorEnvFIle != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	userName := os.Getenv("DB_USERNAME")
	databaseName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	dsn := getDbURLConnection(userName+":"+password+"@tcp("+host+":"+port+")/"+databaseName+"?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func getDbURLConnection(url string) string {
	return url
}
