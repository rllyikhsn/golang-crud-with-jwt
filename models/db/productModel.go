package db

type Product struct {
	Base
	Name           string
	Gambar         string `gorm:"type:text"`
	ProductDetails []ProductDetail
}
