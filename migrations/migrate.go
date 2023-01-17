package main

import (
	"bookstore/config/db"
	"bookstore/config/initializers"
	modelsDb "bookstore/models/db"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDbMySql()
}

func main() {
	db.DB.AutoMigrate(
		&modelsDb.User{},
		&modelsDb.Product{},
		&modelsDb.ProductDetail{},
	)
}
