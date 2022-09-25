package main

import (
	"fmt"
	"strings"
)

func greet(pesan string, nama []string) string {
	//"nama" adalah slice
	//func greet menerima string = "pesan" dan slice = "nama"
	// return string

	var namanama = strings.Join(nama, ", ")
	//Fungsi Join dari package "strings".
	//Berguna untuk menggabungkan data string dari slice maupun array.
	// ("") digunakan sebagai pemisah antar string. kalau " " berarti spasi, kalau ", " berarti koma dan spasi.

	var result string = fmt.Sprintf("%s %s", pesan, namanama)
	//Function Sprintf memiliki fungsi yang sama seperti function Printf,
	//function Sprintf akan me-return sebuah nilai, function Printf tidak.

	return result
}

func main() {
	var names = []string{"aira", "jordan"}
	var printMsg = greet("Heii", names)

	fmt.Println(printMsg)
}
