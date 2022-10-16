package controllers

import (
	"glofox-task/repositories"
)

// Contains all other controllers
type Controllers struct {
	BookingController *BookingController
	ClassController   *ClassController
}

// InitControllers returns a new set of controllers
func InitControllers(repos *repositories.Repositories) *Controllers {
	return &Controllers{
		BookingController: InitBookingController(repos.BookingRepository),
		ClassController:   InitClassController(repos.ClassRepository),
	}
}
