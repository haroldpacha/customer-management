package service

import (
	"log"
	"time"

	"github.com/haroldpacha/customer-management/dto"
	"github.com/haroldpacha/customer-management/entity"
	"github.com/haroldpacha/customer-management/repository"
	"github.com/mashingan/smapping"
)

type CustomerService interface {
	Insert(b dto.CustomerCreateDTO) entity.Customer
	Update(b dto.CustomerUpdateDTO) entity.Customer
	Get(customerID uint64) entity.Customer
	Delete(b entity.Customer)
	All() []entity.Customer
	IsAllowedToEdit(userID string, customerID uint64) bool
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (service *customerService) Insert(b dto.CustomerCreateDTO) entity.Customer {
	newCustomer := entity.Customer{}
	err := smapping.FillStruct(&newCustomer, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	newCustomer.Age = service.getAge(newCustomer.Birthdate)

	res := service.customerRepository.InsertCustomer(newCustomer)
	return res
}

func (service *customerService) Update(b dto.CustomerUpdateDTO) entity.Customer {
	customer := entity.Customer{}
	err := smapping.FillStruct(&customer, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	customer.Age = service.getAge(customer.Birthdate)

	res := service.customerRepository.UpdateCustomer(customer)
	return res
}

func (service *customerService) Get(customerID uint64) entity.Customer {
	return service.customerRepository.GetCustomer(customerID)
}

func (service *customerService) Delete(b entity.Customer) {
	service.customerRepository.DeleteCustomer(b)
}

func (service *customerService) All() []entity.Customer {
	return service.customerRepository.AllCustomers()
}

func (service *customerService) IsAllowedToEdit(userID string, customerID uint64) bool {
	return true
}

func (service *customerService) getAge(birthdate time.Time) uint {
	currentDate := time.Now()
	age := currentDate.Year() - birthdate.Year()

	if currentDate.Month() < birthdate.Month() || (currentDate.Month() == birthdate.Month() && currentDate.Day() < birthdate.Day()) {
		age--
	}

	return uint(age)
}
