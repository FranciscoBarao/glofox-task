package repositories

import "glofox-task/database"

type ClassRepository struct {
	db *database.PostgresqlRepository
}

func NewClassRepository(instance *database.PostgresqlRepository) *ClassRepository {
	return &ClassRepository{
		db: instance,
	}
}

func (repo *ClassRepository) Create() (string, error) {
	return "CLASS", nil
}
