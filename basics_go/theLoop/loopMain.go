package theLoop

import (
	"fmt"
	theinterface "techieasif.com/basics_go/theInterface"
)

func RunTheLoop() {
	simple4Loop(2345678)
	shortHand4Loop(2345678)
	whileLoop(2345678)
	dog := theinterface.Dog{
		Name: "Joru",
	}
	theinterface.ShowAnimals(&dog)
}

func simple4Loop(number int) {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	fmt.Println("Sum --> ", sum)
}

func shortHand4Loop(number int) {
	sum := 1
	///for ; sum < number ;
	for sum < number {
		sum += sum
	}
	fmt.Println("THE short sum :-> ", sum)
}

func whileLoop(number int) {
	sum := 1
	for sum < number {
		sum += sum
	}
	fmt.Println("THE while sum :-> ", sum)
}
