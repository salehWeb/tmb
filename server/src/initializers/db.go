package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase()  {
	var connection string = os.Getenv("CONNECTION_STRING")

	if len(connection) < 2 {
		panic(fmt.Sprintf("No Connection String Found For DataBase"))
	}

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("\n Error Connecting To DataBase error: %s \n", err))
	}

	DB = db
}
