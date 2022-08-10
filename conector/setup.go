package conector

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	user := os.Getenv("DB_USER")

	pwd := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/logindata?charset=utf8mb4&parseTime=True&loc=Local", user, pwd)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&LoginData{})

	DB = db
}

func (LoginData) TableName() string {
	return "logindata"
}
