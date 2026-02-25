package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() string {
	return fmt.Sprintf("EmployeeID: %s, Name: %s, Age: %d", e.EmployeeID, e.Name, e.Age)
}

func main() {
	employee := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E-1001",
	}

	fmt.Println(employee.PrintInfo())
}
