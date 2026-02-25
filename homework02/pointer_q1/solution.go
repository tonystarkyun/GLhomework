package main

import "fmt"

// AddTen adds 10 to the value pointed to by n.
func AddTen(n *int) {
	if n == nil {
		return
	}
	*n += 10
}

func Solve(input int) int {
	value := input
	AddTen(&value)
	return value
}

func main() {
	value := 5
	AddTen(&value)
	fmt.Println(value)
}
