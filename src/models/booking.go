package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model `json:"-" swaggerignore:"true"`
	Name       string     `json:"name" example:"John" valid:"required, alphanum, maxstringlength(50)"`
	Date       CustomTime `json:"date"  example:"2022-01-03" swaggertype:"primitive,string" valid:"required" gorm:"embedded;embeddedPrefix:date_"` // embedded allows the use of CustomTime as a field (So dates can be yyyy-mm-dd)
	ClassID    uint       `json:"-" swaggerignore:"true"`
}

func (booking *Booking) NewBooking(name string, date CustomTime) Booking {
	return Booking{
		Name: name,
		Date: date,
	}
}

func (booking *Booking) GetDate() string {
	return booking.Date.Format("2006-01-02")
}

func (booking *Booking) SetClassID(classID uint) {
	booking.ClassID = classID
}
