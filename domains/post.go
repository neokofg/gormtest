package domains

import "gorm.io/gorm"

type Post struct {
	ID          uint `gorm:"primaryKey"`
	AuthorID    uint
	Name        string
	Text        string
	Description string
}

func GetPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
