package repositories

import (
	"fmt"
	"log"
	"net/http"

	"glofox-task/database"
	"glofox-task/middleware"
	"glofox-task/models"
)

type BookingRepository struct {
	db *database.PostgresqlRepository
}

func NewBookingRepository(instance *database.PostgresqlRepository) *BookingRepository {
	return &BookingRepository{
		db: instance,
	}
}

func (repo *BookingRepository) Create(booking *models.Booking) error {

	// The following 3 checks should be on a service layer since they are Business logic. Additionally, that would abstract the ClassRepository misuse we have here.

	// Check if class exists
	var class models.Class
	if err := repo.db.ReadByCondition(&class, "start_date_time <= ? AND end_date_time >= ?", booking.GetDate(), booking.GetDate()); err != nil {
		if err.Error() == "Record not found" { // So that the message is not a generic "Record not found"
			return middleware.NewCustomError(http.StatusNotFound, "Class Not Found - There is no Class on the provided date")
		}
		return err
	}

	// Check the number of bookings (Of a class in that date)
	var count int64
	if err := repo.db.Count(&models.Booking{}, &count, "class_id = ? AND date_time = ?", class.GetID(), booking.GetDate()); err != nil {
		return err
	}

	// Check if not overbooking
	if class.IsOverbooking(int(count) + 1) {
		log.Println("Error - Overbooking class: " + fmt.Sprintf("%v", class))
		return middleware.NewCustomError(http.StatusConflict, "Overbooking - Too many bookings for that class")
	}

	// Setting one-to-many relation
	booking.SetClassID(class.GetID())

	return repo.db.Create(booking)
}

func (repo *BookingRepository) ReadAll() ([]models.Booking, error) {

	var bookings []models.Booking
	if err := repo.db.ReadAll(&bookings); err != nil {
		return bookings, err
	}

	return bookings, nil
}
