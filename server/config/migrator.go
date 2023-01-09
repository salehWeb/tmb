package config

import (
	"fmt"

	"github.com/salehWeb/chat-app/server/models"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		panic(fmt.Sprintf("\n\n\n\nError Migrate User To DataBase error: %s\n\n\n", err))
	}

}
