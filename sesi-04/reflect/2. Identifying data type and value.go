package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	//fungsi reflect.ValueOf mengembalikan objek dalam tipe reflect.Value
	//reflect.Value berisikan informasi mengenai variabel yang bersangkutan
	//objek reflect.Value memiliki beberapa method, salah satunya Type()
	//method Type() mengembalikan tipe data dari variabel yang bersangkutan dalam bentuk string.

	fmt.Println("tipe variabel :", reflectValue.Type())

	//
	//Untuk menampilkan nilai variabel reflect, harus dipastikan dulu tipe datanya.
	//Ketika tipe data adalah int, maka bisa menggunakan method Int()
	//reflectValue.Int()

	//fmt.Println(reflectValue.Kind()) == int
	//fmt.Println(reflectValue.Int()) == 23

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel : ", reflectValue.Int())
	}

}
