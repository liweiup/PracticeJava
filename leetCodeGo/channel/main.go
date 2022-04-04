package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	ch := make(chan int)
	// wg.Add(2)
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// go func() {
	// 	ch <- 1
	// 	wg.Done()
	// }()
	// wg.Wait()

	ch = make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func(j int) {
			// this line will pause the execution of this goroutine until there is space in the channel, only one msg can be in the channel at one time
			ch <- j
			wg.Done()
		}(j)
	}
	wg.Wait()
}
