package theinterface

import "fmt"

type Employee struct {
		 Name string `json:"employeeName"`
		 JoiningDate string `json:"joiningDate"`
		 BadgeId string `json:"badgeId"`
}


func ShoeEmployeeDetails(){
	e := Employee{
		Name: "ajay",
		JoiningDate: "20/10/2022",
		BadgeId: "b234",
	}

	fmt.Println("Employee Details are: ", e)
}