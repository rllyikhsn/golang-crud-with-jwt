package product

import (
	productRepo "bookstore/repositories/product"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productRepo productRepo.ProductRepo
}

type ProductController interface {
	ProductCreate(c *gin.Context)
	ProductList(c *gin.Context)
	ProductDetail(c *gin.Context)
	ProductUpdate(c *gin.Context)
	ProductDelete(c *gin.Context)
}

func NewProductController(productRepo productRepo.ProductRepo) ProductController {
	return &productController{
		productRepo: productRepo,
	}
}
