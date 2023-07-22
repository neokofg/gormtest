package handler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gormtest/application/utils"
	"gormtest/domains"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем данные из запроса
		var req RegisterRequest
		json.NewDecoder(r.Body).Decode(&req)
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Unable to hash password!", 400)
		}
		var user domains.User
		user.Name = req.Name
		user.Surname = req.Surname
		user.Email = req.Email
		user.Password = string(passwordHash)
		// ...
		var response interface{}
		// Создаем нового пользователя
		if err := db.Create(&user).Error; err != nil {
			w.WriteHeader(500)
			response = map[string]interface{}{
				"message": "failed to create user",
			}
		} else {
			w.WriteHeader(200)
			token, err := utils.GenerateJWT(user)
			if err != nil {
				http.Error(w, "Cannot create JWT token!"+err.Error(), 400)
				return
			}
			response = map[string]interface{}{
				"message": "user created",
				"token":   token,
			}
		}

		jsonResponse, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var req LoginRequest
		json.NewDecoder(r.Body).Decode(&req)

		var user domains.User

		user, err := domains.GetUserByEmail(db, req.Email)
		if err != nil {
			http.Error(w, "Cannot find user with this email!", 401)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(200)

		token, err := utils.GenerateJWT(user)
		if err != nil {
			http.Error(w, "Cannot create JWT token!"+err.Error(), 400)
			return
		}
		response := map[string]interface{}{
			"message": "user founded",
			"token":   token,
		}
		jsonResponse, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
	})
}
