package app

import (
	entyity "ardiman-xyz/go-todo-app/models/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB() *gorm.DB {
	dsn := "ardimansql:mypassword@tcp(127.0.0.1:3306)/belajar_golang_restful_api?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(&entyity.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}