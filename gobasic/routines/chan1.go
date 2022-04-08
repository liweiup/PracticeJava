package main

import "fmt"

func maint() {
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	fmt.Println(<-ch)
	ch <- "D"
	fmt.Println(<-ch)
}
