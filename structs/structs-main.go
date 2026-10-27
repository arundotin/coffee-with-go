package main

import (
	"coffee-with-go/common"
	"fmt"
)

func sts1() {

	employee := common.Employee{
		ID:     101,
		Name:   "Arun",
		Salary: 1200000,
	}

	println(employee.ID)
	println(employee.Name)
	fmt.Printf("Salary: %.2f\n", employee.Salary)
}

func main() {
	sts1()
}
