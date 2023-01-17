package db

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	CreatedBy int8
	UpdatedAt time.Time
	UpdatedBy int8
	DeletedAt time.Time
	DeletedBy int8
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
