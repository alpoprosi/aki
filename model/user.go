package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Phone    string
	Password string

	LastName string
	Patronic string

	IsLandLord bool
	IsAdmin    bool
}

func (User) TableName() string {
	return "users"
}
