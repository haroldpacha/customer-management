package entity

import "time"

type Customer struct {
	ID           uint64     `gorm:"primary_key:auto_increment" json:"id"`
	DNI          string     `gorm:"type:varchar(8);unique" json:"dni"`
	Name         string     `gorm:"type:varchar(120)" json:"name"`
	LastName     string     `gorm:"type:varchar(120)" json:"last_name"`
	Birthdate    time.Time  `gorm:"type:date" json:"birthdate"`
	Age          uint       `json:"age"`
	DepartmentID uint64     `gorm:"not null" json:"-"`
	Department   Department `gorm:"foreignkey:DepartmentID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"department"`
	Address      string     `gorm:"type:varchar(255)" json:"address"`
}
