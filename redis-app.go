package main

import (
	"go-learning/service"
	"fmt"
)

func main() {
	service.SetMini();
	fmt.Println("从redis设置分")
	service.SetHour();
	fmt.Println("从redis设置时")
	hour := service.Get("hour")
	fmt.Println("从redis获取到分\n", hour)
	minute := service.Get("minute")
	fmt.Println("从redis获取到分\n", minute)
}
