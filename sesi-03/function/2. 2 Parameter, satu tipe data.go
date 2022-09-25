package main

import "fmt"

func greet(name, address string) {
	fmt.Println("Hello! My name is", name)
	fmt.Println("I live in", address)
}

func main() {
	greet("Dita", "Bekasi")
}
