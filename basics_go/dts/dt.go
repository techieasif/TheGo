package dts

import "fmt"

func RunDts() {
	enteredNumber := 8
	fmt.Printf("Is %v Prime ->  %v\n", enteredNumber, findMePrime(enteredNumber))
	printTableForOf(121)
	showTheUsageOfEnums()
	fmt.Println("Random Number - ", generateRandomNumber())
	fmt.Println("Value of Pi - ", showValueOfPi())
	a, b := swapValues("Kumar", "ajay")
	fmt.Println("Swapped values: ", a, b)
	showDefaults()
	showInitializedValues()
	showTypeConversion()
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
	fmt.Println(Violet.EnumIndex())
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

func (c Color) EnumIndex() int {
	return int(c)

}

func showTypeConversion() {
	var age = 24
	var ageFloat = float64(age)
	var ageUInt = uint(ageFloat)
	fmt.Println(age, ageFloat, ageUInt)
}
