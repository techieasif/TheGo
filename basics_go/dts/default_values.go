package dts

import "fmt"

var name, gender, city string
var age int
var salary float64

func showDefaults() {
	fmt.Println("Defaults of string:: ", name, gender, city)
	fmt.Println("Defaults of int::", age)
	fmt.Println("Default of float64", salary)
}
