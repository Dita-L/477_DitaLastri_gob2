package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//anonymous struct tanpa pengisian field
	var employee = struct {
		person   person
		division string
	}{}
	employee.person = person{name: "Airell", age: 23}
	employee.division = "Curriculum developer"

	//anonymous struct dengan pengisian field
	var employee2 = struct {
		person   person
		division string
	}{
		person:   person{name: "Ananda", age: 23},
		division: "Finance",
	}

	fmt.Printf("Employee1: %+v\n", employee)
	fmt.Printf("Employee2: %+v\n", employee2)
}
