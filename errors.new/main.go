package main

import (
	"errors"
	"fmt"
)

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func printResult(result float64, err error) (float64, error) {
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	fmt.Println("Result:", result)
	return result, nil
}

func main() {
	printResult(add(1, 2))
	printResult(subtract(1, 2))
	printResult(multiply(1, 2))
	printResult(divide(1, 2))
}
