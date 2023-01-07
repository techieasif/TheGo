package theinterface

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name        string `json:"employeeName"`
	JoiningDate string `json:"joiningDate"`
	BadgeId     string `json:"badgeId"`
}

func ShowEmployeeDetails() {
	e := &Employee{
		Name:        "ajay",
		JoiningDate: "20/10/2022",
		BadgeId:     "b234",
	}
	marshelled, _ := json.Marshal(e)
	fmt.Println("Employee Details are: ", string(marshelled))
}
