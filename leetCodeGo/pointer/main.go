package main

import "fmt"

func main() {
	a := 5
	var b *int = &a

	fmt.Println(a, b, *&a)

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

}
