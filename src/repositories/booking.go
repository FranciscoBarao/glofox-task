package repositories

type BookingRepository struct {
	db string
}

func NewBookingRepository(instance string) *BookingRepository {
	return &BookingRepository{
		db: instance,
	}
}

func (svc *BookingRepository) Create() error {
	return nil
}
