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
	Create(model *models.Booking) error
	ReadAll() ([]models.Booking, error)
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
	controller.repo.Create(&booking)

	render.New().JSON(w, http.StatusOK, booking)
}

func (controller *BookingController) GetAll(w http.ResponseWriter, r *http.Request) {

	// Calls create on repository
	bookings, err := controller.repo.ReadAll()
	if err != nil {
		middleware.ErrorHandler(w, err)
		return
	}

	render.New().JSON(w, http.StatusOK, bookings)
}
