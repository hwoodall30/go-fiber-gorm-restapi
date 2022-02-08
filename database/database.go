package database

import (
	// todo "gofiber/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DATABASE_URI string = "test.db"

func Connect() error {
    var err error
    DB, err = gorm.Open(sqlite.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
    })

    if err != nil {
        panic(err)
    }

	log.Println("Connected to database")

    // DB.AutoMigrate(&todo.Todo{})
    // log.Println("Database migrated")

    return nil
}