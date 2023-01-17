package product

import (
	"bookstore/models/db"
	"bookstore/models/request"
	"time"
)

func (pr *productRepo) ProductUpdate(product db.Product, requestProduct request.Product) (db.Product, error) {

	product.Name = requestProduct.Name
	product.Gambar = requestProduct.Gambar
	product.UpdatedAt = time.Now()

	for _, v := range product.ProductDetails {
		for _, r := range requestProduct.Details {
			if v.ID == uint(r.Id) {
				v.UpdatedAt = time.Now()
				v.Name = r.Name
				v.Stock = int32(r.Stock)
				v.Gambar = r.Gambar
				pr.db.Updates(v)
			}
		}
	}

	result := pr.db.Preload("ProductDetails").Updates(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

func (pr *productRepo) ProductDelete(id string) (db.Product, error) {
	var product db.Product
	result := pr.db.Preload("ProductDetails").First(&product, id)

	if result.Error != nil {
		return product, result.Error
	}

	pr.db.Preload("ProductDetails").Model(&product).Delete(&product)

	return product, nil
}
