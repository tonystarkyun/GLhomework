package main

import "testing"

func TestEmployeePrintInfo(t *testing.T) {
	employee := Employee{
		Person:     Person{Name: "Bob", Age: 28},
		EmployeeID: "EMP-002",
	}

	want := "EmployeeID: EMP-002, Name: Bob, Age: 28"
	if got := employee.PrintInfo(); got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestEmployeeComposition(t *testing.T) {
	employee := Employee{
		Person:     Person{Name: "Carol", Age: 35},
		EmployeeID: "EMP-003",
	}

	if employee.Name != "Carol" {
		t.Fatalf("expected Name Carol, got %s", employee.Name)
	}
	if employee.Age != 35 {
		t.Fatalf("expected Age 35, got %d", employee.Age)
	}
}
