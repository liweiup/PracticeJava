package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 1

func mainx() {
	// msg := "hello"
	// wg.Add(2)
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// go func() {
	// 	fmt.Println("11")
	// 	wg.Done()
	// }()
	// wg.Wait()
	// runtime.GOMAXPROCS(1)
	// fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	m.RLock()
	// 	go sayHello()
	// 	m.Lock()
	// 	go counterIncrement()
	// 	time.Sleep(3 * time.Second)
	// }
	// wg.Wait()
}

var ch = make(chan int, 5)

func sayHello() {
	// m.RUnlock()
	cint := <-ch
	// cint := counter
	fmt.Printf("hello %v \n", cint)
	wg.Done()
}
func counterIncrement(i int) {
	// m.Unlock()
	ch <- counter
	counter++
	wg.Done()
}

func main() {
	// runtime.GOMAXPROCS(16)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	for i := 1; i < 100; i++ {
		wg.Add(2)
		// m.RLock()
		go sayHello()
		// m.Lock()
		go counterIncrement(i)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
