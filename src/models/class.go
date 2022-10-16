package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model `json:"-" swaggerignore:"true"`
	Name       string     `json:"name" example:"Aerobics" valid:"required, alphanum, maxstringlength(50)"`
	StartDate  CustomTime `json:"start_date" example:"2022-01-01" swaggertype:"primitive,string" valid:"required" gorm:"embedded;embeddedPrefix:start_date_"`
	EndDate    CustomTime `json:"end_date" example:"2022-01-03" swaggertype:"primitive,string" valid:"required" gorm:"embedded;embeddedPrefix:end_date_"`
	Capacity   int        `json:"capacity" example:"10" valid:"required, int, range(1|100)"`
	Bookings   []Booking  `json:"bookings,omitempty"`
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

func (class *Class) GetStartDate() string {
	return class.StartDate.Format("2006-01-02")
}

func (class *Class) GetEndDate() string {
	return class.EndDate.Format("2006-01-02")
}

func (class *Class) IsOverbooking(count int) bool {
	return count > class.Capacity
}
