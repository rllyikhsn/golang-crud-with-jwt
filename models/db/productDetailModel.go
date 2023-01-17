package db

type ProductDetail struct {
	Base
	ProductID uint
	Name      string
	Stock     int32
	Gambar    string `gorm:"type:text"`
}
