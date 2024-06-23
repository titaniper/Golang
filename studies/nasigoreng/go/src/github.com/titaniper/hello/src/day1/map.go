package day1

import "fmt"

type DoMap struct {
}

// NOTE: Golang에서 struct의 필드 값을 변경할 때 포인터 리시버를 사용하지 않으면, 해당 메서드가 struct의 복사본을 받게 되어 원본 struct의 필드 값은 변경되지 않습니다. 이는 Golang의 값 전달 방식 때문입니다.
func (d *DoMap) Run() {
	// 맵 생성 및 초기화
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// 요소 추가
	ages["Charlie"] = 35

	// 요소 접근
	fmt.Println("Alice's age:", ages["Alice"])

	// 요소 삭제
	delete(ages, "Bob")

	// 맵 요소 반복
	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	// 요소 존재 여부 확인
	age, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob does not exist in the map")
	}
}

func (DoMap) copy() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
}

func (DoMap) hi() {
	// 배열로부터 슬라이스 생성
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // [2, 3, 4]
	fmt.Println("Slice from array:", slice)

	// 리터럴을 사용하여 슬라이스 생성
	sliceLiteral := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice literal:", sliceLiteral)

	// make 함수를 사용하여 슬라이스 생성
	sliceMake := make([]int, 5)
	fmt.Println("Slice with make:", sliceMake)

	// 슬라이스 길이 및 용량 확인
	fmt.Println("Length of slice:", len(sliceLiteral))
	fmt.Println("Capacity of slice:", cap(sliceLiteral))

	// 슬라이스에 요소 추가
	sliceLiteral = append(sliceLiteral, 60)
	fmt.Println("Slice after append:", sliceLiteral)

	// 슬라이스 복사
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Source slice:", src)
	fmt.Println("Destination slice after copy:", dst)
}
