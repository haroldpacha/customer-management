package entity

import "time"

type Department struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
