package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic("cant handle error")
		}
	}()
	panic("Something bad happend")
	fmt.Println("done panicker")
}

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("stop")
}
