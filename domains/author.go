package domains

type Author struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE"`
	Posts  []Post
}
