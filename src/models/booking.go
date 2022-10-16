package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model `json:"-"`
	Name       string     `json:"name" valid:"required, alphanum, maxstringlength(50)"`
	Date       CustomTime `json:"date" valid:"required" gorm:"embedded;embeddedPrefix:date_"` // embedded allows the use of CustomTime as a field (So dates can be yyyy-mm-dd)
	ClassID    uint       `json:"-"`
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
