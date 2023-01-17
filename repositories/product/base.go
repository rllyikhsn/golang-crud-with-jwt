package product

import (
	"bookstore/models/db"
	"bookstore/models/request"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

type ProductRepo interface {
	ProductCreate(db.Product) (db.Product, error)
	ProductList() ([]db.Product, error)
	ProductDetail(string) (db.Product, error)
	ProductUpdate(db.Product, request.Product) (db.Product, error)
	ProductDelete(string) (db.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}
