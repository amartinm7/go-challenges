package repository

import (
	"challenges/challenge_1/internal/domain/ad"
	"errors"
	"github.com/google/uuid"
)

type adInMemoryRepository struct {
	storage storage
}

type storage struct {
	store []ad.Ad
}

func NewAdInMemoryRepository() ad.AdRepository {
	return &adInMemoryRepository{}
}

func (adRepository *adInMemoryRepository) Save(ad ad.Ad) (ad.Ad, error) {
	adRepository.storage.store = append(adRepository.storage.store, ad)
	return ad, nil
}

func (adRepository *adInMemoryRepository) FindById(id uuid.UUID) (*ad.Ad, error) {
	for _, element := range adRepository.storage.store {
		if element.Id == id {
			return &element, nil
		}
	}
	return nil, errors.New("element not found")
}

func (adRepository *adInMemoryRepository) FindAll() (*[]ad.Ad, error) {
	return &adRepository.storage.store, nil
}
