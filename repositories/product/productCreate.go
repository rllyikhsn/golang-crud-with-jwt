package product

import "bookstore/models/db"

func (pr *productRepo) ProductCreate(product db.Product) (db.Product, error) {

	result := pr.db.Create(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}
