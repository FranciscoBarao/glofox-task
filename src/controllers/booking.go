package controllers

import (
	"net/http"

	"github.com/unrolled/render"

	"glofox-task/repositories"
)

// Declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type bookingRepository interface {
	Create() error
}

type BookingController struct {
	repo bookingRepository
}

func InitBookingController(bookingRepo *repositories.BookingRepository) *BookingController {
	return &BookingController{
		repo: bookingRepo,
	}
}

func (controller *BookingController) Create(w http.ResponseWriter, r *http.Request) {
	render.New().JSON(w, http.StatusOK, "Booking Create")
}
