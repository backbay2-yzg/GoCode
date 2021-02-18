package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "3.14 123.123 .68 haha 1.0 abc 7. ab.3 66.6 123."
	//ret:=regexp.MustCompile(`a[^0-9a-z]c`)
	ret := regexp.MustCompile(`\d\.\d`)
	all := ret.FindAllStringSubmatch(str, -1)
	fmt.Println("all is : ", all)
}
