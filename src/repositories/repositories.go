package repositories

import "glofox-task/database"

// Repositories contains all the repositories structs
type Repositories struct {
	BookingRepository *BookingRepository
	ClassRepository   *ClassRepository
}

// InitRepositories should be called in main.go
func InitRepositories(db *database.PostgresqlRepository) *Repositories {
	bookingRepository := NewBookingRepository(db)
	classRepository := NewClassRepository(db)

	return &Repositories{
		BookingRepository: bookingRepository,
		ClassRepository:   classRepository,
	}
}
