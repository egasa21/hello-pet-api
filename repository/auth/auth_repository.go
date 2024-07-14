package repository

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/egasa21/hello-pet-api/models/user_model"
)

type AuthRepository struct {
	db *database.DB
}

func NewAuthRepository(db *database.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (repo *AuthRepository) Register(user *user_model.User) (err error) {
	err = repo.db.Database.Create(user).Error
	return
}

func (repo *AuthRepository) FindUserByEmail(user *user_model.User, userReq string) (err error) {
	err = repo.db.Database.Where("email = ?", userReq).First(&user).Error
	return
}
