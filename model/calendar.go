package model

import (
	"time"

	"gorm.io/gorm"
)

type Calendar struct {
	gorm.Model
	ObjectID uint
	Settings CalendarSettings

	Orders []Order `gorm:"joinReferences:order_id"`
}

func (Calendar) TableName() string {
	return "calendars"
}

type Order struct {
	gorm.Model
	CalendarID  uint
	UserID      uint
	Name        string
	Description string
	Books       []Book `gorm:"serializer:json"`
}

func (Order) TableName() string {
	return "orders"
}

type Book struct {
	Start time.Time
	End   time.Time
	Price float64
}

type CalendarSettings struct {
	gorm.Model
	WorkDays     WorkDays
	WorkHours    []Hour
	Holidays     []time.Time
	WeekdaysOnly bool
	EachDayWork  bool
}

type WorkDays struct {
	Type DaysType
	Days []Day
}

type DaysType int

var (
	DaysTypeWeek  = 0
	DaysTypeMonth = 1
)

type Day struct {
	Start uint
	End   uint
}

type Hour struct {
	Start uint
	End   uint
}

func DefaultCalendarSettings() CalendarSettings {
	return CalendarSettings{}
}
