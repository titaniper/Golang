package main

import (
	"fmt"
	"github.com/titaniper/hello/src/day1"
	"github.com/titaniper/hello/src/day2"
)

func main() {
	fmt.Println("Hello, World!")

	//doDay1();
	doDay2()
}

func doDay1() {
	fmt.Println("Hello, World!")

	day1.Iterator()

	doMap := &day1.DoMap{}
	doMap.Run()
}

func doDay2() {
	fmt.Println("Hello, World!")

	day2.Channel()
	//day2.Goroutine()
}
