package db

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db global postgres instance
var DB *gorm.DB

func init() {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	DB, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to DB")
		panic(err)
	}
	return
}
