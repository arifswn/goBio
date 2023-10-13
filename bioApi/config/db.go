package config

import (
	"fmt"
	"goBio/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	host := "localhost"
	port := "3306"
	dbname := "go_bio"
	username := "root"
	password := "lupaLagi"

	//dsn (data source name) untuk menghubungkan ke database
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic("Failed connect to database : " + err.Error())
	}

	// first migrate
	AutoMigrate()
	fmt.Println("Database Connected...")
}

func GetDB() *gorm.DB {
	return DB
}

func AutoMigrate() {
	// auto migrate models to database table
	if err := DB.Debug().AutoMigrate(&model.Users{}); err != nil {
		panic("Error migrate database : " + err.Error())
	}
}
