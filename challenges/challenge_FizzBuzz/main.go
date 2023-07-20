package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the fizzbuzz program")
	printInstructions()
	for i := 0; i < 100; i++ {
		fmt.Println(NewFizzBuzz(i).Execute())
	}
	fmt.Println("Ending fizzbuzz program")
}

func printInstructions() {
	fmt.Println("Loop over the 0..99 numbers and...")
	fmt.Println(" When the number is module 5 && 3 is 0, writes down FizzBuzz")
	fmt.Println(" When the number is module 5 is 0, writes down Fizz")
	fmt.Println(" When the number is module 3 is 0, writes down Buzz")
	fmt.Println(" else writes down the number itself")
}

type FizzBuzz interface {
	Execute() string
}

type fizzbuzz struct {
	value string
	index int
}

func NewFizzBuzz(index int) FizzBuzz {
	return fizzbuzz{value: "", index: index}
}

func (f fizzbuzz) Execute() string {
	return f.doFizzBuzzOrElse().
		doFizzOrElse().
		doBuzzOrElse().
		doNoFizzBuzzOrElse().
		value
}

func (f fizzbuzz) doFizzBuzzOrElse() fizzbuzz {
	if f.value == "" && f.index%3 == 0 && f.index%5 == 0 {
		f.value = "FizzBuzz"
	}
	return f
}

func (f fizzbuzz) doFizzOrElse() fizzbuzz {
	if f.value == "" && f.index%3 == 0 {
		f.value = "Fizz"
	}
	return f
}

func (f fizzbuzz) doBuzzOrElse() fizzbuzz {
	if f.value == "" && f.index%5 == 0 {
		f.value = "Buzz"
	}
	return f
}

func (f fizzbuzz) doNoFizzBuzzOrElse() fizzbuzz {
	if f.value == "" && f.index%3 != 0 && f.index%5 != 0 {
		f.value = strconv.Itoa(f.index)
	}
	return f
}
