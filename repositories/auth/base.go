package auth

import (
	"bookstore/models/db"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

type AuthRepo interface {
	Register(dataUser db.User) error
	CheckEmail(string) (db.User, error)
	SaveToken(string, string) (db.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepo {
	return &authRepo{
		db: db,
	}
}
