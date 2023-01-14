package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100.0, 78.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The result %v", result)
}

func divide(val1, val2 float32) (float32, error) {
	var result float32
	if val2 == 0 {
		return result, errors.New("error: cannot divide by zero")
	}

	result = val1 / val2
	return result, nil
}
