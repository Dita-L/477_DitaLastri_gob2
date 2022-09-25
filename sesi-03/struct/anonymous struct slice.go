package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	var employee = []struct {
		person   person
		division string
	}{
		{person: person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: person{name: "Ananda", age: 23}, division: "Finance"},
		{person: person{name: "mailo", age: 23}, division: "Marketing"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}

}
