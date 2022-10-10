package config

import (
	"assessment_efishery/entity"
	"fmt"

	//"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Database() {
	dsn := "host=localhost user=postgres password='' dbname=efishery_academy search_path=oms port=5432 TimeZone=Asia/Jakarta sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database success connected")
}

func Migrate() {
	DB.AutoMigrate(&entity.Products{})
	DB.AutoMigrate(&entity.Users{})
	DB.AutoMigrate(&entity.Cart{})
	DB.AutoMigrate(&entity.Transactions{})
}
