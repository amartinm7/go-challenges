package ad

import "github.com/google/uuid"

type AdRepository interface {
	Save(ad Ad) (Ad, error)
	FindById(id uuid.UUID) (*Ad, error)
	FindAll() (*[]Ad, error)
}
