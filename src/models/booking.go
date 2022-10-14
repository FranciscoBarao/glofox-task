package models

type Booking struct {
	Name string     `json:"name" valid:"required, alphanum, maxstringlength(50)"`
	Date CustomTime `json:"date" valid:"required"`
}

func (booking *Booking) NewBooking(name string, date CustomTime) Booking {
	return Booking{
		Name: name,
		Date: date,
	}
}
