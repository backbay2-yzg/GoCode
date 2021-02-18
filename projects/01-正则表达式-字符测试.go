package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc cat aMc azc cba"
	//ret:=regexp.MustCompile(`a[^0-9a-z]c`)
	ret := regexp.MustCompile(`a.c`)
	all := ret.FindAllStringSubmatch(str, -1)
	fmt.Println("all is : ", all)
}
