package theinterface

import "fmt"

type Animal interface {
	Speaks() string
	hasLegs() int
	color() string
}

func ShowAnimals(a Animal) {
	fmt.Println("Animal speaks ", a.Speaks(), "and has ", a.hasLegs(), "and color of ", a.color())
}

type Dog struct {
	Name string
}

func (d *Dog) Speaks() string {
	return "Bark bark"
}

func (d *Dog) hasLegs() int {
	return 4
}

func (d *Dog) color() string {
	return "brown"
}
