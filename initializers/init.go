package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/AvinFajarF/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	var errDB error

	dsn := os.Getenv("DB")
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDB != nil {
		log.Fatalln("Connection to database error")
	}

	DB.AutoMigrate(&model.User{}, &model.Task{})
}

func msg() {
	fmt.Println("Intialize everything...")
}
