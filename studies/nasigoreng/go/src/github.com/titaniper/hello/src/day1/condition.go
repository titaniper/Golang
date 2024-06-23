package day1

import "fmt"

func DoCondition() {
	x := 10

	if x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is positive")
	}

	// 조건문 내에서 변수 선언
	if y := 20; y > 15 {
		fmt.Println("y is greater than 15")
	}
}
