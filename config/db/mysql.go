package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDbMySql() {
	var err error
	dsn := os.Getenv("DB_MYSQL_DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect Database MySql")
	}
}

func ConnectToDb() gorm.DB {
	return *DB
}
