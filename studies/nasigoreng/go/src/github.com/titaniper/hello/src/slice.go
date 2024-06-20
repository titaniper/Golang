package src

import "fmt"

func DoSlice() {
	// 슬라이스 생성 및 초기화
	nums := []int{1, 2, 3, 4, 5}

	// 요소 추가
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
