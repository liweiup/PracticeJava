package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("sd")
	var i int = 42
	fmt.Printf("%v %T \n", i, i)
	j := strconv.Itoa(i)
	fmt.Printf("%v %T\n", j, j)

	var ui uint8 = 19
	fmt.Printf("%v %T\n", ui, ui)

	var ir = 1
	ir = ir << 1
	fmt.Printf("十进制%v 二进制%b\n ", ir, ir)
	fmt.Printf("十进制%v 二进制%b\n ", ir, ir)
	ir = ir << 1
	fmt.Printf("十进制%v 二进制%b\n ", ir, ir)
	ir = ir << 1
	fmt.Printf("十进制%v 二进制%b\n ", ir, ir)

	ADD := 1 // 增加权限
	UPD := 2 // 修改权限
	SEL := 4 // 修改权限
	fmt.Println(ADD | UPD | SEL)

	arr := []string{"a", "b", "c"}
	arr = append(arr, "ss")

	ar := make([]string, 1)
	ar = append(ar, []string{"a", "b", "c", "d", "e"}...)
	for i := 0; i < len(ar); i++ {
		fmt.Println(ar[i])
	}
	fmt.Printf("%v %T cap %v\n", ar, ar, cap(ar))

	fa := []int{1, 2, 3, 4, 5}
	fb := append(fa[:2], fa[3:]...)
	fmt.Printf("%v %T cap %v\n", fa, fa, cap(fa))
	fmt.Printf("%v %T cap %v\n", fb, fb, cap(fb))

}
