package main

import (
	"fmt"
	go_base "go-learning/go-base"
)

func main() {

	//go_base.BaseType()

	encode := go_base.Md5("Shanghai")
	fmt.Println("golang用md5加密32位：", encode)
}
