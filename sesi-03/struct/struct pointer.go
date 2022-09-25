package main

import (
	"fmt"
	"strings"
)

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

	var pekerja2 *employee = &pekerja //pekerja2 adalah sebuah pointer to struct

	fmt.Println("pekerja1 name:", pekerja.name)
	fmt.Println("pekerja2 name:", pekerja2.name)

	fmt.Println(strings.Repeat("#", 15))

	pekerja2.name = "Ananda"

	fmt.Println("pekerja1 name:", pekerja.name)
	fmt.Println("pekerja2 name:", pekerja2.name)

}
