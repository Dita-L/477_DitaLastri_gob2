package main

import "fmt"

type employee struct {
	name     string
	age      int
	division string
}

func main() {
	var pekerja employee // membuat variabel kosong bernama "pekerja" dengan type "employee"

	pekerja.name = "Airell" //nilai-nilai baru diberikan
	pekerja.age = 23
	pekerja.division = "Curriculum Developer"

	fmt.Println(pekerja)
	fmt.Println(pekerja.name)
}
