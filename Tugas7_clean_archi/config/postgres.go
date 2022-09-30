package config

import (
	"clean-archi/entity"
	"fmt"

	//"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Database() {
	dsn := "host=localhost user=postgres password='' dbname=efishery_academy port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database success connected")
}

func Migrate() {
	DB.AutoMigrate(&entity.User{})
}
