package day1

import "fmt"

type Pointer struct {
	Name string
	Age  int
	City string
}

// Golang에서 struct의 필드 값을 변경할 때 포인터 리시버를 사용하지 않으면, 해당 메서드가 struct의 복사본을 받게 되어 원본 struct의 필드 값은 변경되지 않습니다. 이는 Golang의 값 전달 방식 때문입니다.
func (p *Pointer) Hi() {
	nums := []int{1, 2, 3, 4, 5}

	nums = append(nums, 6)

	// 슬라이스 요소 접근 및 변경
	fmt.Println("First element:", nums[0])
	nums[0] = 10

	// 슬라이스 반복
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 슬라이스 길이와 용량
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
}
