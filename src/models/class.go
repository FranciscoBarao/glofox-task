package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model `json:"-"`
	Name       string     `json:"name" valid:"required, alphanum, maxstringlength(50)"`
	StartDate  CustomTime `json:"start_date" valid:"required" gorm:"embedded"`
	EndDate    CustomTime `json:"end_date" valid:"required" gorm:"embedded"`
	Capacity   int        `json:"capacity" valid:"required, int, range(1,100)"`
}

func (class *Class) NewClass(name string, startDate CustomTime, endDate CustomTime, capacity int) Class {
	return Class{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}
}
