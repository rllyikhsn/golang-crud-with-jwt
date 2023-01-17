package main

import (
	"bookstore/config/db"
	"bookstore/config/initializers"
	authController "bookstore/controllers/auth"
	productController "bookstore/controllers/product"
	"bookstore/middleware"
	authRepository "bookstore/repositories/auth"
	productRepository "bookstore/repositories/product"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDbMySql()
}

func main() {

	dbManajemenBuku := db.ConnectToDb()

	// List All Repositories
	authRepository := authRepository.NewAuthRepository(&dbManajemenBuku)
	productRepository := productRepository.NewProductRepository(&dbManajemenBuku)

	// List All Controllers
	authController := authController.NewAuthController(authRepository)
	productController := productController.NewProductController(productRepository)

	r := gin.Default()

	// List All Route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Route Prouct
	r.GET("/api/product", middleware.RequireAuth, productController.ProductList)
	r.GET("/api/product/:id", middleware.RequireAuth, productController.ProductDetail)
	r.POST("/api/product", middleware.RequireAuth, productController.ProductCreate)
	r.PUT("/api/product/:id", middleware.RequireAuth, productController.ProductUpdate)
	r.DELETE("/api/product/:id", middleware.RequireAuth, productController.ProductDelete)

	r.POST("/register", authController.AuthRegister)
	r.POST("/login", authController.AuthLogin)

	r.Run() // listen and serve on 0.0.0.0:8080
}
