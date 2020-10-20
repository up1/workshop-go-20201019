package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	position string
	person
}

// Implement from interface Stringer
func (e employee) String() string {
	return "XXXXXX"
}

// NewEmployee :: Way to create a new employee
func newEmployee() employee {
	return employee{}
}

func main() {
	e := employee{"Dev", person{"pui", 30}}
	fmt.Println(e)
	fmt.Println(e.name, e.age, e.position)
}
