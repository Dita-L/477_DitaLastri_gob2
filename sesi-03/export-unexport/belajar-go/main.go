package main

import "belajar-go/helpers"

func main() {
	//helpers.greet() tidak bisa digunakan karena "greet" g nya kecil.
	helpers.Greet()

	var person = helpers.Person{}

	person.InvokeGreet()
}
