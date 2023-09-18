package dto

type DepartmentCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type DepartmentUpdateDTO struct {
	ID   uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}
