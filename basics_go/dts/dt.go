package dts

import "fmt"

func RunDts() {
	enteredNumber := 8
	fmt.Printf("Is %v Prime ->  %v\n", enteredNumber, findMePrime(enteredNumber))
	printTableForOf(121)
	showTheUsageOfEnums()
}

// Simple function with return type bool
func findMePrime(num int) bool {
	return num%2 == 1
}

// Simple for-loop.
func printTableForOf(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %d = %v\n", n, i, n*i)
	}
}

// Color Defining new cutom type
type Color float64

// Defining Color based enums
const (
	Red Color = iota
	Black
	Violet
)

func showTheUsageOfEnums() {
	fmt.Println(Red.String())
}

// Function on Color
func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Black:
		return "Black"
	case Violet:
		return "Violet"
	default:
		return "NO_COLOR"
	}
}
