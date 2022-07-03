package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;uniqueIndex" json:"email"`
	Tasks     []Task `json:"tasks"`
}
