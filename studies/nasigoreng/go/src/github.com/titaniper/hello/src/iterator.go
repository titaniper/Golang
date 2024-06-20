package src

import "fmt"

func Iterator() {
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

func a() {
	// 0부터 4까지 반복
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration", i)
	}
}

func b() {
	i := 0
	for i < 5 {
		fmt.Println("While-like iteration", i)
		i++
	}
}

func c() {
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

func d() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}

func e() {
	str := "Hello, World!"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}

func f() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 짝수일 경우 건너뛰기
		}
		fmt.Println("Odd number:", i)
	}
}

func i() {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, nums[i])
	}
}

func j() {
	nums := []int{1, 2, 3, 4, 5}
	for _, value := range nums {
		fmt.Println("Value:", value)
	}
}
