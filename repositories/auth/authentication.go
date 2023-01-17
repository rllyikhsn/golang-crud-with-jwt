package auth

import (
	"bookstore/models/db"
)

func (ar *authRepo) Register(dataUser db.User) error {

	result := ar.db.Create(&dataUser)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ar *authRepo) CheckEmail(email string) (db.User, error) {
	var user db.User
	result := ar.db.First(&user, "email = ?", email)

	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (ar *authRepo) SaveToken(email string, token string) (db.User, error) {
	var user db.User
	result := ar.db.First(&user, "email = ?", email)

	if result.Error != nil {
		return user, result.Error
	}

	user.Token = token
	ar.db.Updates(&user)
	return user, nil
}
