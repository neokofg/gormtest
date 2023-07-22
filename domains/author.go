package domains

type Author struct {
	ID   uint `gorm:"primaryKey"`
	User User `gorm:"embedded"`
}
