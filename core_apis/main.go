package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ENVIRONMENT::", os.Getpagesize())
}
