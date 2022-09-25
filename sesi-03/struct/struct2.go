package main

import "fmt"

type employee struct {
	name     string
	age      int
	division string
}

func main() {
	var pekerja employee

	pekerja.name = "Airell"
	pekerja.age = 23
	pekerja.division = "Curriculum Developer"

	var pekerja2 = employee{name: "Ananda", age: 23, division: "Finance"}
	// nilai langsung diberikan

	fmt.Printf("pekerja1: %+v\n", pekerja)
	fmt.Printf("pekerja2: %+v\n", pekerja2)
	// %+v  untuk memformat sebuah struct menjadi string.

}
