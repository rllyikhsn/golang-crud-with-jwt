package db

type User struct {
	Base
	Email    string
	Password string `gorm:"type:text"`
	Token    string `gorm:"type:text"`
}
