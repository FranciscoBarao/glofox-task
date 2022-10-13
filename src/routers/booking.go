package routers

import (
	"glofox-task/controllers"

	"github.com/go-chi/chi/v5"
)

func AddBookingRouter(router chi.Router, bookingController *controllers.BookingController) {
	router.Post("/api/booking", bookingController.Create)
}
