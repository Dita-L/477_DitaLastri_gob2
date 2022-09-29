package main

import (
	"fmt"
	"time"
)

func main() {
	/*//membuat unbuffered? channel dengan memberikan kapasitas buffer
	c:= make(chan int, 3)
	//3 adalah jumlah kapasitas buffer
	*/

	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main go routine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts recieving data")
	d := <-c1
	fmt.Println("main goroutine recieved data:", d)

	close(c1)
	time.Sleep(time.Second)

}
