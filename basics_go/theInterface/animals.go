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
	//TODO implement me
	panic("implement me")
}

func (d *Dog) hasLegs() int {
	//TODO implement me
	panic("implement me")
}

func (d *Dog) color() string {
	//TODO implement me
	panic("implement me")
}
