package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(connection string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		fmt.Printf("\n\n\n\nError Connecting To DataBase error: %s\n\n\n", err)
	}

	return db
}
