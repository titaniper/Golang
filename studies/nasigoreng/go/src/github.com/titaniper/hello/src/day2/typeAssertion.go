package day2

// **타입 단언(Type Assertion)**은
//- 인터페이스 타입 변수의 실제 값을 특정 타입으로 추론하는 데 사용됩니다.
//- 인터페이스는 다양한 타입의 값을 담을 수 있는데, 타입 단언을 통해 인터페이스에 저장된 값을 우리가 기대하는 특정 타입으로 변환할 수 있습니다.
//

import (
	"fmt"
)

func TypeAssertion() {
	var i interface{} = "Hello, Go!"

	// 타입 단언 사용 - 성공적인 경우
	str, ok := i.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Type assertion to string failed")
	}

	// 타입 단언 사용 - 실패한 경우
	num, ok := i.(int)
	if ok {
		fmt.Println("Integer value:", num)
	} else {
		fmt.Println("Type assertion to int failed")
	}

	// 타입 단언을 사용하여 값을 바로 가져오는 경우 - 패닉 발생 가능
	// num2 := i.(int) // 이 경우 타입 단언이 실패하면 런타임 에러가 발생
	// fmt.Println("Integer value:", num2)
}
