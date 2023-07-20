package repository

import (
	"challenges/challenge_1/internal/domain/ad"
	mock2 "challenges/challenge_1/internal/infrastructure/ad/repository/mock"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdRepository_Save(t *testing.T) {
	// Inicializa el mock de sql.DB
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error al crear el mock de sql.DB: %s", err)
	}
	defer db.Close()

	t.Run("should save with success", func(t *testing.T) {
		var motorAd = mock2.MotorAd
		sqlMock.ExpectExec(
			"INSERT INTO ADS (id, title, description, price, time_stamp) VALUES ($1, $2, $3, $4, $5)").
			WithArgs(motorAd.Id.String(), motorAd.Title, motorAd.Description.Value, motorAd.Price, motorAd.Timestamp.Format("2006-01-02")).
			WillReturnResult(sqlmock.NewResult(1, 1))

		savedAd, error := NewAdRepository(db).Save(motorAd)
		assert.Nil(t, error)
		assert.Equal(t, savedAd, motorAd)

		if err := sqlMock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectativas no cumplidas: %s", err)
		}
	})

	t.Run("should save with error", func(t *testing.T) {
		var motorAd = ad.Ad{}
		sqlMock.ExpectExec(
			"INSERT INTO ADS (id, title, description, price, time_stamp) VALUES ($1, $2, $3, $4, $5)").
			WithArgs(motorAd.Id.String(), motorAd.Title, motorAd.Description.Value, motorAd.Price, motorAd.Timestamp.Format("2006-01-02")).
			WillReturnError(errors.New("error trying to persist an AD on database"))

		_, error := NewAdRepository(db).Save(motorAd)
		assert.Error(t, error)

		if err := sqlMock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectativas no cumplidas: %s", err)
		}
	})
}

func TestAdRepository_FindById(t *testing.T) {
	// Inicializa el mock de sql.DB
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error al crear el mock de sql.DB: %s", err)
	}
	defer db.Close()

	t.Run("should findById with success", func(t *testing.T) {
		var motorAd = mock2.MotorAd
		sqlMock.ExpectExec(
			"SELECT * FROM ADS WHERE ID = $1").
			WithArgs(motorAd.Id.String()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		foundAd, error := NewAdRepository(db).FindById(motorAd.Id)
		assert.Nil(t, error)
		assert.Equal(t, foundAd, motorAd)

		if err := sqlMock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectativas no cumplidas: %s", err)
		}
	})

}

func TestAdRepository_FindAll(t *testing.T) {
	// Inicializa el mock de sql.DB
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error al crear el mock de sql.DB: %s", err)
	}
	defer db.Close()

	t.Run("should findAll with success", func(t *testing.T) {
		var motorAds = make([]ad.Ad, 1)
		motorAds[0] = mock2.MotorAd
		sqlMock.ExpectExec(
			"SELECT * FROM ADS").
			WillReturnResult(sqlmock.NewResult(1, 1))

		foundAds, error := NewAdRepository(db).FindAll()
		assert.Nil(t, error)
		assert.Equal(t, foundAds, motorAds)

		if err := sqlMock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectativas no cumplidas: %s", err)
		}
	})
}
