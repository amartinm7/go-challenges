package example1

import (
	"fmt"
	"sync"
)

type GenericFunc func() (interface{}, error)

func GetString() (interface{}, error) {
	return "Hello, world!", nil
}

func ExecuteAndReturn[T any](wg *sync.WaitGroup, fn GenericFunc) (T, error) {
	var result T
	var err error
	go func() {
		defer wg.Done()
		res, e := fn()
		if e != nil {
			fmt.Errorf("oups")
		}
		result = res.(T)
		err = e
	}()
	return result, err
}
