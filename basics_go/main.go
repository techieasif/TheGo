package main

import (
	"fmt"

	"github.com/samber/lo"
	theinterface "techieasif.com/basics_go/theInterface"
)

func main() {
	//dts.RunDts()
	// theLoop.RunTheLoop()
	theinterface.ShoeEmployeeDetails()
	names := lo.Uniq([]string{"Samuel", "John", "Samuel", "john"})
	fmt.Println(names)
}
