package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model `json:"-"`
	Name       string     `json:"name" valid:"required, alphanum, maxstringlength(50)"`
	Date       CustomTime `json:"date" valid:"required" gorm:"embedded"` // embedded allows the use of CustomTime as a field (So dates can be yyyy-mm-dd)
}

func (booking *Booking) NewBooking(name string, date CustomTime) Booking {
	return Booking{
		Name: name,
		Date: date,
	}
}
