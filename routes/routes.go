package routes

import (
	"github.com/rs/cors"
	"gorm.io/gorm"
	"gormtest/application/handler"
	"log"
	"net/http"
)

func InitializeRoutes(db *gorm.DB) {
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
