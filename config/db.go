package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_APP") //root:@tcp(127.0.0.1:3306)/gin-rest?charset=utf8mb4&parseTime=True&loc=Local
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed")
	}
	return db
}
