package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gormtest/domains"
)

func InitializeDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=kiku port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connected")
	}
	db.AutoMigrate(
		&domains.User{},
		&domains.Author{},
		&domains.Post{},
	)
	return db
}
