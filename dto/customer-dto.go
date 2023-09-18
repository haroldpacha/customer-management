package dto

import "time"

type CustomerUpdateDTO struct {
	ID           uint64    `json:"id" form:"id" binding:"required"`
	DNI          string    `json:"dni" form:"dni" binding:"required,min=8,max=8"`
	Name         string    `json:"name" form:"name" binding:"required"`
	LastName     string    `json:"last_name" form:"last_name" binding:"required"`
	Birthdate    time.Time `json:"birthdate" form:"birthdate" binding:"required"`
	DepartmentID uint64    `json:"department_id,omitempty" form:"department_id,omitempty"`
	Address      string    `json:"address" form:"address" binding:"required"`
}

type CustomerCreateDTO struct {
	DNI          string    `json:"dni" form:"dni" binding:"required,min=8,max=8"`
	Name         string    `json:"name" form:"name" binding:"required"`
	LastName     string    `json:"last_name" form:"last_name" binding:"required"`
	Birthdate    time.Time `json:"birthdate" form:"birthdate" binding:"required"`
	DepartmentID uint64    `json:"department_id,omitempty" form:"department_id,omitempty"`
	Address      string    `json:"address" form:"address" binding:"required"`
}
