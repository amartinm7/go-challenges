package example1

type GenericFunc func() (interface{}, error)

func GetString() (interface{}, error) {
	return "Hello, world!", nil
}

func ExecuteAndReturn[T any](fn GenericFunc) (T, error) {
	result, err := fn()
	return result.(T), err
}
