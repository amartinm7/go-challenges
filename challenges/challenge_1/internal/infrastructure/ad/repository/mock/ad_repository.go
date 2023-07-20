package mock

import (
	"challenges/challenge_1/internal/domain/ad"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

var MotorAd, _ = ad.NewAd(uuid.New().String(), "Opel Omega", "como nuevo.", 15000, "2023-03-01")

type MockAdRepository struct {
	mock.Mock
}

func NewMockAdRepository() *MockAdRepository { return &MockAdRepository{} }

func (m *MockAdRepository) Save(newAd ad.Ad) (ad.Ad, error) {
	args := m.Called(newAd)
	return args.Get(0).(ad.Ad), args.Error(1)
}

func (m *MockAdRepository) FindById(id uuid.UUID) (*ad.Ad, error) {
	args := m.Called(id)
	return args.Get(0).(*ad.Ad), args.Error(1)
}

func (m *MockAdRepository) FindAll() (*[]ad.Ad, error) {
	args := m.Called()
	return args.Get(0).(*[]ad.Ad), args.Error(1)
}
