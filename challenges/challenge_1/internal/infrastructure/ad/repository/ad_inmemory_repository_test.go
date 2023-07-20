package repository

import (
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/ad/repository/mock"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdRepository_SaveAd(t *testing.T) {
	var motorAd = mock.MotorAd
	savedAd, error := NewAdInMemoryRepository().Save(motorAd)
	assert.Nil(t, error)
	assert.Equal(t, savedAd, motorAd)
}

func TestAdRepository_FindById_success(t *testing.T) {
	var motorAd = mock.MotorAd
	adRepository := NewAdInMemoryRepository()
	savedAd, error := adRepository.Save(motorAd)
	assert.Nil(t, error)
	assert.Equal(t, savedAd, motorAd)
	foundAd, error := adRepository.FindById(motorAd.Id)
	assert.Nil(t, error)
	assert.Equal(t, *foundAd, motorAd)
}

func TestAdRepository_FindById_error(t *testing.T) {
	var motorAd = mock.MotorAd
	adRepository := NewAdInMemoryRepository()
	savedAd, _ := adRepository.Save(motorAd)
	assert.Equal(t, savedAd, motorAd)
	_, error := adRepository.FindById(uuid.New())
	assert.Error(t, error)
	assert.Equal(t, error, errors.New("element not found"))
}

func TestAdRepository_FindAll_success(t *testing.T) {
	var motorAd = mock.MotorAd
	var motorAds = make([]ad.Ad, 1)
	motorAds[0] = motorAd
	adRepository := NewAdInMemoryRepository()
	savedAd, error := adRepository.Save(motorAd)
	assert.Nil(t, error)
	assert.Equal(t, savedAd, motorAd)
	foundAds, error := adRepository.FindAll()
	assert.Nil(t, error)
	assert.Equal(t, *foundAds, motorAds)
}
