package example1

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAd(t *testing.T) {

	t.Run("should return hello world", func(t *testing.T) {
		expectedAd := "Hello, world!"
		wg := sync.WaitGroup{}
		wg.Add(1)
		result, err := ExecuteAndReturn[string](&wg, GetString)
		wg.Wait()
		fmt.Println(result)
		assert.Nil(t, err)
		assert.Equal(t, expectedAd, result)
	})
}
