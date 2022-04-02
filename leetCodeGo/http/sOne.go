package main

import (
	"fmt"
)

func printss() {
	fmt.Println("123")
	defer fmt.Println("456")
	panic("errror")
}
