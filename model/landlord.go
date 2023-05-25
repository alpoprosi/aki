package model

import "gorm.io/gorm"

type LandLord struct {
	gorm.Model
	CategoryID    int64
	JuridicalName string
	RegistrarJob  string
	INN           string
	Descriprion   string
}

func (LandLord) TableName() string {
	return "landlords"
}
