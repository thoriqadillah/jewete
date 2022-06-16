package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

func GetInstance() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("env error : " + errEnv.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3303)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Error connecting database : " + err.Error())
	}

	instance = db
	return db
}

func Close() {
	dbCon, err := instance.DB()
	if err != nil {
		panic("Error closing database : " + err.Error())
	}

	dbCon.Close()
}
