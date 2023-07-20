package createad

import (
	"challenges/challenge_1/internal/infrastructure/ad/repository/mock"
	"testing"
)

func TestPostAd_Execute(t *testing.T) {

	mockAdRepository := new(mock.MockAdRepository)
	// mock := mock.NewMockAdRepository()
	mockAdRepository.On("Save", mock.MotorAd).Return(mock.MotorAd, nil)

	NewPostAdService(mockAdRepository).Execute(mock.MotorAd)

	mockAdRepository.AssertCalled(t, "Save", mock.MotorAd)
}
