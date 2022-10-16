package repositories

import (
	"glofox-task/database"
	"glofox-task/models"
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

	// Check if dates are available

	return repo.db.Create(class)
}

func (repo *ClassRepository) ReadAll() ([]models.Class, error) {

	var classes []models.Class
	if err := repo.db.ReadAll(&classes); err != nil {
		return classes, err
	}

	return classes, nil
}
