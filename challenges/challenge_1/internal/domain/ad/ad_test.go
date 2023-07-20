package ad

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewAd(t *testing.T) {

	t.Run("Create new Ad", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2022-12-01")
		expectedAd := Ad{
			Id:          uuid.New(),
			Title:       "Opel Omega",
			Description: Description{"como nuevo."},
			Price:       15000,
			Timestamp:   date,
		}
		ad, err := NewAd(expectedAd.Id.String(), "Opel Omega", "como nuevo", 15000, "2022-12-01")
		assert.Nil(t, err)
		assert.Equal(t, ad, ad)
	})

	t.Run("Validate description is less than 50 chars ", func(t *testing.T) {
		_, err := NewAd(uuid.New().String(), "Opel Omega", "como nuevo. Comprueba que todo está bien: daños y accidentes, lecturas de kilometraje, inspecciones, etc.", 15000, time.Now().String())
		assert.NotNil(t, err)
		assert.Error(t, err, errors.New("Description can be more than 50 chars"))
	})
}
