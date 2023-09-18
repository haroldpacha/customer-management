package repository

import (
	"github.com/haroldpacha/customer-management/entity"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	InsertDepartment(department entity.Department) entity.Department
	UpdateDepartment(department entity.Department) entity.Department
	ShowDepartment(departmentID uint64) entity.Department
	DeleteDepartment(department entity.Department)
	AllDepartments() []entity.Department
}

type departmentConnection struct {
	connection *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentConnection{
		connection: db,
	}
}

func (db *departmentConnection) InsertDepartment(department entity.Department) entity.Department {
	db.connection.Save(&department)
	db.connection.Find(&department)
	return department
}

func (db *departmentConnection) UpdateDepartment(department entity.Department) entity.Department {
	var tempDepartment entity.Department
	db.connection.Find(&tempDepartment, department.ID)

	/* if department.Amount == 0 {
		department.Amount = tempDepartment.Amount
	}
	if department.Total == 0 {
		department.Total = tempDepartment.Total
	} */

	db.connection.Save(&department)
	return department
}

func (db *departmentConnection) ShowDepartment(departmentID uint64) entity.Department {
	var department entity.Department
	db.connection.Find(&department, departmentID)
	return department
}

func (db *departmentConnection) DeleteDepartment(department entity.Department) {
	db.connection.Delete(&department)
}

func (db *departmentConnection) AllDepartments() []entity.Department {
	var departments []entity.Department
	db.connection.Find(&departments)
	return departments
}
