package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person   person //ada struct "person" didalam struct "employee", sebagai field
	division string
}

func main() {
	var pekerja1 = employee{}

	pekerja1.person.name = "Airell"
	pekerja1.person.age = 23
	pekerja1.division = "Curriculum Developer"

	fmt.Printf("%+v", pekerja1)
}
