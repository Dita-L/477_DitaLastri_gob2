package main

import "fmt"

func main() {

	//membuat slice
	var fruit = []string{"apel", "pisang", "jambu"}
	fmt.Println(fruit)

	//membuat slice menggunakan "make"
	var angka = make([]int, 3)
	fmt.Println(angka)
	fmt.Printf("%T\n", angka)
	fmt.Printf("%#v\n", angka)

	//mengakses elemen pada slice
	angka[1] = 7
	fmt.Println(angka)

	//menambahkan elemen dengan "append"
	fruit = append(fruit, "mangga", "rambutan")
	fmt.Println(fruit)

	//menggabungkan 2 atau lebih slice dengan "append" dan elipsis (...)
	var fruit2 = []string{"pir", "anggur"}

	fruit = append(fruit, fruit2...)
	fmt.Println(fruit)

	//menyalin slice
	abc := copy(fruit2, fruit)

	fmt.Println(abc)
	fmt.Println(fruit)
	fmt.Println(fruit2)

	//slicing
	var buah = []string{"apel", "pisang", "jambu", "mangga"}
	var x = buah[1:2]
	fmt.Println(buah)
	fmt.Println(x)

	//Backing Array
	x[0] = "rambutan"
	fmt.Println(buah)
	fmt.Println(x)

}
