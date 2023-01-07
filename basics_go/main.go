package main

import (
	"techieasif.com/basics_go/theLoop"

	theinterface "techieasif.com/basics_go/theInterface"
)

func main() {
	//dts.RunDts()
	theLoop.RunTheLoop()
	theinterface.ShowEmployeeDetails()
	//names := lo.Uniq([]string{"Samuel", "John", "Samuel", "john"})
	//fmt.Println(names)
}
