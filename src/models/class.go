package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model `json:"-"`
	Name       string     `json:"name" valid:"required, alphanum, maxstringlength(50)"`
	StartDate  CustomTime `json:"start_date" valid:"required" gorm:"embedded;embeddedPrefix:start_date_"`
	EndDate    CustomTime `json:"end_date" valid:"required" gorm:"embedded;embeddedPrefix:end_date_"`
	Capacity   int        `json:"capacity" valid:"required, int, range(1|100)"`
	Bookings   []Booking
}

func (class *Class) NewClass(name string, startDate CustomTime, endDate CustomTime, capacity int) Class {
	return Class{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}
}

func (class *Class) GetName() string {
	return class.Name
}

func (class *Class) GetID() uint {
	return class.ID
}

func (class *Class) IsOverbooking(count int) bool {
	return count > class.Capacity
}
