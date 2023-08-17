package domain

type Pattern struct {
	ID     uint `gorm:"primaryKey"`
	Roma   string
	KanaID uint
}
