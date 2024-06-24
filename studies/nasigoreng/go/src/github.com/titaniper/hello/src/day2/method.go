package day2

import "fmt"

// Person 구조체 정의
//type Person struct {
//	Name string
//	Age  int
//}

// 메서드는 특정 타입에 속한 함수를 의미합니다. 메서드는 구조체와 함께 자주 사용됩니다.
// 포인터 리시버: 메서드 리시버가 구조체 포인터인 경우, 메서드 내에서 리시버의 필드를 수정할 수 있습니다.
// 값 리시버: 메서드 리시버가 값인 경우, 메서드 내에서 리시버의 필드를 수정해도 원본 값은 변경되지 않습니다.
// 메서드 오버로딩은 불가능: Go에서는 메서드 오버로딩을 지원하지 않습니다. 메서드 이름은 고유해야 합니다.

// Person 타입에 대한 메서드 정의
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// 포인터 리시버 메서드
func (p *Person) HaveBirthday() {
	p.Age++
}

type Circle struct {
	Radius float64
}

// 값 리시버 메서드
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func Method() {
	p := Person{Name: "Alice", Age: 30}
	p.Greet() // "Hello, my name is Alice" 출력
}
