package main

import "fmt"

func main() {

  //Operator Aritmatika
	  var a = 6
  	var b = 7
  	fmt.Println(a + b)
  	fmt.Println(b % a)

  //OPERATOR LOGIKA
  	val1 := true && true
  	val2 := true && false
  	val3 := val1 || true
  	val4 := val1 || false
  	val5 := true && false || true
  	val6 := !(val5)

  	fmt.Println(val1, val2, val3, val4, val5, val6)

  //OPERATOR RELASIONAL

	  var x int = 31
	  var y int = 23

	  fmt.Println(y < x)
	  fmt.Println(!(y < x))

}
