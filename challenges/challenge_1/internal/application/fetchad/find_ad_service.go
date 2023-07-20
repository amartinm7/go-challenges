package fetchad

import (
	"challenges/challenge_1/internal/domain/ad"
	"github.com/google/uuid"
)

type fetchAd struct {
	repository ad.AdRepository
}

func NewFetchAdService(repository ad.AdRepository) FetchAdService {
	return &fetchAd{repository: repository}
}

func (service *fetchAd) Execute(id uuid.UUID) (*ad.Ad, error) {
	return service.repository.FindById(id)
}

type FetchAdService interface {
	Execute(id uuid.UUID) (*ad.Ad, error)
}
