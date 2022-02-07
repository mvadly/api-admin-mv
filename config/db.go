package config

import (
	entity "api-adminmv/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	InitEnv()
	Username := os.Getenv("DB_USERNAME")
	Password := os.Getenv("DB_PASSWORD")
	Hostname := os.Getenv("DB_HOSTNAME")
	DbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Hostname, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database not connected")
	}

	db.AutoMigrate(
		entity.UserBasicAuth{},
		entity.User{},
	)

	return db

}

func CloseDB() {
	sql, err := ConnectDB().DB()
	if err != nil {
		fmt.Println(err)
	}

	sql.Close()
}
