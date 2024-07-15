package doctor_model

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Age      int8
	Address  string
	Phone    string
	Position string
}
