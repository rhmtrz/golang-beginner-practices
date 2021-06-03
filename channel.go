package main

import (
	"fmt"
	"time"
)

func sendValue(c chan int) {
	fmt.Println("Executing goroutine")
	time.Sleep(1 * time.Second)
	c <- 120
	println("Finished executing goroutine")

}

func main() {

	value := make(chan int, 2)
	defer close(value)
	go sendValue(value)
	go sendValue(value)

	val := <-value
	fmt.Println(val)
	time.Sleep(1 * time.Second)
}
