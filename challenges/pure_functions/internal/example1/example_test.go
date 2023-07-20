package example1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAd(t *testing.T) {

	t.Run("should return hello world", func(t *testing.T) {
		expectedAd := "Hello, world!"
		result, err := ExecuteAndReturn[string](GetString)
		fmt.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, expectedAd, result)
	})

	//t.Run("Validate description is less than 50 chars ", func(t *testing.T) {
	//	_, err := NewAd(uuid.New().String(), "Opel Omega", "como nuevo. Comprueba que todo está bien: daños y accidentes, lecturas de kilometraje, inspecciones, etc.", 15000, time.Now().String())
	//	assert.NotNil(t, err)
	//	assert.Error(t, err, errors.New("Description can be more than 50 chars"))
	//})
}
