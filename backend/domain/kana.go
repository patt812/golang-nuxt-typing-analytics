package domain

type Kana struct {
	ID   uint   `gorm:"primaryKey"`
	Kana string `gorm:"unique,collate:utf8mb4_bin"` // avoid case-insensitive unique constraint
}
