package main

import (
	"fmt"
	"github.com/titaniper/hello/src/day1"
)

func main() {
	fmt.Println("Hello, World!")

	day1.Iterator()

	doMap := &day1.DoMap{}
	doMap.Run()
}
