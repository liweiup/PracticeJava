package main

import "fmt"



type Write interface {
	Write([]byte) (int, error)
}

type ConsoleWrite struct {
}

func (c ConsoleWrite) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
