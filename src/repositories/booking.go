package repositories

import (
	"glofox-task/database"
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

	return repo.db.Create(booking)
}

func (repo *BookingRepository) ReadAll() ([]models.Booking, error) {

	var bookings []models.Booking
	if err := repo.db.ReadAll(&bookings, "", ""); err != nil {
		return bookings, err
	}

	return bookings, nil
}
