package repository

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/egasa21/hello-pet-api/models/product_model"
)

type ProductRepository struct {
	db *database.DB
}

func NewProductRepository(db *database.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) Create(product *product_model.Product) (err error) {
	err = repo.db.Database.Create(product).Error
	return
}

func (repo *ProductRepository) FindById(product *product_model.Product, productId string) (err error) {
	err = repo.db.Database.First(&product, "id=?", productId).Error
	return
}

func (repo *ProductRepository) Update(product *product_model.Product) (err error) {
	err = repo.db.Database.Save(product).Error
	return
}

func (repo *ProductRepository) Delete(productId string) (err error) {
	err = repo.db.Database.Delete(&product_model.Product{}, "id=?", productId).Error
	return
}
