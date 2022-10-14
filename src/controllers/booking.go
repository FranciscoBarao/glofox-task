package controllers

import (
	"net/http"

	"github.com/unrolled/render"

	"glofox-task/middleware"
	"glofox-task/models"
	"glofox-task/repositories"
	"glofox-task/utils"
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

	// Body to Booking Struct
	var booking models.Booking
	if err := utils.DecodeJSONBody(w, r, &booking); err != nil {
		middleware.ErrorHandler(w, err)
		return
	}

	// Validate Booking input
	if err := utils.ValidateStruct(&booking); err != nil {
		middleware.ErrorHandler(w, err)
		return
	}

	// Calls create on repository
	controller.repo.Create()

	render.New().JSON(w, http.StatusOK, booking)
}
