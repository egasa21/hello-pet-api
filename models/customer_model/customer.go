package customer_model

import (
	"github.com/egasa21/hello-pet-api/models/user_model"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID         uint            `gorm:"primaryKey;autoIncrement:false"`
	UserID     uint            `gorm:"not null"`
	User       user_model.User `gorm:"foreignKey:UserID"`
	Name       string
	Address    string
	Phone      string
	AnimalType string
}
