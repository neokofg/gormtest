package domains

type Post struct {
	ID          uint   `gorm:"primaryKey"`
	Author      Author `gorm:"embedded"`
	Name        string
	Text        string
	Description string
}
