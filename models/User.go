package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string         `gorm:"not null"`
	LastName  string         `gorm:"not null"`
	Email     string         `gorm:"not null;unique_index"`
	Telephone string         `gorm:"not null"`
	Age       string         `gorm:"not null"`
	Dni       string         `gorm:"not null;unique"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
