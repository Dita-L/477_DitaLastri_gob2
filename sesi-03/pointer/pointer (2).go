package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstperson string = "John"
	var secondperson *string = &firstperson

	fmt.Println("firstperson (value) :", firstperson)
	fmt.Println("firstperson (memory address) :", &firstperson)
	fmt.Println("secondperson (value) :", *secondperson)
	fmt.Println("secondperson (memory address) :", secondperson)

	fmt.Println("\n", strings.Repeat("#", 15), "\n")

	*secondperson = "Doe"

	fmt.Println("firstperson (value) :", firstperson)
	fmt.Println("firstperson (memory address) :", &firstperson)
	fmt.Println("secondperson (value) :", *secondperson)
	fmt.Println("secondperson (memory address) :", secondperson)

}
