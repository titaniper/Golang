package day2

import (
	"errors"
	"fmt"
)

// Go에서는 오류 처리를 위해 error 인터페이스를 사용합니다.
// - 함수는 error 타입의 값을 반환하여 오류 발생을 알립니다.
// - 커스텀 오류 타입: error 인터페이스를 구현하여 사용자 정의 오류 타입을 만들 수 있습니다.
// - fmt.Errorf 함수: 오류 메시지를 포맷팅하여 생성할 수 있습니

// 오류를 반환하는 함수 정의
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func ErrorHandling() {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func formatError() {
	// NOTE: 포맷
	err := fmt.Errorf("an error occurred: %s", "something went wrong")
	fmt.Println(err)
}

// 커스텀 오류 타입
type MyError struct {
	When string
	What string
}

func Test() *string {
	return nil
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %s, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		When: "now",
		What: "something went wrong",
	}
}
