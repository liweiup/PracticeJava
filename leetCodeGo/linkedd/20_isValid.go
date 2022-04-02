package main

import (
	"fmt"
	//"strings"
)

func InSliceString(e string, slice []string) bool {
	for _, s := range slice {
		if s == e {
			return true
		}
	}
	return false
}
func isValid(s string) bool {
	bArr := []byte(s)
	strArr := []string{}
	for _, b := range bArr {
		if !InSliceString(string(b),strArr) {
			strArr = append(strArr, string(b))
		} else {
			fmt.Println()
		}
	}
	fmt.Println(strArr)
	return true
}

//func isValid_v2(s string) bool {
//	//strArr := []string{
//	//	"{}",
//	//	"[]",
//	//	"()",
//	//}
//	//strings.Replace()
//}
