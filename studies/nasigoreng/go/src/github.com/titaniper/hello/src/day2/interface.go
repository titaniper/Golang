package day2

import "fmt"

// 인터페이스는 메서드들의 집합을 정의하며, 그 메서드들을 구현하는 모든 타입이 해당 인터페이스를 구현했다고 할 수 있습니다.
// - 인터페이스는 Go의 다형성을 지원합니다.
// 빈 인터페이스 (interface{}): 모든 타입을 저장할 수 있는 특수한 인터페이스입니다.
// 타입 스위치: 인터페이스 변수가 담고 있는 실제 타입을 검사할 때 사용됩니다.

// 인터페이스 정의
type Describer interface {
	Describe() string
}

// Person 구조체 정의
//type Person struct {
//	Name string
//	Age  int
//}

// Person 구조체에 대한 메서드 정의
func (p Person) Describe() string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

// 인터페이스 사용 예시
func Interface() {
	var d Describer
	p := Person{Name: "Alice", Age: 30}
	d = p
	fmt.Println(d.Describe())

	// 빈 인터페이스 사용 예시
	var i interface{}
	i = "Hello"
	fmt.Println(i)
}

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeSwitch() {
	describe(21)
	describe("hello")
	describe(true)
}
