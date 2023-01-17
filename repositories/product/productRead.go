package product

import "bookstore/models/db"

func (pr *productRepo) ProductList() ([]db.Product, error) {
	var products []db.Product
	result := pr.db.Preload("ProductDetails").Order("id desc").Find(&products)

	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func (pr *productRepo) ProductDetail(id string) (db.Product, error) {
	var product db.Product
	result := pr.db.Preload("ProductDetails").First(&product, id)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}
