package model

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	Name         string
	Description  string
	CalendarID   int64
	MapPlaceMark MapPlaceMark `gorm:"serializer:json"`
}

func (Object) TableName() string {
	return "objects"
}

type MapPlaceMark struct {
	Latitude  float64
	Longitude float64
}
