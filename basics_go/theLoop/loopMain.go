package theLoop

import "fmt"

func RunTheLoop() {
	simple4Loop(2345678)
}

func simple4Loop(number int) {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	fmt.Println("Sum --> ", sum)
}
