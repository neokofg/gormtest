package main

import (
	"fmt"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gormtest/application/handler"
	"gormtest/domains"
	"log"
	"net/http"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=kiku port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connected")
	}
	db.AutoMigrate(&domains.User{}, &domains.Author{}, &domains.Post{})
	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // your app's origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
	mux.Handle("/register", handler.Register(db))
	mux.Handle("/login", handler.Login(db))
	mux.Handle("/post", handler.CreatePost(db))
	mux.Handle("/posts", handler.GetPosts(db))
	corsHandler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":8000", corsHandler))
}
