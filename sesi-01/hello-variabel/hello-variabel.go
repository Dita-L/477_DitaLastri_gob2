package main

import "fmt"

func main() {
	//VARIABEL DENGAN DATA TYPE
	var name string = "Dita"
	var age int = 24
	fmt.Println("nama saya ", name)
	fmt.Println("usia saya ", age)

	//VARIABEL TANPA DATA TYPE
	var Nama = "dita"
	var Usia = 24
	fmt.Printf("%T, %T\n", Nama, Usia)

	//SHORT DECLARATION
	kotaAsal := "Bekasi"
	fmt.Printf("kota asal saya %s\n", kotaAsal)

	// MULTIPLE VARIABLE
	var nama, usia, TTL = "Dita", 24, 271097
	fmt.Println(nama, usia, TTL)

	// fmt.Printf USAGE
	fmt.Printf("nama saya %s, saya lahir di %s. Sekarang saya berusia %d\n", nama, kotaAsal, usia)

	//UNDERSCORE VARIABLE
	var testing string = "ini hanya contoh"
	var tes1, tes2, tes3 = 1, 2, 3
	_, _, _, _ = testing, tes1, tes2, tes3

}
