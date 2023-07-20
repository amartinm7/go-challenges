package listads

import (
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/ad/repository/mock"
	"testing"
)

func TestFetchAllAds_Execute(t *testing.T) {
	var motorAds = make([]ad.Ad, 10)
	motorAds[0] = mock.MotorAd
	mockAdRepository := new(mock.MockAdRepository)
	mockAdRepository.On("FindAll").Return(&motorAds, nil)
	NewFetchAllAds(mockAdRepository).Execute()
	mockAdRepository.AssertCalled(t, "FindAll")
}
