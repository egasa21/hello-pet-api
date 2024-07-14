package repository

import "github.com/egasa21/hello-pet-api/infra/database"

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

func GetUserByEmail() {

}
