package main

import "fmt"

func greet(name string, age int8) {
	fmt.Printf("My name is %s and i'm %d years old", name, age)
}

func main() {
	greet("Dita", 24)

}
