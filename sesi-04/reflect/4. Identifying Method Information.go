package main

import (
	"fmt"
	"reflect"
)

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "john wick", Grade: 2} //instance struct (struct instant? wkwk)
	fmt.Println("nama:", s1.Name)

	var reflectValue = reflect.ValueOf(s1) //refleksi value s1

	var method = reflectValue.MethodByName("SetName") //refleksi method untuk mengubah nama
	method.Call([]reflect.Value{                      //Call(), untuk mengeksekusi
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s1.Name)
}
