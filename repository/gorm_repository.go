package repository

import "github.com/egasa21/hello-pet-api/infra/database"

type GormRepository struct {
	db *database.DB
}

func NewGormRepository(db *database.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

type Repository interface {
	GetById(id uint, dest interface{}) error
}
