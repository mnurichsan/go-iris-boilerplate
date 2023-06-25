package config

import (
	"fmt"
	"iris-boilerplate/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(configuration Config) *gorm.DB {

	dbName := configuration.Get("DB_NAME")
	dbUser := configuration.Get("DB_USER")
	dbPassword := configuration.Get("DB_PASSWORD")
	dbHost := configuration.Get("DB_HOST")
	dbPort := configuration.Get("DB_PORT")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		fmt.Println("err", err.Error())
	}

	fmt.Println("Berhasil Connect ke DB")

	DB.AutoMigrate(&entity.User{})

	fmt.Println("Migration Berhasil")

	return DB
}
