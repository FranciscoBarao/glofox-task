package repositories

type ClassRepository struct {
	db string
}

func NewClassRepository(instance string) *ClassRepository {
	return &ClassRepository{
		db: instance,
	}
}

func (svc *ClassRepository) Create() (string, error) {
	return "CLASS", nil
}
