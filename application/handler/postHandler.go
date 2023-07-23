package handler

import (
	"encoding/json"
	"gorm.io/gorm"
	"gormtest/application/utils"
	"gormtest/domains"
	"net/http"
)

type CreatePostRequest struct {
	Name        string `json:"name"`
	Text        string `json:"text"`
	Description string `json:"description"`
}

func CreatePost(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreatePostRequest
		json.NewDecoder(r.Body).Decode(&req)
		tokenString := r.Header.Get("Authorization")
		user_id, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Cannot Authorize!"+err.Error(), 400)
			return
		}
		user, err := domains.GetUserById(db, user_id)
		tx := db.Begin()

		// Создание автора
		var author domains.Author

		if tx.Where("user_id = ?", user.ID).First(&author).Error != nil {
			author = domains.Author{User: user}
			tx.Create(&author)
		}
		// Создание поста
		post := domains.Post{
			AuthorID: author.ID,
			Name:     req.Name,
			Text:     req.Text,
		}
		tx.Create(&post)

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}
		response := map[string]interface{}{
			"message": "post created",
			"post":    post,
		}
		jsonResponse, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
	})
}
func GetPosts(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		posts, err := domains.GetPosts(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		response := map[string]interface{}{
			"message": "posts founded",
			"posts":   posts,
		}
		jsonResponse, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
	})
}
