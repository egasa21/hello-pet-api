package repository

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/egasa21/hello-pet-api/models/customer_model"
)

type CustomerRepository struct {
	db *database.DB
}

func NewCustomerRepository(db *database.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (repo *CustomerRepository) CreateCustomer(customer *customer_model.Customer) (err error) {
	err = repo.db.Database.Create(customer).Error
	return
}

func (repo *CustomerRepository) GetCustomerById(customer *customer_model.Customer, customerId string) (err error) {
	err = repo.db.Database.Preload("User").First(&customer, "id=?", customerId).Error
	return
}

func (repo *CustomerRepository) UpdateCustomer(customer *customer_model.Customer) (err error) {
	err = repo.db.Database.Save(customer).Error
	return
}

func (repo *CustomerRepository) DeleteCustomer(customerId string) (err error) {
	err = repo.db.Database.Delete(&customer_model.Customer{}, "id=?", customerId).Error
	return
}

func (repo *CustomerRepository) GetAllCustomers(customers []customer_model.Customer) (err error) {
	err = repo.db.Database.Find(&customers).Error
	return
}

func (repo *CustomerRepository) LoadUser(customer *customer_model.Customer) (err error) {
	err = repo.db.Database.Model(customer).Association("User").Find(&customer.User)
	return
}
