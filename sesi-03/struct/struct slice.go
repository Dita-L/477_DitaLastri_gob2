package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var people = []person{
		{name: "Airell", age: 23},
		{name: "Ananda", age: 23},
		{name: "mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

}
