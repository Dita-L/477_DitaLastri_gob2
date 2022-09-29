package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	//var pass []byte
	var user string

	fmt.Println("Input Username:")
	fmt.Scanln(&user)

	//fmt.Println(pass)

	if valid, err := validUser(user); err == nil {
		fmt.Println(valid)
		pass, _ := gopass.GetPasswdMasked()
		//fmt.Scanln(&pass)

		if val, ror := validPass(pass); ror == nil {
			fmt.Println(val)
		} else {
			fmt.Println(ror.Error())
		}

	} else {
		fmt.Println(err.Error())
	}

}

func validUser(u string) (string, error) {
	if u == "dita.lastri" {
		return "Input Password", nil
	} else {
		return "", errors.New("Incorrect Username")
	}
}

func validPass(p []byte) (string, error) {
	pasw := string(p[:])
	if pasw == "1234567" {
		return "Welcome!", nil
	} else {
		return "", errors.New("Incorrect Password")
	}
}
