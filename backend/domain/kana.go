package domain

type Kana struct {
	ID   uint   `gorm:"primaryKey"`
	Kana string `gorm:"unique"`
}
