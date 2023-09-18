package service

import (
	"log"

	"github.com/haroldpacha/customer-management/dto"
	"github.com/haroldpacha/customer-management/entity"
	"github.com/haroldpacha/customer-management/repository"
	"github.com/mashingan/smapping"
)

type DepartmentService interface {
	Insert(r dto.DepartmentCreateDTO) entity.Department
	Update(r dto.DepartmentUpdateDTO) entity.Department
	Show(departmentID uint64) entity.Department
	Delete(r entity.Department)
	All() []entity.Department
}

type departmentService struct {
	departmentRepository repository.DepartmentRepository
}

func NewDepartmentService(departmentRepository repository.DepartmentRepository) DepartmentService {
	return &departmentService{
		departmentRepository: departmentRepository,
	}
}

func (service *departmentService) Insert(r dto.DepartmentCreateDTO) entity.Department {
	newDepartment := entity.Department{}
	err := smapping.FillStruct(&newDepartment, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	return service.departmentRepository.InsertDepartment(newDepartment)
}

func (service *departmentService) Update(r dto.DepartmentUpdateDTO) entity.Department {
	department := entity.Department{}
	err := smapping.FillStruct(&department, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	return service.departmentRepository.UpdateDepartment(department)
}

func (service *departmentService) Show(departmentID uint64) entity.Department {
	return service.departmentRepository.ShowDepartment(departmentID)
}

func (service *departmentService) Delete(department entity.Department) {
	service.departmentRepository.DeleteDepartment(department)
}

func (service *departmentService) All() []entity.Department {
	return service.departmentRepository.AllDepartments()
}
