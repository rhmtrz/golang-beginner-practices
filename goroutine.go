package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func loopDigit(dig int) {
	for i := 0; i < dig; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	wg.Done()
}

var counter = 0

func main() {

	fmt.Println("Concurrency with goroutines")
	runtime.GOMAXPROCS(10)
	wg.Add(2)
	go loopDigit(5)

	mes := "Hello world"
	go func(m string) {
		fmt.Println(m)
		wg.Done()
	}(mes)
	mes = "hello Rahmat"

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	fmt.Printf("Threads 2: %v\n", runtime.GOMAXPROCS(-1))
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hiii #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
