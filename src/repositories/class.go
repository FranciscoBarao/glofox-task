package repositories

import (
	"fmt"
	"glofox-task/database"
	"glofox-task/middleware"
	"glofox-task/models"
	"log"
	"net/http"
)

type ClassRepository struct {
	db *database.PostgresqlRepository
}

func NewClassRepository(instance *database.PostgresqlRepository) *ClassRepository {
	return &ClassRepository{
		db: instance,
	}
}

func (repo *ClassRepository) Create(class *models.Class) error {

	// Check if dates are available by checking if any date overlaps with new dates
	var exists models.Class
	repo.db.ReadByCondition(&exists, "start_date_time BETWEEN ? AND ? OR end_date_time BETWEEN ? AND ?", class.GetStartDate(), class.GetEndDate(), class.GetStartDate(), class.GetEndDate())

	if exists.GetID() != 0 { // No date exists
		log.Println("Error - overlapping classes: " + fmt.Sprintf("%v", class) + " and " + fmt.Sprintf("%v", exists))
		return middleware.NewCustomError(http.StatusBadRequest, "Overlapping - New class overlapps with another already existing class")
	}

	return repo.db.Create(class)
}

func (repo *ClassRepository) ReadAll() ([]models.Class, error) {

	var classes []models.Class
	if err := repo.db.ReadAll(&classes); err != nil {
		return classes, err
	}

	return classes, nil
}
