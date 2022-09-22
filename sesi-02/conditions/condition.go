package main

import "fmt"

func main() {

	//IF
	var lampu = "hijau"

	if lampu == "hijau" {
		fmt.Println("Silahkan jalan")
	} else if lampu == "merah" {
		fmt.Println("Berhenti!")
	} else {
		fmt.Println("Hati-hati")
	}

	//SWITCH
	var nilai = 8

	switch nilai {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//SWITCH DENGAN OPERATOR PERBANDINGAN, DAN "FALLTHROUGH"
	var nilai2 = 5

	switch {
	case nilai2 == 8:
		fmt.Println("perfect")
	case (nilai2 > 3) && (nilai2 < 8):
		fmt.Println("awesome")
		fallthrough
	case nilai2 <= 5:
		fmt.Println("try harder")
		fallthrough
	case nilai2 == 5:
		fmt.Println("wow")
	default:
		fmt.Println("not bad")
	}

	//KONDISIONAL BERSARANG. IF + SWITCH
	var score = 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it!")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
