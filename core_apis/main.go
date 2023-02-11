package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("Pause for a moment %T")
	}
	f()
}

func advFn(arg int) {
	fmt.Println("Inside advance function", arg)
}
