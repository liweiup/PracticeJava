package main

import (
	"fmt"
	"time"
)

func mainxx() {
	msg := "hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "11"
	time.Sleep(100 * time.Millisecond)
}
