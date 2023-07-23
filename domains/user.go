package domains

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Surname   string
	Email     string `gorm:"unique"`
	Password  string
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User

	if result := db.Where("email = ?", email).First(&user); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
func GetUserById(db *gorm.DB, id float64) (User, error) {
	var user User

	if result := db.Where("id = ?", id).First(&user); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
