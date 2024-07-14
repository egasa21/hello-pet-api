package user_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"type:varchar(255);unique"`
	Email    string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255);not null"`
}
