package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint    `gorm:"primary_key" json:"id"`
	Username  string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Roles     []Role  `gorm:"many2many:user_roles;" json:"roles"`
	Groups    []Group `gorm:"many2many:user_groups;" json:"groups"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
	DeletedAt int64   `json:"deleted_at"`
}

