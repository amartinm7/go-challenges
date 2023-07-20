package main

import (
	"fmt"

	"github.com/antonio-martin/learning-go/challenges/pure_functions/internal/example1"
)

func main() {
	result, _ := example1.ExecuteAndReturn[string](example1.GetString)
	fmt.Println(result)
}
