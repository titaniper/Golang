package day2

import "fmt"

// Person 구조체 정의
// - 구조체는 필드들을 묶어서 하나의 타입으로 정의할 수 있는 방법입니다
// - 구조체는 객체 지향 언어의 클래스와 유사하지만 상속은 지원하지 않습니다.

type Address struct {
	City, State string
}

// Embedded Struct: 다른 구조체를 필드로 포함할 수 있습니다. 이를 통해 간접적으로 상속과 비슷한 기능을 구현할 수 있습니다.
// 익명 필드: 필드 이름을 생략하고 타입만 지정할 수 있습니다. 익명 필드는 해당 타입의 필드로 간주됩니다.
type Person struct {
	Name    string
	Age     int
	Address // 익명 필드
}

func Struct() {
	// 구조체 인스턴스 생성
	p := Person{Name: "Alice", Age: 30}

	// 구조체 필드 접근
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
