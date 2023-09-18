package repository

import (
	"github.com/haroldpacha/customer-management/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	InsertCustomer(customer entity.Customer) entity.Customer
	UpdateCustomer(customer entity.Customer) entity.Customer
	GetCustomer(customerID uint64) entity.Customer
	DeleteCustomer(customer entity.Customer)
	AllCustomers() []entity.Customer
}

type customerConnection struct {
	connection *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerConnection{
		connection: db,
	}
}

func (db *customerConnection) InsertCustomer(customer entity.Customer) entity.Customer {
	db.connection.Save(&customer)
	db.connection.Preload("Department").Find(&customer)
	return customer
}

func (db *customerConnection) UpdateCustomer(customer entity.Customer) entity.Customer {
	db.connection.Save(&customer)
	db.connection.Preload("Department").Find(&customer)
	return customer
}

func (db *customerConnection) GetCustomer(customerID uint64) entity.Customer {
	var customer entity.Customer
	db.connection.Preload("Department").Find(&customer, customerID)
	return customer
}

func (db *customerConnection) DeleteCustomer(customer entity.Customer) {
	db.connection.Delete(&customer)
}

func (db *customerConnection) AllCustomers() []entity.Customer {
	var customers []entity.Customer
	db.connection.Preload("Department").Find(&customers)
	return customers
}
