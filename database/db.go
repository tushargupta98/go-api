package database

import (
	"fmt"
	"github.com/tushargupta98/todoapi/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect : db.go")
		return nil, err
	}
	fmt.Println("DB Connection successful")

	DB = db
	_ = DB.AutoMigrate(&model.ToDo{})
	return db, nil
}
