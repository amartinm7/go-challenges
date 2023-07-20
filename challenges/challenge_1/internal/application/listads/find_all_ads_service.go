package listads

import "challenges/challenge_1/internal/domain/ad"

type fetchAllAds struct {
	repository ad.AdRepository
}

func NewFetchAllAds(repository ad.AdRepository) FetchAllAdsService {
	return &fetchAllAds{repository: repository}
}

func (service *fetchAllAds) Execute() (*[]ad.Ad, error) {
	return service.repository.FindAll()
}

type FetchAllAdsService interface {
	Execute() (*[]ad.Ad, error)
}
