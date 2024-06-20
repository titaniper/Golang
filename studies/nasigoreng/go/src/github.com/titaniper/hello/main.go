package main

import (
	"fmt"
	hello "github.com/titaniper/hello/src" // myproject는 모듈 이름입니다
)

func main() {
	fmt.Println("Hello, World!")

	hello.Iterator()

	doMap := &hello.DoMap{}
	doMap.Run()
}
