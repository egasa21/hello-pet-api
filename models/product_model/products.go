package product_model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"not null"`
	Price float32
	Stock int16
}
