package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")

	c := make(chan string)
	go count("Rahmat", c)

	for mes := range c {
		fmt.Println(mes)
	}

}

func count(s string, c chan string) {
	for i := 0; i < 6; i++ {
		c <- s
		time.Sleep(500 * time.Millisecond)
	}

	close(c)
}
