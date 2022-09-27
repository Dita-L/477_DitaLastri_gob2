package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	//mengakses nilai menggunakan interface method
	fmt.Println("tipe variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface())
	//method interface() mengembalikan/return interface kosong == interface{}

	//mengakses nilai asli
	var nilai = reflectValue.Interface().(int)
	fmt.Println(nilai)
}
