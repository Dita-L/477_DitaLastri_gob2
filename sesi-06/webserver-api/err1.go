package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var err error

	number, err = strconv.Atoi("123GHIJ")
	// karena ada huruf, pasti error
	// error tersebut yang akan ditampung variabel "err"

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
		//menggunakan method Error untuk menampilkan pesan jika terjadi sebuah error.
		//Method Error berasal dari tipe data error.

		//Tapi pakai fmt.Println(err) outputnya sama aja. kenapa?
	}

}
