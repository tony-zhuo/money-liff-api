package database

import (
	"fmt"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connection() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		//TODO: error handle
		logger := log.TeeDefault()
		logger.Error("db connection error: ", log.String("message", err.Error()))
		panic("db connection error.")
	}

	return db
}

func getDsn() string {
	if os.Getenv("APP_ENV") == "production" {
		return os.Getenv("DATABASE_URL")
	} else {
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Taipei",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_PASSWORD"),
		)
	}
}
