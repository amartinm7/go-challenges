package fetchad

import (
	"challenges/challenge_1/internal/infrastructure/ad/repository/mock"
	"testing"
)

func TestFetchAd_Execute(t *testing.T) {
	mockAdRepository := new(mock.MockAdRepository)
	mockAdRepository.On("FindById", mock.MotorAd.Id).Return(&mock.MotorAd, nil)
	NewFetchAdService(mockAdRepository).Execute(mock.MotorAd.Id)
	mockAdRepository.AssertCalled(t, "FindById", mock.MotorAd.Id)
}
